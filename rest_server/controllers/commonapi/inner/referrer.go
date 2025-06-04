package inner

import (
	"time"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/point_manager_server"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/LumiWave/inno-dashboard/rest_server/util"
)

func ProcReferrer(innoUID, myInnoUID string, auid int64, referrerlRess []*context.RefreralResponse) {
	// referrerlRes db에 추천인 등록이 정상적으로 되었다면 포인트 추가 해준다
	if len(referrerlRess) == 0 {
		return
	}

	// 추천인 코드가 있다면 프리세일 정보가 존재 하는지 확인해서 추천인과, 피추천인 둘다 포인트 보상처리를 해준다.
	// 1. 유요한 innoUID 인지 체크
	referreeAUID, _, _, err := model.GetDB().USPAU_Get_Accounts_By_InnoUID(innoUID)
	if err != nil {
		return
	}

	// 2. presale 정보 수집
	preSales := getPreSales()

	for _, preSale := range preSales {
		// presale 기간인지 체크
		//if checkPreSalsPeriod(preSale) {
		for _, referrerlRes := range referrerlRess {
			if referrerlRes.SalesID == preSale.SalesID && referrerlRes.IsRegistered {
				// 정상적으로 등록된 경우에만 포인트 적립
				if preSale.ReferralRewardAppID == 0 {
					continue
				}

				// 내 포인트 업데이트
				if muid, databaseid, pointQuantity, err := checkMyPoint(preSale.ReferralRewardAppID, preSale.ReferralRewardPointID, auid, myInnoUID); err != nil {
					log.Debugf("checkMyPoint err : %v", err)
				} else {
					req := &point_manager_server.ReqPointAppUpdate{
						AppID:      preSale.ReferralRewardAppID,
						MUID:       muid,
						PointID:    preSale.ReferralRewardPointID,
						DatabaseID: databaseid,

						PreQuantity:    pointQuantity,
						AdjustQuantity: preSale.ReferrerPointQuantity,
					}
					if res, err := point_manager_server.GetInstance().PutPointAppUpdate(req); err != nil {
						log.Errorf("PutPointAppUpdate err : %v, req:%v", err, req)
					} else {
						if res.Return != resultcode.Result_Success {
							log.Errorf("PutPointAppUpdate fail return : %v, message : %v, req : %v ", res.Return, res.Message, req)
						}
					}
				}

				// 피추천인 포인트 업데이트
				if muid, databaseid, pointQuantity, err := checkMyPoint(preSale.ReferralRewardAppID, preSale.ReferralRewardPointID, referreeAUID, innoUID); err != nil {
					log.Debugf("checkMyPoint err : %v", err)
				} else {
					req := &point_manager_server.ReqPointAppUpdate{
						AppID:      preSale.ReferralRewardAppID,
						MUID:       muid,
						PointID:    preSale.ReferralRewardPointID,
						DatabaseID: databaseid,

						PreQuantity:    pointQuantity,
						AdjustQuantity: preSale.RefereePointQuantity,
					}
					if res, err := point_manager_server.GetInstance().PutPointAppUpdate(req); err != nil {
						log.Errorf("PutPointAppUpdate err : %v, req:%v", err, req)
					} else {
						if res.Return != resultcode.Result_Success {
							log.Errorf("PutPointAppUpdate fail return : %v, message : %v, req : %v ", res.Return, res.Message, req)
						}
					}
				}
			}
		}
		//}
	}
}

func checkMyPoint(appID, pointID, auid int64, innoUID string) (int64, int64, int64, error) { // muid, databaseid, pointQuantity, err
	_, membersMap, err := model.GetDB().USPAU_GetList_Members(auid)
	if err != nil {
		log.Debugf("USPAU_GetList_Members err: %v", err)
		return 0, 0, 0, err
	}

	bFind := false
	for _, member := range membersMap {
		if member.AppID != appID {
			continue
		}
		// 포인트 서버에서 현재 실제 정보 가져와서 포인트 정보가 존재한다면 리턴
		if memberInfo, err := point_manager_server.GetInstance().GetPointAppList(member.MUID, member.DatabaseID); err == nil {
			for _, point := range memberInfo.Points {
				if pointID == point.PointID { // 포인트를 찾았으면 정보 리턴
					return member.MUID, memberInfo.DatabaseID, point.Quantity, nil
				}
			}
		} else {
			log.Errorf("point_manager_server GetPointAppList error : %v", err)
			return 0, 0, 0, err
		}
	}

	if !bFind {
		// 최초 실행이여서 포인트정보가 생성되어 있지 않은 경우에 강제로 db에 포인트 생성
		if muid, databaseid, err := model.GetDB().USPAU_Auth_Members(innoUID, appID); err != nil {
			log.Errorf("USPAU_Auth_Members error : %v", err)
			return 0, 0, 0, err
		} else {
			req := &point_manager_server.ReqPointMemberRegister{
				AUID:       auid,
				MUID:       muid,
				AppID:      appID,
				DataBaseID: databaseid,
			}
			if res, err := point_manager_server.GetInstance().PostPointMemberRegister(req); err != nil {
				log.Debugf("PostPointMemberRegister err : %v", err)
				return 0, 0, 0, err
			} else {
				if res.Return == resultcode.Result_Success {
					for _, newPoint := range res.Value.Points {
						if newPoint.PointID == pointID {
							return muid, databaseid, newPoint.Quantity, nil // 최초 등록이기때문에 point를 직접 찾아서 뿌리지 않고 그냥 0으로 셋팅한다.
						}
					}
				} else {
					log.Debugf("PostPointMemberRegister return:%v, message:%v", res.Return, res.Message)
				}
			}
		}
	}

	return 0, 0, 0, err
}

func checkPreSalsPeriod(saleinfo *context.PreSales) bool {
	now := time.Now().UTC()
	//sales 시간 체크
	diff, err := util.IsTimeBetween(now, saleinfo.OpenStartSDT, saleinfo.OpenEndSDT)
	if err != nil {
		//return nil, errors.New("Sales Time Parse Error - " + saleinfo.OpenStartSDT + "/" + saleinfo.OpenEndSDT)
		return false
	}
	if !diff {
		//return nil, errors.New("This is not a time for presales - " + saleinfo.OpenStartSDT + "/" + saleinfo.OpenEndSDT)
		return false
	}
	return true
}

func getPreSales() []*context.PreSales {
	var presales []*context.PreSales
	if cachePreSales, err := model.GetDB().GetCachePreSales(); err != nil || len(cachePreSales) == 0 {
		if dbPreSales, err := model.GetDB().USPPR_Scan_PreSales(); err != nil {
			return nil
		} else {
			model.GetDB().SetCachePreSales(dbPreSales)
			presales = dbPreSales
		}
	} else {
		presales = cachePreSales
	}

	return presales
}

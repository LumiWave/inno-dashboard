package commonapi

import (
	"net/http"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/baseutil/otp_google"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 지갑 정보 조회
func GetMeWallets(c echo.Context, reqMeCoin *context.ReqMeCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletList, err := model.GetDB().USPAU_GetList_AccountCoins(reqMeCoin.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
	} else {
		resp.Value = walletList
		// tmpWalletList := []*context.MeCoin{}
		// for _, wallet := range walletList {
		// 	if wallet.CoinSymbol == "MATIC" {
		// 		continue
		// 	}
		// 	tmpWalletList = append(tmpWalletList, wallet)
		// }
		//resp.Value = tmpWalletList
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 포인트 리스트 조회
func GetMePointList(c echo.Context, reqMePoint *context.ReqMePoint) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if pointList, err := model.GetDB().USPAU_GetList_AccountPoints(reqMePoint.AUID, 0); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_PointList_Scan_Error)
	} else {
		if _, membersMap, err := model.GetDB().USPAU_GetList_Members(reqMePoint.AUID); err != nil {
			resp.SetReturn(resultcode.Result_Get_MemberList_Scan_Error)
		} else {
			for _, member := range membersMap {
				// 포인트 서버에서 현재 실제 정보 가져와서 merge
				if memberInfo, err := point_manager_server.GetInstance().GetPointAppList(member.MUID, member.DatabaseID); err == nil {
					for _, point := range memberInfo.Points {
						for _, mePoint := range pointList {
							if point.PointID == mePoint.PointID {
								mePoint.Quantity = point.Quantity
							}
						}
					}
				} else {
					log.Errorf("point_manager_server GetPointAppList error : %v", err)
				}
			}

			resp.Value = []*context.MePoint{}
			if pointList != nil {
				resp.Value = pointList
			}
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 코인 리스트 조회
func GetMeCoinList(c echo.Context, reqMeCoin *context.ReqMeCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if coinList, err := model.GetDB().USPAU_GetList_AccountCoins(reqMeCoin.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_CoinList_Scan_Error)
	} else {
		resp.Value = []*context.MeCoin{}
		if coinList != nil {
			resp.Value = coinList
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// google otp : qrcode용 uri 조회
func GetOtpUri(ctx *context.InnoDashboardContext) error {
	resp := new(base.BaseResponse)
	resp.Success()

	conf := config.GetInstance().Otp

	resp.Value = context.MeOtpUri{
		OtpUri: otp_google.MakeTimebaseUri(ctx.GetValue().InnoUID, ctx.GetValue().InnoUID, conf.IssueName),
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetOtpVerify(ctx *context.InnoDashboardContext, params *context.MeOtpVerify) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if !otp_google.VerifyTimebase(ctx.GetValue().InnoUID, params.OtpCode) {
		resp.SetReturn(resultcode.Result_Get_Me_Verify_otp_Error)
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostCoinReload(ctx *context.InnoDashboardContext, params *context.CoinReload) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// req := &point_manager_server.CoinReload{
	// 	AUID: params.AUID,
	// }

	// if res, err := point_manager_server.GetInstance().PostCoinReload(req); err == nil {
	// 	resp.Value = res.Value
	// } else {
	// 	log.Errorf("point_manager_server GetPointAppList error : %v", err)
	// 	resp.SetReturn(resultcode.ResultInternalServerError)
	// }

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqGetWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if res, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		resp.Value = res
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqPostWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if res, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		if walletData, ok := res.WalletData[params.WalletPlatform]; !ok {
			resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_UnsupportWallet_Error)
		} else {
			if walletData.IsRegistered {
				resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegistered_Error)
			} else {
				for _, basecoin := range model.GetDB().BaseCoins.Coins {
					if basecoin.WalletPlatform == params.WalletPlatform {
						if errType, err := model.GetDB().USPAU_Cnct_AccountWallets(params.AUID, basecoin.BaseCoinID, params.WalletAddress); err != nil {
							switch errType {
							case 2:
								resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegisteredDB_Error)
							case 3:
								resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegistered_AnotherAccount_Error)
							default:
								resp.SetReturn(resultcode.Result_DBError)
							}
						}
					}
				}
			}
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func DeleteWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqDeleteWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if res, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		if walletData, ok := res.WalletData[params.WalletPlatform]; !ok {
			resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_UnsupportWallet_Error)
		} else {
			if walletData.IsRegistered {
				if registDT, err := time.Parse(time.RFC3339, walletData.RegistDT); err != nil {
					log.Errorf("wallet registDT parse error : %v", err)
					resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_System_Error)
				} else {
					limitDT := registDT.Add(time.Hour * context.DeleteWalletHour)
					cmp := limitDT.Compare(time.Now())
					if cmp > 0 {
						resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_DeleteTime_Error) //24시간이 안됨
					} else {
						if walletData.WalletAddress != params.WalletAddress {
							resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_Diffrent_Wallet_Error) //현재 등록되어있는 지갑주소와 다름
						} else {
							for _, basecoin := range model.GetDB().BaseCoins.Coins {
								if basecoin.WalletPlatform == params.WalletPlatform {
									if err := model.GetDB().USPAU_Dscnct_AccountWallets(params.AUID, basecoin.BaseCoinID, params.WalletAddress); err != nil {
										resp.SetReturn(resultcode.Result_DBError)
									}
								}
							}
						}
					}
				}
			} else {
				resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_NoRegistered_Wallet_Error)
			}
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetWalletRegistInfo(auid int64) (*context.ResGetWalletRegist, int) {
	if UserWallets, err := model.GetDB().USPAU_GetList_AccountWallets(auid); err != nil {
		return nil, resultcode.Result_Get_Me_AUID_Empty
	} else {
		res := &context.ResGetWalletRegist{
			WalletData: make(map[string]*context.WalletRegistInfo),
		}
		for _, walletName := range model.GetDB().RegistWalletNames {
			res.WalletData[walletName] = &context.WalletRegistInfo{}
		}
		for _, userWallet := range UserWallets {
			if baseCoin, ok := model.GetDB().BaseCoinMapByCoinID[userWallet.BaseCoinID]; ok {
				if _, ok := res.WalletData[baseCoin.WalletPlatform]; ok {
					if auid > context.UserTypeLimit {
						res.WalletData[baseCoin.WalletPlatform].UserType = 2
					} else {
						res.WalletData[baseCoin.WalletPlatform].UserType = 1
					}
					switch userWallet.ConnectionStatus {
					case 1:
						res.WalletData[baseCoin.WalletPlatform].IsRegistered = true
						res.WalletData[baseCoin.WalletPlatform].WalletAddress = userWallet.WalletAddress
						res.WalletData[baseCoin.WalletPlatform].RegistDT = userWallet.ModifiedDT
					case 2:
						res.WalletData[baseCoin.WalletPlatform].LastDeleteWalletAddress = userWallet.WalletAddress
						res.WalletData[baseCoin.WalletPlatform].LastDeleteDT = userWallet.ModifiedDT
					default:
					}
				}
			}
		}
		//미등록상태이면 마지막 등록 주소는 내보내지않는다
		for _, userWallet := range res.WalletData {
			if !userWallet.IsRegistered {
				userWallet.LastDeleteWalletAddress = ""
				userWallet.LastDeleteDT = ""
			}
		}

		return res, 0
	}
}

package commonapi

import (
	"net/http"

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

	if walletList, err := model.GetDB().GetListAccountCoins(reqMeCoin.AUID); err != nil {
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

	if pointList, err := model.GetDB().GetListAccountPoints(reqMePoint.AUID, 0); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_PointList_Scan_Error)
	} else {
		if _, membersMap, err := model.GetDB().GetListMembers(reqMePoint.AUID); err != nil {
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

	if coinList, err := model.GetDB().GetListAccountCoins(reqMeCoin.AUID); err != nil {
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

func GetWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletRegistMap, err := model.GetDB().USPAU_GetList_AccountWallets(5); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_AUID_Empty)
	} else {
		resp.Value = &context.ResWalletRegist{}
		//체크로직
		if len(walletRegistMap) > 0 {
			// for _, wallet := range walletRegistMap {

			// }
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

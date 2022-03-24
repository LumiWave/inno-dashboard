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

// 전체 포인트, 코인 정보 리스트 조회
func GetSwapList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	swapList := context.SwapList{
		PointList: model.GetDB().ScanPoints,
		AppPoints: model.GetDB().AppPoints,
		CoinList:  model.GetDB().Coins,
		Swapable:  model.GetDB().SwapAble,
	}

	resp.Value = swapList

	return c.JSON(http.StatusOK, resp)
}

// Swap 가능 정보 조회 (최소, 변동률, 수수료)
func GetSwapEnable(c echo.Context, reqSwapEnable *context.ReqSwapEnable) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// Swap 처리
func PostSwap(ctx *context.InnoDashboardContext, reqSwapInfo *context.ReqSwapInfo) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// otp check
	if config.GetInstance().Otp.EnableSwap {
		if !otp_google.VerifyTimebase(ctx.GetValue().InnoUID, reqSwapInfo.OtpCode) {
			resp.SetReturn(resultcode.Result_Get_Me_Verify_otp_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	Lockkey := model.MakeMemberSwapLockKey(ctx.GetValue().AUID)
	unLock, err := model.AutoLock(Lockkey)
	if err != nil {
		return err
	}
	defer unLock()

	swapInfo := &point_manager_server.ReqSwapInfo{
		AUID: ctx.GetValue().AUID,
		SwapPoint: point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.AdjustPointQuantity,
			PointQuantity:         0,
		},
		SwapCoin: point_manager_server.SwapCoin{
			CoinID:               reqSwapInfo.CoinID,
			WalletAddress:        "",
			PreviousCoinQuantity: 0,
			AdjustCoinQuantity:   reqSwapInfo.AdjustCoinQuantity,
			CoinQuantity:         0,
		},
		LogID:   0,
		EventID: reqSwapInfo.EventID,
	}

	// SwapPoint 정보 추가
	if _, membersMap, err := model.GetDB().GetListMembers(swapInfo.AUID); err != nil {
		log.Errorf(resultcode.ResultCodeText[resultcode.Result_Get_MemberList_Scan_Error])
		resp.SetReturn(resultcode.Result_Get_MemberList_Scan_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		if member, ok := membersMap[swapInfo.AppID]; ok {
			swapInfo.MUID = member.MUID
			swapInfo.DatabaseID = member.DatabaseID
		} else {
			// swap 하려는 app point 정보가 없다.
			log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
			resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	// SwapCoin 정보 추가
	if coinList, err := model.GetDB().GetListAccountCoins(swapInfo.AUID); err != nil {
		log.Errorf("GetListAccountCoins error : %v", err)
		resp.SetReturn(resultcode.Result_Get_Me_CoinList_Scan_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		for _, coin := range coinList {
			if coin.CoinID == swapInfo.CoinID {
				swapInfo.PreviousCoinQuantity = coin.Quantity
				swapInfo.WalletAddress = coin.WalletAddress
				swapInfo.BaseCoinID = coin.BaseCoinID
				swapInfo.CoinQuantity = swapInfo.PreviousCoinQuantity + swapInfo.AdjustCoinQuantity
				break
			}
		}
		// 내 코인 정보 존재 확인 체크
		if len(swapInfo.WalletAddress) == 0 {
			log.Errorf(resultcode.ResultCodeText[resultcode.Result_Invalid_CoinID_Error])
			resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	swapInfo.LogID = context.LogID_exchange

	// 아래 체크 사항은 point manager server에서 처리한다.
	// 최소 변환 비율에 맞는지 체크
	// 전환 비율 계산 후 타당성 확인
	if resSwap, err := point_manager_server.GetInstance().PostPointCoinSwap(swapInfo); err != nil {
		log.Errorf("PostPointCoinSwap error : %v", err)
		resp.SetReturn(resultcode.Result_Unknown_Swap_Error)
	} else {
		if resSwap.Common.Return != 0 {
			resp.Return = resSwap.Return
			resp.Message = resSwap.Message
		} else {
			resp.Value = resSwap.ReqSwapInfo
		}
	}
	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

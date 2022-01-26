package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
)

// 외부 지갑으로 코인 전송
func PostTransfer(ctx *context.InnoDashboardContext, reqCoinTransfer *context.ReqCoinTransfer) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// 전송 코인량 검증, 수수료 계산 검증
	// 전송 할만큼의 코인 보유량 검증 (전송량 + 수수료)
	coinInfo, ok := model.GetDB().CoinsMap[reqCoinTransfer.CoinID]
	if !ok { // 플랫폼에 존재 하지 않는 코인
		resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	meCoin := &context.MeCoin{}
	if walletList, err := model.GetDB().GetListAccountCoins(ctx.GetValue().AUID); walletList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		for _, meWallet := range walletList {
			if meWallet.CoinID == reqCoinTransfer.CoinID {
				meCoin = meWallet
			}
		}
		if meCoin.CoinID == 0 { // 나에게 존재하지 않은 코인 전송을 요청
			resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
		if meCoin.Quantity < (reqCoinTransfer.Quantity + coinInfo.ExchangeFees) { // 보유량 부족
			resp.SetReturn(resultcode.Result_CoinTransfer_NotEnough_Coin)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	// 필요한 정보를 모아서 point-manager "외부 지갑으로 토큰 전송" 요청
	req := &point_manager_server.ReqCoinTransfer{
		AUID:          reqCoinTransfer.AUID,
		CoinID:        reqCoinTransfer.CoinID,
		CoinSymbol:    meCoin.CoinSymbol,
		ToAddress:     reqCoinTransfer.ToAddress,
		Quantity:      reqCoinTransfer.Quantity,
		TransferFee:   coinInfo.ExchangeFees,
		TotalQuantity: reqCoinTransfer.Quantity + coinInfo.ExchangeFees,
	}

	if res, err := point_manager_server.GetInstance().PostCoinTransfer(req); err != nil {
		resp.SetReturn(resultcode.ResultInternalServerError)
	} else {
		if res.Common.Return != 0 {
			resp.Return = res.Return
			resp.Message = res.Message
		} else {
			resp.Value = res.Value
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

// 코인 외부 지갑 전송 중인 상태 정보 요청
func GetCoinTransferExistInProgress(ctx *context.InnoDashboardContext, params *context.GetCoinTransferExistInProgress) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if res, err := point_manager_server.GetInstance().GetCoinTransferExistInProgress(params.AUID); err != nil {
		resp.SetReturn(resultcode.ResultInternalServerError)
	} else {
		resp.Return = res.Return
		resp.Message = res.Message
		if resp.Return == 0 {
			resp.Value = res.Value
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

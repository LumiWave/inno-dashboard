package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 외부 지갑으로 코인 전송
func PostTransfer(ctx *context.InnoDashboardContext, reqCoinTransfer *context.ReqCoinTransfer) error {
	resp := new(base.BaseResponse)
	resp.Success()

	log.Debugf("PostTransfer : %v", reqCoinTransfer)

	if !model.GetExternalTransferEnable() {
		resp.SetReturn(resultcode.Result_Error_IsCoinTransferExternalMaintenance)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	// 전송 코인량 검증, 수수료 계산 검증
	// 전송 할만큼의 코인 보유량 검증 (전송량 + 수수료)
	coinInfo, ok := model.GetDB().CoinsMap[reqCoinTransfer.CoinID]
	if !ok { // 플랫폼에 존재 하지 않는 코인
		resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	meCoin := &context.MeCoin{}
	meBaseCoin := &context.MeCoin{}
	if walletList, err := model.GetDB().GetListAccountCoins(ctx.GetValue().AUID); walletList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		for _, meWallet := range walletList {
			if meWallet.CoinID == reqCoinTransfer.CoinID {
				meCoin = meWallet
			}
			// if meWallet.CoinSymbol == model.GetDB().BaseCoinMapByCoinID[meWallet.BaseCoinID].BaseCoinSymbol {
			// 	meBaseCoin = meWallet
			// }
		}
		for _, meWallet := range walletList {
			if meWallet.CoinSymbol == model.GetDB().BaseCoinMapByCoinID[meCoin.BaseCoinID].BaseCoinSymbol {
				meBaseCoin = meWallet
			}
		}
		if meCoin.CoinID == 0 { // 나에게 존재하지 않은 코인 전송을 요청
			resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
		tempBaseCoin := model.GetDB().BaseCoinMapByCoinID[meCoin.BaseCoinID]
		if tempBaseCoin.IsUsedParentWallet { // 부모지갑에서 출금은 수수료까지 보유량 체크를 해야함
			if meCoin.Quantity < (reqCoinTransfer.Quantity + coinInfo.ExchangeFees) { // 보유량 부족
				resp.SetReturn(resultcode.Result_CoinTransfer_NotEnough_Coin)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			}
		} else { // 자식 지갑은 수수료를 따로 제외하지 않는다. 대신 basecoin 지갑에 해당 가스비가 남아 있는지 확인 한다.
			if meCoin.Quantity < reqCoinTransfer.Quantity { // 보유량 부족
				resp.SetReturn(resultcode.Result_CoinTransfer_NotEnough_Coin)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			}
			// 가스비 체크
			key := model.MakeCoinFeeKey(meCoin.CoinSymbol)
			if coinFee, err := model.GetDB().GetCacheCoinFee(key); err != nil {
				log.Errorf("GetCacheCoinFee err : %v", err)
				resp.SetReturn(resultcode.Result_CoinFee_NotExist)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			} else {
				if meBaseCoin.Quantity < coinFee.TransactionFee { // 가스비 부족
					resp.SetReturn(resultcode.Result_CoinFee_LackOfGas)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			}

		}

	}

	// 부모지갑에서 전송 해야 하는지 자식 지갑에서 전송 해야 하는지 체크
	baseCoin := model.GetDB().BaseCoinMapByCoinID[meCoin.BaseCoinID]

	if baseCoin.IsUsedParentWallet {
		// 필요한 정보를 모아서 point-manager 부모 지갑에서 전송 호출
		req := &point_manager_server.ReqCoinTransferFromParentWallet{
			AUID:       reqCoinTransfer.AUID,
			CoinID:     reqCoinTransfer.CoinID,
			CoinSymbol: meCoin.CoinSymbol,
			ToAddress:  reqCoinTransfer.ToAddress,
			Quantity:   reqCoinTransfer.Quantity,
		}

		if res, err := point_manager_server.GetInstance().PostCoinTransferFromParentWallet(req); err != nil {
			resp.SetReturn(resultcode.ResultInternalServerError)
		} else {
			if res.Common.Return != 0 {
				resp.Return = res.Return
				resp.Message = res.Message
			} else {
				resp.Value = res.Value
			}
		}
	} else {
		// 필요한 정보를 모아서 point-manager 특정 지갑 전송 호출
		req := &point_manager_server.ReqCoinTransferFromUserWallet{
			AUID:           reqCoinTransfer.AUID,
			CoinID:         reqCoinTransfer.CoinID,
			CoinSymbol:     meCoin.CoinSymbol,
			BaseCoinSymbol: baseCoin.BaseCoinSymbol,
			FromAddress:    meCoin.WalletAddress,
			ToAddress:      reqCoinTransfer.ToAddress,
			Quantity:       reqCoinTransfer.Quantity,
		}

		if res, err := point_manager_server.GetInstance().PostCoinTransferFromUserWallet(req); err != nil {
			resp.SetReturn(resultcode.ResultInternalServerError)
		} else {
			if res.Common.Return != 0 {
				resp.Return = res.Return
				resp.Message = res.Message
			} else {
				resp.Value = res.Value
			}
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

// 코인 외부 지갑 전송 중인 상태 정보 존재 하지 않는지 요청
func GetCoinTransferNotExistInProgress(ctx *context.InnoDashboardContext, params *context.GetCoinTransferExistInProgress) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if res, err := point_manager_server.GetInstance().GetCoinTransferNotExistInProgress(params.AUID); err != nil {
		resp.SetReturn(resultcode.ResultInternalServerError)
	} else {
		if res.Return == 0 {
			// 존재 해서 에러 발생
			resp.Return = 12200
			resp.Message = "Transfer inprogress"
			resp.Value = res.Value
		} else if res.Return != 12201 {
			resp.Return = res.Return
			resp.Message = res.Message
		} else if res.Return == 12202 {
			// 존재 하지 않으면 성공 처리
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetCoinTransferFee(c echo.Context, params *context.GetCoinFee) error {
	resp := new(base.BaseResponse)
	resp.Success()

	redisKey := model.MakeCoinFeeKey(params.BaseCoinSymbol)
	if coinFee, err := model.GetDB().GetCacheCoinFee(redisKey); err != nil {
		log.Errorf("GetCacheCoinFee err : %v", err)
		resp.SetReturn(resultcode.Result_CoinFee_NotExist)
	} else {
		resp.Value = coinFee
	}

	return c.JSON(http.StatusOK, resp)
}

package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Coin Transfer
type ReqCoinTransfer struct {
	AUID          int64  `json:"au_id" url:"au_id"`                   // 계정의 UID (Access Token에서 가져옴)
	CoinSymbol    string `json:"coin_symbol" url:"coin_symbol"`       // 코인 심볼
	ToAddress     string `json:"to_address" url:"to_address"`         // 보낼 지갑 주소
	Quantity      string `json:"quantity" url:"quantity"`             // 보낼 코인량
	TransferFee   string `json:"transfer_fee" url:"transfer_fee"`     // 전송 수수료
	TotalQuantity string `json:"total_quantity" url:"total_quantity"` // 보낼 코인량 + 전송 수수료
}

func (o *ReqCoinTransfer) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_CoinSymbol_Empty)
	}
	if len(o.ToAddress) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_ToAddress_Empty)
	}
	if len(o.Quantity) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_Quantity_Empty)
	}
	if len(o.TransferFee) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_TransferFee_Empty)
	}
	if len(o.TotalQuantity) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_TotalQuantity_Empty)
	}

	return nil
}

////////////////////////////////////////

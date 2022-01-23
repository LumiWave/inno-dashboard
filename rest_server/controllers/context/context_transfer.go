package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Coin Transfer
type ReqCoinTransfer struct {
	AUID      int64   `json:"au_id" url:"au_id"`           // 계정의 UID (Access Token에서 가져옴)
	CoinID    int64   `json:"coin_id" url:"coin_id"`       // 코인 심볼
	ToAddress string  `json:"to_address" url:"to_address"` // 보낼 지갑 주소
	Quantity  float64 `json:"quantity" url:"quantity"`     // 보낼 코인량
}

func (o *ReqCoinTransfer) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	if o.CoinID == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_CoinSymbol_Empty)
	}
	if len(o.ToAddress) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_ToAddress_Empty)
	}
	if o.Quantity == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinTransfer_Quantity_Empty)
	}

	return nil
}

////////////////////////////////////////

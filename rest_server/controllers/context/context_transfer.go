package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Coin Transfer
type ReqCoinTransfer struct {
	AUID        int64   `json:"au_id" url:"au_id"`               // 계정의 UID (Access Token에서 가져옴)
	CoinID      int64   `json:"coin_id" url:"coin_id"`           // 코인 심볼
	FromAddress string  `json:"from_address" url:"from_address"` // 전송자 지갑 주소
	ToAddress   string  `json:"to_address" url:"to_address"`     // 보낼 지갑 주소
	Quantity    float64 `json:"quantity" url:"quantity"`         // 보낼 코인량
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

///////// 외부 지갑으로 코인 전송 중인 상태 정보 요청
type GetCoinTransferExistInProgress struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func NewGetCoinTransferExistInProgress() *GetCoinTransferExistInProgress {
	return new(GetCoinTransferExistInProgress)
}

func (o *GetCoinTransferExistInProgress) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	} else if ctx.GetValue() == nil || o.AUID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AUID_Empty)
	}
	return nil
}

////////////////////////////////////////

///////// 코인 가스비 조회
type GetCoinFee struct {
	BaseCoinSymbol string `json:"coin_symbol" query:"coin_symbol"`
}

func NewGetCoinFee() *GetCoinFee {
	return new(GetCoinFee)
}

func (o *GetCoinFee) CheckValidate() *base.BaseResponse {
	if len(o.BaseCoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_CoinFee_BaseSymbol_Empty)
	}
	return nil
}

type ResGetCoinFee struct {
	BaseCoinID     int64   `json:"base_coin_id"`
	BaseCoinSymbol string  `json:"base_coin_symbol"`
	CoinID         int64   `json:"coin_id"`
	ConiSymbol     string  `json:"coin_symbol"`
	TransactionFee float64 `json:"transaction_fee"`
	GasPrice       float64 `json:"gas_price"`
}

////////////////////////////////////////

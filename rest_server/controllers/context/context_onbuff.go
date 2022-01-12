package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

// 가격 정보
type ReqPriceInfo struct {
	CoinSymbol string `json:"coin_symbol" query:"coin_symbol"`
}
type PriceInfo struct {
	CoinSymbol     string  `json:"coin_symbol,omitempty"`
	ONITPrice      float64 `json:"onit_price,omitempty"`
	KRWPrice       float64 `json:"krw_price,omitempty"`
	PriceTimeStamp int64   `json:"price_timestamp,omitempty"`
}

// 시세 정보
type QuoteInfo struct {
	StartDate string  `json:"start_date,omitempty"`
	EndDate   string  `json:"end_date,omitempty"`
	Price     float64 `json:"price,omitempty"`
	Prepare   float64 `json:"prepare,omitempty"`
}

// 시세 History
type QuoteHistory struct {
	PageInfoResponse `json:"page_info"`
	CoinSymbol       string      `json:"coin_symbol,omitempty"`
	Candle           string      `json:"candle,omitempty"`
	ItemCount        int64       `json:"item_count,omitempty"`
	Quotes           []QuoteInfo `json:"quotes"`
}

// 유동량
type LiquidityInfo struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Liquidity int64  `json:"liquidity,omitempty"`
	Prepare   int64  `json:"prepare,omitempty"`
}

// 유동량 History
type LiquidityHistory struct {
	PageInfoResponse `json:"page_info"`
	ItemCount        int64  `json:"item_count,omitempty"`
	Candle           string `json:"candle,omitempty"`
	Liquiditys       []LiquidityInfo
}

func (o *ReqPriceInfo) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Empty_CoinSymbol)
	}
	return nil
}

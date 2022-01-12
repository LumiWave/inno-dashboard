package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/util"
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

// Candle Base
type ReqBaseCandle struct {
	CoinSymbol string `json:"coin_symbol" query:"coin_symbol"`
	Count      string `json:"count" query:"count"`
	To         string `json:"to" query:"to"` // 마지막 캔들 시각 (UTC)
}

// Candle Minutes
type ReqCandleMinutes struct {
	ReqBaseCandle
	Unit string `json:"unit" query:"unit"`
}

// Candle Days
type ReqCandleDays struct {
	ReqBaseCandle
	ConvertingPriceUnit string `json:"converting_price_unit" query:"converting_price_unit"`
}

// Candle Weeks
type ReqCandleWeeks struct {
	ReqBaseCandle
}

// Candle Months
type ReqCandleMonths struct {
	ReqBaseCandle
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
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	return nil
}

func (o *ReqCandleMinutes) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	if len(o.Count) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCount)
	}
	if len(o.Unit) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyUnit)
	}
	if !util.Contains(o.Unit, []string{"1", "3", "5", "15", "10", "30", "60", "240"}) {
		return base.MakeBaseResponse(resultcode.Result_Upbit_InvalidUnit)
	}
	return nil
}

func (o *ReqCandleDays) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	if len(o.Count) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCount)
	}
	return nil
}

func (o *ReqCandleWeeks) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	if len(o.Count) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCount)
	}
	return nil
}

func (o *ReqCandleMonths) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	if len(o.Count) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCount)
	}
	return nil
}

package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/util"
)

///////// 가격 정보
type ReqPriceInfo struct {
	CoinSymbol string `json:"coin_symbol" query:"coin_symbol"`
}

type PriceInfo struct {
	CoinSymbol     string  `json:"coin_symbol,omitempty"`
	ONITPrice      float64 `json:"onit_price,omitempty"`
	KRWPrice       float64 `json:"krw_price,omitempty"`
	PriceTimeStamp int64   `json:"price_timestamp,omitempty"`
}

func (o *ReqPriceInfo) CheckValidate() *base.BaseResponse {
	if len(o.CoinSymbol) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Upbit_EmptyCoinSymbol)
	}
	return nil
}

////////////////////////////////////////

///////// Candle Base
type ReqBaseCandle struct {
	CoinSymbol string `json:"coin_symbol" query:"coin_symbol"`
	Count      string `json:"count" query:"count"`
	To         string `json:"to" query:"to"` // 마지막 캔들 시각 (UTC)
}

////////////////////////////////////////

///////// Candle Minutes
type ReqCandleMinutes struct {
	ReqBaseCandle
	Unit string `json:"unit" query:"unit"`
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

////////////////////////////////////////

///////// Candle Days
type ReqCandleDays struct {
	ReqBaseCandle
	ConvertingPriceUnit string `json:"converting_price_unit" query:"converting_price_unit"`
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

////////////////////////////////////////

///////// Candle Weeks
type ReqCandleWeeks struct {
	ReqBaseCandle
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

////////////////////////////////////////

///////// Candle Months
type ReqCandleMonths struct {
	ReqBaseCandle
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

////////////////////////////////////////

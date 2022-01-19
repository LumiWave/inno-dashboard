package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Point Info
type PointInfo struct {
	PointId              int64   `json:"point_id,omitempty"`
	PointName            string  `json:"point_name,omitempty"`
	IconUrl              string  `json:"icon_url,omitempty"`
	MinExchangeQuantity  int64   `json:"minimum_exchange_quantity"`
	ExchangeRatio        float64 `json:"exchange_ratio"`
	DaliyLimitedQuantity int64   `json:"daliy_limited_quantity,omitempty"`
}

type PointList struct {
	Points []PointInfo `json:"points"`
}

type AppPointInfo struct {
	AppId   int64        `json:"app_id,omitempty"`
	AppName string       `json:"app_name,omitempty"`
	IconUrl string       `json:"icon_url"`
	Points  []*PointInfo `json:"points"`
}

type AppPoints struct {
	Apps []*AppPointInfo `json:"apps"`
}

////////////////////////////////////////

///////// Coin Info
type CoinInfo struct {
	CoinId          int64   `json:"coin_id,omitempty"`
	CoinSymbol      string  `json:"coin_symbol,omitempty"`
	ContractAddress string  `json:"contract_address,omitempty"`
	IconUrl         string  `json:"icon_url,omitempty"`
	ExchangeFees    float64 `json:"exchange_fees"`
}

type CoinList struct {
	Coins []*CoinInfo `json:"coins"`
}

////////////////////////////////////////

///////// AppPointDailyInfo
type ReqAppPoint struct {
	AppId int64 `json:"app_id" query:"app_id"`
}

func (o *ReqAppPoint) CheckValidate() *base.BaseResponse {
	if o.AppId == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_App_AppID_Empty)
	}
	return nil
}

type AppPointDailyInfo struct {
	PointID               int64  `json:"point_id" query:"point_id"`
	PointName             string `json:"point_name" query:"point_name"`
	DailyQuantity         int64  `json:"daily_quantity" query:"daily_quantity"`
	DailyExchangeQuantity int64  `json:"daily_exchange_quantity" query:"daily_exchange_quantity"`
}

////////////////////////////////////////

///////// AppCoinDailyInfo
type ReqAppCoin struct {
	AppId int64 `json:"app_id" query:"app_id"`
}

func (o *ReqAppCoin) CheckValidate() *base.BaseResponse {
	if o.AppId == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_App_AppID_Empty)
	}
	return nil
}

type AppCoinDailyInfo struct {
	CoinID                int64   `json:"coin_id" query:"coin_id"`
	CoinSymbol            string  `json:"coin_symbol" query:"coin_symbol"`
	DailyQuantity         float64 `json:"daily_quantity" query:"daily_quantity"`
	DailyExchangeQuantity float64 `json:"daily_exchange_quantity" query:"daily_exchange_quantity"`
}

////////////////////////////////////////

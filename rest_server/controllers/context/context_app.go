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
	MinExchangeQuantity  int64   `json:"minimum_exchange_quantity,omitempty"`
	ExchangeRatio        float64 `json:"exchange_ratio,omitempty"`
	DaliyLimitedQuantity int64   `json:"daliy_limited_quantity,omitempty"`
}

type PointList struct {
	Points []*PointInfo `json:"points"`
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
	CoinName        string  `json:"coin_name"`
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
type ReqAppPointDaily struct {
	AppID   int64 `json:"app_id" query:"app_id"`
	PointID int64 `json:"point_id" query:"point_id"`
}

func (o *ReqAppPointDaily) CheckValidate() *base.BaseResponse {
	if o.AppID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_AppID_Error)
	}
	return nil
}

type ResAppPointDaily struct {
	AppID                    int64 `json:"app_id" query:"app_id"`
	PointID                  int64 `json:"point_id" query:"point_id"`
	TodayAcqQuantity         int64 `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayAcqExchangeQuantity int64 `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
}

////////////////////////////////////////

///////// AppCoinDailyInfo
type ReqAppCoinDaily struct {
	CoinID int64 `json:"coin_id" query:"coin_id"`
}

func (o *ReqAppCoinDaily) CheckValidate() *base.BaseResponse {
	if o.CoinID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_App_CoinID_Empty)
	}
	return nil
}

type ResAppCoinDaily struct {
	CoinID                   int64   `json:"coin_id" query:"coin_id"`
	TodayAcqQuantity         float64 `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayAcqExchangeQuantity float64 `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
}

////////////////////////////////////////

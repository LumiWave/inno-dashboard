package context

import (
	"time"

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
type ReqAppPointDaily struct {
	AppID   int64 `json:"app_id" query:"app_id"`
	PointID int64 `json:"point_id" query:"point_id"`
}

func (o *ReqAppPointDaily) CheckValidate() *base.BaseResponse {
	if o.AppID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_App_AppID_Empty)
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
		return base.MakeBaseResponse(resultcode.Result_Get_App_AppID_Empty)
	}
	return nil
}

type ResAppCoinDaily struct {
	CoinID                   int64   `json:"coin_id" query:"coin_id"`
	TodayAcqQuantity         float64 `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayAcqExchangeQuantity float64 `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
}

////////////////////////////////////////

///////// Get List Coin Liquidity (코인 유동량)
type ReqCoinLiquidity struct {
	BaseDate time.Time `json:"base_date" query:"base_date"` // 기준날짜
	CoinID   int64     `json:"coin_id" query:"coin_id"`     // 코인ID
	Interval int64     `json:"interval" query:"interval"`   // 기간 간격 (0:오늘)
}

type CoinLiquidity struct {
	BaseDate             time.Time `json:"base_date" query:"base_date"`                           // 기준날짜
	AcqQuantity          float64   `json:"acq_quantity" query:"acq_quantity"`                     // 획득량
	AcqCount             int64     `json:"acq_count" query:"acq_count"`                           // 획득 횟수
	CnsmQuantity         float64   `json:"cnsm_quntity" query:"cnsm_quntity"`                     // 소모량
	CnsmCount            int64     `json:"cnsm_count" query:"cnsm_count"`                         // 소모 횟수
	AcqExchangeQuantity  float64   `json:"acq_exchange_quantity" query:"acq_exchange_quantity"`   // 획득 전환량
	PointsToCoinsCount   int64     `json:"points_to_coins_count" query:"points_to_coins_count"`   // 획득 전환 횟수
	CnsmExchangeQuantity float64   `json:"cnsm_exchange_quantity" query:"cnsm_exchange_quantity"` // 소모 전환량
	CoinsToPointsCount   int64     `json:"coins_to_points_count" query:"coins_to_points_count"`   // 소모 전환 횟수
}

////////////////////////////////////////

///////// Get List Point Liquidity (포인트 유동량)
type ReqPointLiquidity struct {
	BaseDate time.Time `json:"base_date" query:"base_date"` // 기준날짜
	AppID    int64     `json:"app_id" query:"app_id"`       // 앱ID
	PointID  int64     `json:"point_id" query:"point_id"`   // 포인트ID
	Interval int64     `json:"interval" query:"interval"`   // 기간 간격 (0:오늘)
}

type PointLiquidity struct {
	BaseDate             time.Time `json:"base_date" query:"base_date"`                           // 기준날짜
	AcqQuantity          int64     `json:"acq_quantity" query:"acq_quantity"`                     // 획득량
	AcqCount             int64     `json:"acq_count" query:"acq_count"`                           // 획득 횟수
	CnsmQuantity         int64     `json:"cnsm_quntity" query:"cnsm_quntity"`                     // 소모량
	CnsmCount            int64     `json:"cnsm_count" query:"cnsm_count"`                         // 소모 횟수
	AcqExchangeQuantity  int64     `json:"acq_exchange_quantity" query:"acq_exchange_quantity"`   // 획득 전환량
	PointsToCoinsCount   int64     `json:"points_to_coins_count" query:"points_to_coins_count"`   // 획득 전환 횟수
	CnsmExchangeQuantity int64     `json:"cnsm_exchange_quantity" query:"cnsm_exchange_quantity"` //소모 전환량
	CoinsToPointsCount   int64     `json:"coins_to_points_count" query:"coins_to_points_count"`   // 소모 전환 횟수
}

////////////////////////////////////////

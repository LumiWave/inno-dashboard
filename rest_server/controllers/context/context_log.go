package context

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Get List Coin Liquidity (코인 유동량)
var CandleTypeOfCoin = map[string]string{
	"hour":  "USPW_GetList_HourlyCoins",
	"day":   "USPW_GetList_DailyCoins",
	"week":  "USPW_GetList_WeeklyCoins",
	"month": "USPW_GetList_MonthlyCoins",
}

type ReqCoinLiquidity struct {
	CoinID int64  `json:"coin_id" query:"coin_id"` // 코인ID
	Candle string `json:"candle" query:"candle"`   // 조회 종류 "hour", "day", "week", "month"
	PageInfo

	//internal param
	BaseDate string // 기준날짜
	Interval int64  // 기간 간격 (0:오늘)
}

func NewReqCoinLiquidity() *ReqCoinLiquidity {
	return new(ReqCoinLiquidity)
}

func (o *ReqCoinLiquidity) CheckValidate() *base.BaseResponse {
	if o.CoinID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_CoinID_Error)
	}

	return nil
}

type CoinLiquidity struct {
	BaseDate             time.Time `json:"base_date" query:"base_date"` // 기준날짜
	BaseDateToNumber     int64     `json:"base_date_number" query:"base_date_number"`
	AcqQuantity          float64   `json:"acq_quantity" query:"acq_quantity"`                     // 획득량
	AcqCount             int64     `json:"acq_count" query:"acq_count"`                           // 획득 횟수
	CnsmQuantity         float64   `json:"cnsm_quantity" query:"cnsm_quantity"`                   // 소모량
	CnsmCount            int64     `json:"cnsm_count" query:"cnsm_count"`                         // 소모 횟수
	AcqExchangeQuantity  float64   `json:"acq_exchange_quantity" query:"acq_exchange_quantity"`   // 획득 전환량
	PointsToCoinsCount   int64     `json:"points_to_coins_count" query:"points_to_coins_count"`   // 획득 전환 횟수
	CnsmExchangeQuantity float64   `json:"cnsm_exchange_quantity" query:"cnsm_exchange_quantity"` // 소모 전환량
	CoinsToPointsCount   int64     `json:"coins_to_points_count" query:"coins_to_points_count"`   // 소모 전환 횟수
}

type ResCoinLiquidity struct {
	Count          int              `json:"coin_liquidity_count"`
	CoinLiquiditys []*CoinLiquidity `json:"coin_liquiditys"`
}

////////////////////////////////////////

///////// Get List Point Liquidity (포인트 유동량)
var CandleTypeOfPoint = map[string]string{
	"hour":  "USPW_GetList_HourlyPoints",
	"day":   "USPW_GetList_DailyPoints",
	"week":  "USPW_GetList_WeeklyPoints",
	"month": "USPW_GetList_MonthlyPoints",
}

type ReqPointLiquidity struct {
	AppID   int64  `json:"app_id" query:"app_id"`     // 앱ID
	PointID int64  `json:"point_id" query:"point_id"` // 포인트ID
	Candle  string `json:"candle" query:"candle"`     // 조회 종류 "hour", "day", "week", "month"
	PageInfo

	// internal param
	BaseDate string // 기준날짜
	Interval int64  // 기간 간격 (0:오늘)
}

func NewReqPointLiquidity() *ReqPointLiquidity {
	return new(ReqPointLiquidity)
}

func (o *ReqPointLiquidity) CheckValidate() *base.BaseResponse {
	if o.AppID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_AppID_Error)
	}
	if o.PointID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_PointID_Error)
	}
	return nil
}

type PointLiquidity struct {
	BaseDate             time.Time `json:"base_date" query:"base_date"` // 기준날짜
	BaseDateToNumber     int64     `json:"base_date_number" query:"base_date_number"`
	AcqQuantity          int64     `json:"acq_quantity" query:"acq_quantity"`                     // 획득량
	AcqCount             int64     `json:"acq_count" query:"acq_count"`                           // 획득 횟수
	CnsmQuantity         int64     `json:"cnsm_quantity" query:"cnsm_quantity"`                   // 소모량
	CnsmCount            int64     `json:"cnsm_count" query:"cnsm_count"`                         // 소모 횟수
	AcqExchangeQuantity  int64     `json:"acq_exchange_quantity" query:"acq_exchange_quantity"`   // 획득 전환량
	PointsToCoinsCount   int64     `json:"points_to_coins_count" query:"points_to_coins_count"`   // 획득 전환 횟수
	CnsmExchangeQuantity int64     `json:"cnsm_exchange_quantity" query:"cnsm_exchange_quantity"` //소모 전환량
	CoinsToPointsCount   int64     `json:"coins_to_points_count" query:"coins_to_points_count"`   // 소모 전환 횟수
}

type ResPointLiquidity struct {
	Count           int               `json:"point_liquidity_count"`
	PointLiquiditys []*PointLiquidity `json:"point_liquiditys"`
}

////////////////////////////////////////

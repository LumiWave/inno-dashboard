package context

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Wallet Info
type ReqMeWallet struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func (o *ReqMeWallet) CheckValidate() *base.BaseResponse {
	if o.AUID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AUID_Empty)
	}
	return nil
}

type MeWalletInfo struct {
	CoinID            int64     `json:"coin_id" query:"coin_id"`
	WalletAddress     string    `json:"wallet_address" query:"wallet_address"`
	Quantity          float64   `json:"quantity" query:"quantity"`
	TodayAcqQuantity  float64   `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity float64   `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	ResetDate         time.Time `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

///////// Me Point List
type ReqMePoint struct {
	AUID int64 `json:"au_id" query:"au_id"`
	MUID int64 `json:"mu_id" query:"mu_id"`
}

func (o *ReqMePoint) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}

	return nil
}

type MePoint struct {
	AppID                int64  `json:"app_id" query:"app_id"`
	PointID              int64  `json:"point_id" query:"point_id"`
	Quantity             int64  `json:"quantity"`
	TodayLimitedQuantity int64  `json:"today_limited_quantity" query:"today_limited_quantity"`
	TodayAcqQuantity     int64  `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity    int64  `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	ResetDate            string `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

///////// Me Coin List
type ReqMeCoin struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func (o *ReqMeCoin) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

type MeCoin struct {
	CoinID                int64  `json:"coin_id" query:"coin_id"`
	CoinSymbol            string `json:"coin_symbol" query:"coin_symbol"`
	DailyQuantity         string `json:"daily_quantity" query:"daily_quantity"`
	DailyExchangeQuantity string `json:"daily_exchange_quantity" query:"daily_exchange_quantity"`
}

////////////////////////////////////////

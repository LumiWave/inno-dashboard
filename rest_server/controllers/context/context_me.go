package context

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
)

////////////////////////////////////////

// /////// Me Point List
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
	AppID                     int64  `json:"app_id" query:"app_id"`
	PointID                   int64  `json:"point_id" query:"point_id"`
	Quantity                  int64  `json:"quantity"`
	TodayAcqQuantity          int64  `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity         int64  `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	TodayAcqExchangeQuantity  int64  `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
	TodayCnsmExchangeQuantity int64  `json:"today_cnsm_exchange_quantity" query:"today_cnsm_exchange_quantity"`
	ResetDate                 string `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

// /////// Me Coin List
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
	CoinID                    int64     `json:"coin_id" query:"coin_id"`
	BaseCoinID                int64     `json:"base_coin_id" query:"base_coin_id"`
	CoinSymbol                string    `json:"coin_symbol" query:"coin_symbol"`
	WalletAddress             string    `json:"wallet_address" query:"wallet_address"`
	Quantity                  float64   `json:"quantity" query:"quantity"`
	TodayAcqQuantity          float64   `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity         float64   `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	TodayAcqExchangeQuantity  float64   `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
	TodayCnsmExchangeQuantity float64   `json:"today_cnsm_exchange_quantity" query:"today_cnsm_exchange_quantity"`
	ResetDate                 time.Time `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

// /////// Member
type Member struct {
	MUID       int64 `json:"mu_id"`
	AppID      int64 `json:"app_id"`
	DatabaseID int64 `json:"database_id"`
}

////////////////////////////////////////

// /////// otp : qrcode 용 uri 조회
type MeOtpUri struct {
	OtpUri string `json:"otp_uri"`
}

////////////////////////////////////////

// /////// otp : qrcode 용 uri 조회
type MeOtpVerify struct {
	OtpCode string `json:"otp_code" query:"otp_code"`
}

func NewMeOtpVerify() *MeOtpVerify {
	return new(MeOtpVerify)
}

func (o *MeOtpVerify) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	return nil
}

////////////////////////////////////////

// /////// coin mainnet 보정
type CoinReload struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func NewCoinReload() *CoinReload {
	return new(CoinReload)
}

func (o *CoinReload) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

////////////////////////////////////////

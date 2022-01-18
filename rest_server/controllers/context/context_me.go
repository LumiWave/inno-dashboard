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
	CoinID        int64     `json:"coin_id" query:"coin_id"`
	WalletAddress string    `json:"wallet_address" query:"wallet_address"`
	Quantity      string    `json:"quantity" query:"quantity"`
	DailyQuantity string    `json:"daily_quantity" query:"daily_quantity"`
	ResetDate     time.Time `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

///////// Me Point List
type ReqMePoint struct {
	AppId int64 `json:"app_id" query:"app_id"`
}

func (o *ReqMePoint) CheckValidate() *base.BaseResponse {
	if o.AppId == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AppID_Empty)
	}
	return nil
}

type MePoint struct {
	PointID               int64  `json:"point_id" query:"point_id"`
	PointName             string `json:"point_name" query:"point_name"`
	DailyQuantity         int64  `json:"daily_quantity" query:"daily_quantity"`
	DailyExchangeQuantity int64  `json:"daily_exchange_quantity" query:"daily_exchange_quantity"`
}

type MePointList struct {
	MePointList []MePoint
}

////////////////////////////////////////

///////// Me Coin List
type ReqMeCoin struct {
	AppId int64 `json:"app_id" query:"app_id"`
}

func (o *ReqMeCoin) CheckValidate() *base.BaseResponse {
	if o.AppId == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AppID_Empty)
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

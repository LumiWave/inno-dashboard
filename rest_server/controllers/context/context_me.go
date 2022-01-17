package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Wallet Info
type WalletInfo struct {
	CoinID        int64  `json:"coin_id" query:"coin_id"`
	CoinSymbol    string `json:"coin_symbol" query:"coin_symbol"`
	WalletAddress string `json:"wallet_address" query:"wallet_address"`
	CoinQuantity  string `json:"coin_quantity" query:"coin_quantity"`
}

type Wallets struct {
	Wallets []WalletInfo `json:"wallets" query:"wallets"`
}

////////////////////////////////////////

///////// Me Point List
type ReqMePoint struct {
	AppId int64 `json:"app_id" query:"app_id"`
}

func (o *ReqMePoint) CheckValidate() *base.BaseResponse {
	if o.AppId == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_PointList_Empty)
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
		return base.MakeBaseResponse(resultcode.Result_Get_Me_CoinList_Empty)
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

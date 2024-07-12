package context

import (
	"github.com/LumiWave/baseapp/base"
)

// /////// Point Info
type PointInfo struct {
	PointId   int64  `json:"point_id,omitempty"`
	PointName string `json:"point_name,omitempty"`
	IconUrl   string `json:"icon_url,omitempty"`
	//MinExchangeQuantity             int64   `json:"minimum_exchange_quantity,omitempty"`
	//ExchangeRatio                   float64 `json:"exchange_ratio,omitempty"`
	DaliyLimitAcqQuantity         int64 `json:"daliy_limit_acq_quantity,omitempty"`          // 게임을 통한 일일 제한 획득량
	DailyLimitExchangeAcqQuantity int64 `json:"daily_limit_exchange_acq_quantity,omitempty"` // 일일 제한 전환 획득량
}

type PointList struct {
	Points []*PointInfo `json:"points"`
}

type AppPointInfo struct {
	AppId            int64        `json:"app_id,omitempty"`
	AppName          string       `json:"app_name,omitempty"`
	IconUrl          string       `json:"icon_url"`
	GooglePlayPath   string       `json:"google_play_path"`
	AppleStorePath   string       `json:"apple_store_path"`
	BrandingPagePath string       `json:"branding_page_path"`
	Points           []*PointInfo `json:"points"`
}

type AppPoints struct {
	Apps []*AppPointInfo `json:"apps"`
}

////////////////////////////////////////

// /////// Coin Info
type CoinInfo struct {
	BaseCoinID                    int64   `json:"base_coin_id"`
	CoinId                        int64   `json:"coin_id,omitempty"`
	CoinName                      string  `json:"coin_name"`
	CoinSymbol                    string  `json:"coin_symbol,omitempty"`
	ContractAddress               string  `json:"contract_address,omitempty"`
	Decimal                       int64   `json:"dicimal"`
	ExplorePath                   string  `json:"explore_path"`
	IconUrl                       string  `json:"icon_url,omitempty"`
	DailyLimitExchangeAcqQuantity float64 `json:"daily_limit_exchange_acq_quantity"` // 일일 제한 전환 획득량
	ExchangeFees                  float64 `json:"exchange_fees"`
	IsRechargeable                bool    `json:"is_rechargeable"`
	RechargeURL                   string  `json:"recharge_url"`
	CustomProperties              string  `json:"custom_properties"`
}

type CoinList struct {
	Coins []*CoinInfo `json:"coins"`
}

////////////////////////////////////////

// /////// BaseCoinInfo
type BaseCoinInfo struct {
	BaseCoinID         int64   `json:"base_coin_id"`
	BaseCoinName       string  `json:"base_coin_name"`
	BaseCoinSymbol     string  `json:"base_coin_symbol"`
	IsUsedParentWallet bool    `json:"is_used_parent_wallet"`
	AllowWalletTypes   []int64 `json:"allow_wallet_types"`
}

type BaseCoinList struct {
	Coins []*BaseCoinInfo `json:"base_coins"`
}

// //////////////////////////////////////
type WalletTypeList struct {
	WalletTypes []*WalletType `json:"wallet_types"`
}
type WalletType struct {
	WalletName   string `json:"wallet_name"`
	WalletTypeID int64  `json:"wallet_type_id"`
}

// /////// AppPointDailyInfo
type ReqAppPointDaily struct {
	AppID   int64 `json:"app_id" query:"app_id"`
	PointID int64 `json:"point_id" query:"point_id"`
}

func (o *ReqAppPointDaily) CheckValidate() *base.BaseResponse {
	return nil
}

type ResPointDaily struct {
	PointID                  int64 `json:"point_id" query:"point_id"`
	TodayAcqQuantity         int64 `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayAcqExchangeQuantity int64 `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
}

type ResAppPointDaily struct {
	AppID          int64            `json:"app_id" query:"app_id"`
	ResPointDailys []*ResPointDaily `json:"point_dailys"`
}

////////////////////////////////////////

// /////// AppCoinDailyInfo
type ReqAppCoinDaily struct {
	CoinID int64 `json:"coin_id" query:"coin_id"`
}

func (o *ReqAppCoinDaily) CheckValidate() *base.BaseResponse {
	return nil
}

type ResAppCoinDaily struct {
	CoinID                   int64   `json:"coin_id" query:"coin_id"`
	TodayAcqQuantity         float64 `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayAcqExchangeQuantity float64 `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
}

////////////////////////////////////////

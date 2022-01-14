package context

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

type MePoint struct {
	PointID       int64  `json:"point_id" query:"point_id"`
	PointName     string `json:"point_name" query:"point_name"`
	TotalQuantity int64  `json:"total_quantity" query:"total_quantity"`
	TodayQuantity int64  `json:"today_quantity" query:"today_quantity"`
}

type MePointList struct {
	MePointList []MePoint
}

////////////////////////////////////////

///////// Me Coin List

type MeCoin struct {
	CoinID        int64  `json:"coin_id" query:"coin_id"`
	CoinName      string `json:"coin_name" query:"coin_name"`
	TotalQuantity int64  `json:"total_quantity" query:"total_quantity"`
	TodayQuantity int64  `json:"today_quantity" query:"today_quantity"`
}

type MeCoinList struct {
	MeCoinList []MeCoin
}

////////////////////////////////////////

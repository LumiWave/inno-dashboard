package context

type Meta struct {
	PointList
	AppPoints
	BaseCoinList
	CoinList
	Swapable []*Swapable `json:"swapable"`
	WalletTypeList
}

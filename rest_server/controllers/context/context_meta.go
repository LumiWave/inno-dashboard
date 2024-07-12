package context

type SwapAble struct {
	SwapAbleP2C any `json:"p2c"`
	SwapAbleC2P any `json:"c2p"`
	SwapAbleC2C any `json:"c2c"`
}

type Meta struct {
	PointList
	AppPoints
	BaseCoinList
	CoinList
	SwapAble `json:"swapable"`
	WalletTypeList
}

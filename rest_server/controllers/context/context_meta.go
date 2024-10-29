package context

type SwapAble struct {
	SwapAbleP2C any   `json:"p2c"`
	SwapAbleC2P any   `json:"c2p"`
	SwapAbleC2C any   `json:"c2c"`
	SwapAbleP2P any   `json:"p2p"`
	ExpireCycle int64 `json:"expire_cycle"`
}

type Meta struct {
	PointList
	AppPoints
	BaseCoinList
	CoinList
	SwapAble `json:"swapable"`
	WalletTypeList
}

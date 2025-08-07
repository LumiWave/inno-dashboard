package context

type SwapAble struct {
	SwapAbleP2C any   `json:"p2c"`
	SwapAbleC2P any   `json:"c2p"`
	SwapAbleC2C any   `json:"c2c"`
	SwapAbleP2P any   `json:"p2p"`
	ExpireCycle int64 `json:"expire_cycle"`
}

type SwapTier struct {
	SwapP2CTier any `json:"p2c_tier"`
}

type SwapTierCondition struct {
	SwapP2CTierCondition any `json:"p2c_tier_condition"`
}

type Meta struct {
	PointList
	AppPoints
	BaseCoinList
	CoinList
	SwapAble          `json:"swapable"`
	SwapTier          `json:"swap_tier"`
	SwapTierCondition `json:"swap_tier_condition"`
	WalletTypeList
}

package context

type Meta struct {
	PointList
	AppPoints
	CoinList

	Swapable []*Swapable `json:"swapable"`
}

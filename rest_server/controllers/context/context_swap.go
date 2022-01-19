package context

///////// Swap Info 전체 포인트, 코인 정보 리스트 조회
type Swapable struct {
	AppID   int64 `json:"app_id"`
	CoinID  int64 `json:"coin_id"`
	PointID int64 `json:"point_id"`
}

type SwapList struct {
	AppPoints
	CoinList

	Swapable []*Swapable `json:"swapable"`
}

////////////////////////////////////////

///////// Swap 가능 정보 조회

type ReqSwapEnable struct {
	FromType     string `json:"from_type" query:"from_type"`         // 출발지 타입 "point" | "coin"
	FromID       string `json:"from_id" query:"from_id"`             // 출발지 id
	FromQuantity string `json:"from_quantity" query:"from_quantity"` // 양
}

type RespSwapEnable struct {
	MinimumReceived      int64   `json:"minimum_received,omitempty"`       // 최소 스왑 가능 금액
	PriceImpact          float64 `json:"price_impact,omitempty"`           // 가격 변동률
	LiquidityProviderFee float64 `json:"liquidity_provider_fee,omitempty"` // 수수료

	ToType     string `json:"to_type,omitempty"`     // 도착지 타입 point or coin
	ToID       int64  `json:"to_id,omitempty"`       // 도착지 id
	ToQuantity string `json:"to_quantity,omitempty"` // 양
}

////////////////////////////////////////

///////// Swap 처리

type BaseSwapInfo struct {
	Type          string `json:"type"`                     // point or coin
	ID            int64  `json:"id"`                       // point_id or coin_id
	PointQuantity int64  `json:"point_quantity,omitempty"` // 포인트 양
	CoinQuantity  string `json:"coin_quantity,omitempty"`  // 코인 양
}

type FromInfo struct {
	BaseSwapInfo
}

type ToInfo struct {
	BaseSwapInfo
}

type SwapInfo struct {
	FromInfo `json:"from_info"` // 출발지
	ToInfo   `json:"to_info"`   // 도착지
}

////////////////////////////////////////

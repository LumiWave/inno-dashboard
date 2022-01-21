package point_manager_server

type Common struct {
	Return  int    `json:"return"`
	Message string `json:"message"`
}

///////// point list 응답
type Point struct {
	PointID          int64  `json:"point_id"`
	Quantity         int64  `json:"quantity"`
	TodayQuantity    int64  `json:"today_quantity"`
	ResetDate        string `json:"reset_date"`
	PreviousQuantity int64  `json:"-"`
}

type MePointValue struct {
	MyUUID     string   `json:"-"`
	DatabaseID int64    `json:"database_id"`
	MUID       int64    `json:"mu_id"`
	Points     []*Point `json:"points"`
}

type MePointInfo struct {
	Common
	MePointValue `json:"value"`
}

////////////////////////////////////////

///////// swap 요청
type SwapPoint struct {
	MUID                  int64 `json:"mu_id"`                   // 검색
	AppID                 int64 `json:"app_id"`                  // 요청 인자
	DatabaseID            int64 `json:"database_id"`             // 검색
	PointID               int64 `json:"point_id"`                // 요청 인자
	PreviousPointQuantity int64 `json:"previous_point_quantity"` // point manager server 이관
	AdjustPointQuantity   int64 `json:"adjust_point_quantity"`   // 요청 인자
	PointQuantity         int64 `json:"point_quantity"`          // point manager server 이관
}

type SwapCoin struct {
	CoinID               int64   `json:"coin_id"`                // 요청 인자
	WalletAddress        string  `json:"walletaddress"`          // 검색
	PreviousCoinQuantity float64 `json:"previous_coin_quantity"` // 검색
	AdjustCoinQuantity   float64 `json:"adjust_coin_quantity"`   // 요청 인자
	CoinQuantity         float64 `json:"coin_quantity"`          // 검색
}

type ReqSwapInfo struct {
	AUID int64 `json:"au_id"`

	SwapPoint `json:"point"`
	SwapCoin  `json:"coin"`

	LogID   int64 `json:"log_id"`   // 2: 전환
	EventID int64 `json:"event_id"` // 3: point->coin,  4: coin->point
}

type ResSwapInfo struct {
	Common
	ReqSwapInfo `json:"value"`
}

////////////////////////////////////////

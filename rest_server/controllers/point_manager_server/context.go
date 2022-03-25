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
	BaseCoinID           int64   `json:"base_coin_id"`           // 요청 인자
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

///////// 부모지갑에서 coin 전송 요청
type ReqCoinTransferFromParentWallet struct {
	AUID       int64   `json:"au_id" url:"au_id"`             // 계정의 UID (Access Token에서 가져옴)
	CoinID     int64   `json:"coin_id" url:"coin_id"`         // 코인 ID
	CoinSymbol string  `json:"coin_symbol" url:"coin_symbol"` // 코인 심볼
	ToAddress  string  `json:"to_address" url:"to_address"`   // 보낼 지갑 주소
	Quantity   float64 `json:"quantity" url:"quantity"`       // 보낼 코인량

	// internal used
	TransferFee   float64 `json:"transfer_fee" url:"transfer_fee"`     // 전송 수수료
	TotalQuantity float64 `json:"total_quantity" url:"total_quantity"` // 보낼 코인량 + 전송 수수료
	ReqId         string  `json:"reqid"`
	TransactionId string  `json:"transaction_id"`
}

type ResCoinTransferFromParentWallet struct {
	Common
	Value ReqCoinTransferFromParentWallet `json:"value"`
}

////////////////////////////////////////

///////// 특정지갑에서 coin 전송 요청
type ReqCoinTransferFromUserWallet struct {
	AUID           int64   `json:"au_id" url:"au_id"`                       // 계정의 UID (Access Token에서 가져옴)
	CoinID         int64   `json:"coin_id" url:"coin_id"`                   // 코인 ID
	CoinSymbol     string  `json:"coin_symbol" url:"coin_symbol"`           // 코인 심볼
	BaseCoinSymbol string  `json:"base_coin_symbol" url:"base_coin_symbol"` // base 코인 심볼
	FromAddress    string  `json:"from_address" url:"from_address"`         // 보내는 지갑 주소
	ToAddress      string  `json:"to_address" url:"to_address"`             // 보낼 지갑 주소
	Quantity       float64 `json:"quantity" url:"quantity"`                 // 보낼 코인량

	// internal used
	TransferFee   float64 `json:"transfer_fee" url:"transfer_fee"`     // 전송 수수료
	TotalQuantity float64 `json:"total_quantity" url:"total_quantity"` // 보낼 코인량 + 전송 수수료
	ReqId         string  `json:"reqid"`
	TransactionId string  `json:"transaction_id"`
}
type ResCoinTransferFromUserWallet struct {
	Common
	Value ReqCoinTransferFromUserWallet `json:"value"`
}

////////////////////////////////////////

///////// 코인 가스비 요청
type ReqCoinFee struct {
	Symbol string `query:"symbol"`
}

type ResCoinFeeInfoValue struct {
	Fast    string `json:"fast"`
	Slow    string `json:"slow"`
	Average string `json:"average"`
	BaseFee string `json:"basefee"`
	Fastest string `json:"fastest"`
}

type ResCoinFeeInfo struct {
	Common
	ResCoinFeeInfoValue `json:"value"`
}

////////////////////////////////////////

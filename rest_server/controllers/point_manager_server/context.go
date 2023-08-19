package point_manager_server

import "time"

type Common struct {
	Return  int    `json:"return"`
	Message string `json:"message"`
}

// /////// point list 응답
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

// /////// swap 요청
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
	CoinID             int64   `json:"coin_id"` // 요청 인자
	CoinSymbol         string  `json:"coin_symbol"`
	BaseCoinID         int64   `json:"base_coin_id"` // 요청 인자
	BaseCoinSymbol     string  `json:"base_coin_symbol"`
	WalletAddress      string  `json:"walletaddress"`
	AdjustCoinQuantity float64 `json:"adjust_coin_quantity"` // 요청 인자
}

type ReqSwapInfo struct {
	AUID int64 `json:"au_id"`

	SwapPoint `json:"point"`
	SwapCoin  `json:"coin"`

	TxType int64 `json:"tx_type"` // 3: point->coin,  4: coin->point

	SwapFee           float64 `json:"swap_fee"` // point->coin 시 전환시 부모지갑에 전송될 코인량 coin->point는 0 고정
	SwapWalletAddress string  `json:"swap_fee_to_wallet"`
	InnoUID           string  `json:"inno_uid"`
	TxID              int64   `json:"tx_id"`
	CreateAt          int64   `json:"create_at"`
}

type ResSwapInfo struct {
	Common
	ReqSwapInfo `json:"value"`
}

////////////////////////////////////////

// swap 상태 갱신 요청
type ReqSwapGasFee struct {
	TxStatus          int64  `json:"tx_status"`
	TxHash            string `json:"tx_hash"`
	FromWalletAddress string `json:"from_wallet_address"`
}

type ResSwapGasFee struct {
	Common
}

////////////////////////////////////////

// /////// 부모지갑에서 coin 전송 요청
type ReqCoinTransferFromParentWallet struct {
	AUID       int64   `json:"au_id" url:"au_id"`             // 계정의 UID (Access Token에서 가져옴)
	CoinID     int64   `json:"coin_id" url:"coin_id"`         // 코인 ID
	CoinSymbol string  `json:"coin_symbol" url:"coin_symbol"` // 코인 심볼
	ToAddress  string  `json:"to_address" url:"to_address"`   // 보낼 지갑 주소
	Quantity   float64 `json:"quantity" url:"quantity"`       // 보낼 코인량

	// internal used
	ReqId         string `json:"reqid"`
	TransactionId string `json:"transaction_id"`
}

type ResCoinTransferFromParentWallet struct {
	Common
	Value ReqCoinTransferFromParentWallet `json:"value"`
}

////////////////////////////////////////

// /////// 특정지갑에서 coin 전송 요청
type ReqCoinTransferFromUserWallet struct {
	AUID           int64   `json:"au_id" url:"au_id"`                       // 계정의 UID (Access Token에서 가져옴)
	CoinID         int64   `json:"coin_id" url:"coin_id"`                   // 코인 ID
	CoinSymbol     string  `json:"coin_symbol" url:"coin_symbol"`           // 코인 심볼
	BaseCoinSymbol string  `json:"base_coin_symbol" url:"base_coin_symbol"` // base 코인 심볼
	FromAddress    string  `json:"from_address" url:"from_address"`         // 보내는 지갑 주소
	ToAddress      string  `json:"to_address" url:"to_address"`             // 보낼 지갑 주소
	Quantity       float64 `json:"quantity" url:"quantity"`                 // 보낼 코인량

	// internal used
	ReqId         string `json:"reqid"`
	TransactionId string `json:"transaction_id"`
}
type ResCoinTransferFromUserWallet struct {
	Common
	Value ReqCoinTransferFromUserWallet `json:"value"`
}

////////////////////////////////////////

// /////// 코인 가스비 요청
type ReqCoinFee struct {
	Symbol string `query:"symbol"`
}

type ResCoinFeeInfoValue struct {
	GasPrice string `json:"gas_price"`
	Decimal  int64  `json:"decimal"`
}

type ResCoinFeeInfo struct {
	Common
	ResCoinFeeInfoValue `json:"value"`
}

////////////////////////////////////////

// /////// 코인 코인 mainnet reload
type CoinReload struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

type ResCoinReloadValue struct {
	CoinID                    int64     `json:"coin_id"`
	BaseCoinID                int64     `json:"base_coin_id"`
	WalletAddress             string    `json:"wallet_address"`
	Quantity                  float64   `json:"quantity"`
	TodayAcqQuantity          float64   `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity         float64   `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	TodayAcqExchangeQuantity  float64   `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
	TodayCnsmExchangeQuantity float64   `json:"today_cnsm_exchange_quantity" query:"today_cnsm_exchange_quantity"`
	ResetDate                 time.Time `json:"reset_date" query:"reset_date"`
}

type ResCoinReload struct {
	Common
	Value []ResCoinReloadValue `json:"value"`
}

////////////////////////////////////////

// /////// 지갑 잔액
type ReqBalance struct {
	Symbol  string `query:"symbol"`
	Address string `query:"address"`
}

type ResBalance struct {
	Common
	Value struct {
		Balance string `json:"balance"`
		Address string `json:"address"`
		Decimal int64  `json:"decimal"`
	} `json:"value"`
}

////////////////////////////////////////

// /////// 전체 지갑 잔액
type ReqBalanceAll struct {
	AUID int64 `query:"au_id"`
}

type Balance struct {
	CoinID     int64  `json:"coin_id"`
	BaseCoinID int64  `json:"base_coin_id"`
	Symbol     string `json:"symbol"`
	Balance    string `json:"balance"`
	Address    string `json:"address"`
	Decimal    int64  `json:"decimal"`
}
type ResBalanceAll struct {
	Common
	Value struct {
		Balances []*Balance `json:"balances"`
	} `json:"value"`
}

////////////////////////////////////////

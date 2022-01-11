package context

// 가격 정보
type PriceInfo struct {
	OpeningPrice float64 `json:"opening_price,omitempty"`
	HighPrice    float64 `json:"high_price,omitempty"`
	LowPrice     float64 `json:"low_price,omitempty"`
	TradePrice   float64 `json:"trade_price,omitempty"`
}

// 시세 정보
type QuoteInfo struct {
	StartDate string  `json:"start_date,omitempty"`
	EndDate   string  `json:"end_date,omitempty"`
	Price     float64 `json:"price,omitempty"`
	Prepare   float64 `json:"prepare,omitempty"`
}

// 시세 History
type QuoteHistory struct {
	PageInfoResponse `json:"page_info"`
	CoinSymbol       string      `json:"coin_symbol,omitempty"`
	Candle           string      `json:"candle,omitempty"`
	ItemCount        int64       `json:"item_count,omitempty"`
	Quotes           []QuoteInfo `json:"quotes"`
}

// 유동량
type LiquidityInfo struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Liquidity int64  `json:"liquidity,omitempty"`
	Prepare   int64  `json:"prepare,omitempty"`
}

// 유동량 History
type LiquidityHistory struct {
	PageInfoResponse `json:"page_info"`
	ItemCount        int64  `json:"item_count,omitempty"`
	Candle           string `json:"candle,omitempty"`
	Liquiditys       []LiquidityInfo
}

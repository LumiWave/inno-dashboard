package context

///////// Coin Transfer
type CoinTransfer struct {
	WalletAddress string `json:"wallet_address,omitempty"`
	CoinSymbol    string `json:"coin_symbol,omitempty"`
	Quantity      string `json:"quantity,omitempty"`
}

////////////////////////////////////////

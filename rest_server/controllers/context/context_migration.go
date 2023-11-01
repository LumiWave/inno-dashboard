package context

// 마이그레이션 데이터
type MIGCoin struct {
	CoinID   int64   `json:"coin_id"`
	Quantity float64 `json:"quantity"`

	//메모리 보관용
	WalletAddress string `json:"wallet"`
	Ts            int64  `json:"ts"`
}

type MIGNFT struct {
	NFTPackID int64 `json:"nft_pack_id"`
	CoinID    int64 `json:"coin_id"`
	NFTID     int64 `json:"nft_id"`
}

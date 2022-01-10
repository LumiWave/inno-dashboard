package context

type PointInfo struct {
	PointId   int64  `json:"point_id,omitempty"`
	PointName string `json:"point_name,omitempty"`
	IconUrl   string `json:"icon_url,omitempty"`
}

type PointList struct {
	Points []PointInfo `json:"points"`
}

type AppInfo struct {
	AppId   int64  `json:"app_id,omitempty"`
	AppName string `json:"app_name,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
	PointId int64  `json:"point_id,omitempty"`
}

type Apps struct {
	Apps []AppInfo `json:"apps"`
}

type CoinInfo struct {
	CoinId          int64  `json:"coin_id,omitempty"`
	CoinSymbol      string `json:"coin_symbol,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	IconUrl         string `json:"icon_url,omitempty"`
}

type Coins struct {
	Coins []CoinInfo `json:"coins"`
}

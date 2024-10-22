package model

import "fmt"

const (
	PubSub      = "pubsub"
	InternalCmd = "internal_cmd"
)

const (
	PubSub_cmd_healthcheck           = "HealthCheck"
	PubSub_type_maintenance          = "Maintenance"
	PubSub_type_Swap                 = "Swap"
	PubSub_type_CoinTransferExternal = "CoinTransferExternal"
	PubSub_type_meta_refresh         = "MetaRefresh"
)

const (
	Webhook_type_meta_refresh = iota
)

type PSHeader struct {
	Type string `json:"type"`
}

type PSHealthCheck struct {
	PSHeader
	Value struct {
		Timestamp int64 `json:"ts"`
	} `json:"value"`
}

type PSMaintenance struct {
	PSHeader
	Value struct {
		Enable    bool   `json:"enable"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	} `json:"value"`
}

type PSSwap struct {
	PSHeader
	Value struct {
		ToCoinEnable  bool `json:"to_coin_enable"`
		ToPointEnable bool `json:"to_point_enable"`
		ToC2CEnable   bool `json:"to_c2c_enable"`
	} `json:"value"`
}

type PSCoinTransferExternal struct {
	PSHeader
	Value struct {
		Enable bool `json:"enable"`
	} `json:"value"`
}

type PSMetaRefresh struct {
	PSHeader
	Value struct {
		Enable bool `json:"enable"`
	} `json:"value"`
}

func MakePubSubKey(val string) string {
	return fmt.Sprintf("%s:%s", PubSub, val)
}

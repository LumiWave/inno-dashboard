package inno_market

import (
	"github.com/LumiWave/baseInnoClient/context"
	"github.com/LumiWave/baseInnoClient/market"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
)

var gMarketServer *market.Server

func GetInstance() *market.Server {
	return gMarketServer
}

func InitMarketServer(conf *config.ServerConfig) error {
	MarketServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.InnoMarket.InternalpiDomain,
			ExtHostUri: conf.InnoMarket.ExternalpiDomain,
			IntVer:     conf.InnoMarket.InternalVer,
			ExtVer:     conf.InnoMarket.ExternalVer,
		},
	}

	gMarketServer = market.NewServerInfo(MarketServerInfo)
	return nil
}

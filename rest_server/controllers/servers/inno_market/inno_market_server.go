package inno_market

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/market"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
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

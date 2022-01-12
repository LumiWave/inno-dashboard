package upbit

import (
	"strings"

	baseupbit "github.com/ONBUFF-IP-TOKEN/baseMarket/market/upbit"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

const (
	KRW_BTC = "KRW-BTC"
)

var gUpbitInfo *baseupbit.UpbitInfo

func InitUpbitInfo() {
	gUpbitInfo = &baseupbit.UpbitInfo{
		HostInfo: baseupbit.HostInfo{
			HostUri: "https://api.upbit.com",
			Ver:     "/v1",
		},
	}
}

func GetUpbitInfo() *baseupbit.UpbitInfo {
	return gUpbitInfo
}

func GetQuoteTicker(coinSymbol string) (*context.PriceInfo, error) {
	var priceInfo *context.PriceInfo

	if resp, err := baseupbit.GetQuoteTicker(gUpbitInfo, coinSymbol+","+KRW_BTC); err != nil {
		return nil, err
	} else {
		// BTC-ONIT to KRW
		onitPrice := float64(0)
		btcPrice := float64(0)
		var priceTimeStamp int64
		for _, ticker := range *resp {
			if strings.EqualFold(ticker.Market, coinSymbol) {
				onitPrice = ticker.TradePrice
				priceTimeStamp = ticker.TradeTimeStamp
			} else if strings.EqualFold(ticker.Market, KRW_BTC) {
				btcPrice = ticker.TradePrice
			}
		}

		priceInfo = &context.PriceInfo{
			CoinSymbol:     coinSymbol,
			ONITPrice:      onitPrice,
			KRWPrice:       onitPrice * btcPrice,
			PriceTimeStamp: priceTimeStamp,
		}
	}

	return priceInfo, nil
}

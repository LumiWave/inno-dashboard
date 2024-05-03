package commonapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

func GetMeta(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// tmpCoinList := context.CoinList{}
	// for _, coin := range model.GetDB().Coins.Coins {
	// 	if coin.CoinSymbol == "MATIC" {
	// 		continue
	// 	}

	// 	tmpCoinList.Coins = append(tmpCoinList.Coins, coin)
	// }

	swapList := context.Meta{
		PointList: model.GetDB().ScanPoints,
		AppPoints: model.GetDB().AppPoints,
		CoinList:  model.GetDB().Coins,
		//CoinList:     tmpCoinList,
		Swapable:     model.GetDB().SwapAble,
		BaseCoinList: model.GetDB().BaseCoins,
	}

	resp.Value = swapList

	return c.JSON(http.StatusOK, resp)
}

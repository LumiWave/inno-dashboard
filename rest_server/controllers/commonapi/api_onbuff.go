package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/upbit"
	"github.com/labstack/echo"
)

//  현재 시세 조회
func GetCoinPrice(c echo.Context, reqPriceInfo *context.ReqPriceInfo) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if priceInfo, err := upbit.GetQuoteTicker(reqPriceInfo.CoinSymbol); err != nil {
		resp.SetReturn(resultcode.Result_Upbit_TickerMarkets)
	} else {
		resp.Value = priceInfo
	}

	return c.JSON(http.StatusOK, resp)
}

// 시세 history 조회
func GetCoinHistoryPrice(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 유동량 조회
func GetCoinHistoryLiquidity(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

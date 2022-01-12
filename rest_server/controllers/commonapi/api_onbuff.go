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

// 시세 history 조회 (Minutes)
func GetCoinCandleMinutes(c echo.Context, reqcandle *context.ReqCandleMinutes) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if value, err := upbit.GetCandleMinutes(reqcandle); err != nil {
		resp.SetReturn(resultcode.Result_Upbit_CandleDays)
	} else {
		resp.Value = value
	}

	return c.JSON(http.StatusOK, resp)
}

// 시세 history 조회 (Days)
func GetCoinCandleDays(c echo.Context, reqcandle *context.ReqCandleDays) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if value, err := upbit.GetCandleDays(reqcandle); err != nil {
		resp.SetReturn(resultcode.Result_Upbit_CandleDays)
	} else {
		resp.Value = value
	}

	return c.JSON(http.StatusOK, resp)
}

// 시세 history 조회 (Weeks)
func GetCoinCandleWeeks(c echo.Context, reqcandle *context.ReqCandleWeeks) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if value, err := upbit.GetCandleWeeks(reqcandle); err != nil {
		resp.SetReturn(resultcode.Result_Upbit_CandleWeeks)
	} else {
		resp.Value = value
	}

	return c.JSON(http.StatusOK, resp)
}

// 시세 history 조회 (Months)
func GetCoinCandleMonths(c echo.Context, reqcandle *context.ReqCandleMonths) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if value, err := upbit.GetCandleMonths(reqcandle); err != nil {
		resp.SetReturn(resultcode.Result_Upbit_CandleMonths)
	} else {
		resp.Value = value
	}
	return c.JSON(http.StatusOK, resp)
}

// 유동량 조회
func GetCoinHistoryLiquidity(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

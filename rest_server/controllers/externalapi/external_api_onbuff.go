package externalapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 현재 시세 조회
func (o *ExternalAPI) GetCoinPrice(c echo.Context) error {
	reqPriceInfo := new(context.ReqPriceInfo)

	// Request json 파싱
	if err := c.Bind(reqPriceInfo); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqPriceInfo.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.GetCoinPrice(c, reqPriceInfo)
}

// 시세 history 조회 (Minutes)
func (o *ExternalAPI) GetCoinCandleMinutes(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	reqCandleMinutes := new(context.ReqCandleMinutes)

	// Request json 파싱
	if err := c.Bind(reqCandleMinutes); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqCandleMinutes.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinCandleMinutes(c, reqCandleMinutes)
}

// 시세 history 조회 (Days)
func (o *ExternalAPI) GetCoinCandleDays(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	reqCandleDays := new(context.ReqCandleDays)

	// Request json 파싱
	if err := c.Bind(reqCandleDays); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqCandleDays.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinCandleDays(c, reqCandleDays)
}

// 시세 history 조회 (Weeks)
func (o *ExternalAPI) GetCoinCandleWeeks(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	reqCandleWeeks := new(context.ReqCandleWeeks)

	// Request json 파싱
	if err := c.Bind(reqCandleWeeks); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqCandleWeeks.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinCandleWeeks(c, reqCandleWeeks)
}

// 시세 history 조회 (Months)
func (o *ExternalAPI) GetCoinCandleMonths(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	reqCandleMonths := new(context.ReqCandleMonths)

	// Request json 파싱
	if err := c.Bind(reqCandleMonths); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqCandleMonths.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinCandleMonths(c, reqCandleMonths)
}

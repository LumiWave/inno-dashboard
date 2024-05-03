package externalapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 전체 포인트 리스트 조회
func (o *ExternalAPI) GetPointList(c echo.Context) error {
	return commonapi.GetPointList(c)
}

// 전체 App 리스트 조회
func (o *ExternalAPI) GetAppList(c echo.Context) error {
	return commonapi.GetAppList(c)
}

// 전체 코인 리스트 조회
func (o *ExternalAPI) GetCoinList(c echo.Context) error {
	return commonapi.GetCoinList(c)
}

// App 포인트 별 당일 누적/전환량 정보 조회
func (o *ExternalAPI) GetAppPoint(c echo.Context) error {
	reqAppPointDaily := new(context.ReqAppPointDaily)

	// Request json 파싱
	if err := c.Bind(reqAppPointDaily); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqAppPointDaily.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetAppPointAll(c, reqAppPointDaily)
}

// 코인 별 당일 누적/전환량 조회
func (o *ExternalAPI) GetAppCoin(c echo.Context) error {
	reqAppCoinDaily := new(context.ReqAppCoinDaily)

	// Request json 파싱
	if err := c.Bind(reqAppCoinDaily); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqAppCoinDaily.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.GetAppCoinDailyAll(c, reqAppCoinDaily)
}

// App 포인트 별 유동량 history 조회
func (o *ExternalAPI) GetAppPointHistory(c echo.Context) error {
	params := context.NewReqPointLiquidity()

	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("context bind error : %v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetAppPointHistory(c, params)
}

// 코인별 유동량 history 조회
func (o *ExternalAPI) GetAppCoinHistory(c echo.Context) error {
	params := context.NewReqCoinLiquidity()

	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("context bind error : %v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetAppCoinHistory(c, params)
}

// 하루 최대 수집 포인트 양 조회
func (o *ExternalAPI) GetAppPointMax(c echo.Context) error {
	return commonapi.GetAppPointMax(c)
}

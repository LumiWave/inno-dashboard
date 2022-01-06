package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
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
	return commonapi.GetAppPoint(c)
}

// App 포인트 별 유동량 history 조회
func (o *ExternalAPI) GetAppPointHistory(c echo.Context) error {
	return commonapi.GetAppPointHistory(c)
}

// 코인 별 당일 누적/전환량 조회
func (o *ExternalAPI) GetAppCoin(c echo.Context) error {
	return commonapi.GetAppCoin(c)
}

// 코인별 유동량 history 조회
func (o *ExternalAPI) GetAppCoinHistory(c echo.Context) error {
	return commonapi.GetAppCoinHistory(c)
}

// 하루 최대 수집 포인트 양 조회
func (o *ExternalAPI) GetAppPointMax(c echo.Context) error {
	return commonapi.GetAppPointMax(c)
}

package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

//  현재 시세 조회
func GetCoinPrice(c echo.Context) error {
	return commonapi.GetCoinPrice(c)
}

// 시세 history 조회
func GetCoinHistoryPrice(c echo.Context) error {
	return commonapi.GetCoinHistoryPrice(c)
}

// 유동량 조회
func GetCoinHistoryLiquidity(c echo.Context) error {
	return commonapi.GetCoinHistoryLiquidity(c)
}

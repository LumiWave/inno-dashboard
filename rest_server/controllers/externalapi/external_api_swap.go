package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 전체 포인트, 코인 정보 리스트 조회
func (o *ExternalAPI) GetSwapList(c echo.Context) error {
	return commonapi.GetSwapList(c)
}

// Swap 가능 정보 조회 (최소, 변동률, 수수료)
func (o *ExternalAPI) GetSwapEnable(c echo.Context) error {
	return commonapi.GetSwapEnable(c)
}

// Swap 처리
func (o *ExternalAPI) PostSwap(c echo.Context) error {
	return commonapi.PostSwap(c)
}

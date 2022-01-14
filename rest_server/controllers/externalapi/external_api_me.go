package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 지갑 정보 조회
func (o *ExternalAPI) GetMeWallets(c echo.Context) error {
	return commonapi.GetMeWallets(c)
}

// App 별 총/금일 누적 포인트 리스트 조회
func (o *ExternalAPI) GetMePointList(c echo.Context) error {
	return commonapi.GetMePointList(c)
}

// App 별 총/금일 누적 코인 리스트 조회
func (o *ExternalAPI) GetMeCoinList(c echo.Context) error {
	return commonapi.GetMeCoinList(c)
}

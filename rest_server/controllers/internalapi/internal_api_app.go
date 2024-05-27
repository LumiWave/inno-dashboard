package internalapi

import (
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 전체 포인트 리스트 DB reload
func (o *InternalAPI) PostReLoadPointList(c echo.Context) error {
	return commonapi.PostReLoadPointList(c)
}

// 전체 App 리스트 DB reload
func (o *InternalAPI) PostReLoadAppList(c echo.Context) error {
	return commonapi.PostReLoadAppList(c)
}

// 전체 코인 리스트 DB reload
func (o *InternalAPI) PostReloadCoinList(c echo.Context) error {
	return commonapi.PostReloadCoinList(c)
}

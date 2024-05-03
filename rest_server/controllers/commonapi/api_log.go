package commonapi

import (
	"net/http"

	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi/inner"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// App 포인트 별 유동량 history 조회
func GetAppPointHistory(c echo.Context, req *context.ReqPointLiquidity) error {
	resp := inner.GetPointHistory(req)
	return c.JSON(http.StatusOK, resp)
}

// 코인별 유동량 history 조회
func GetAppCoinHistory(c echo.Context, req *context.ReqCoinLiquidity) error {
	resp := inner.GetCoinHistory(req)
	return c.JSON(http.StatusOK, resp)
}

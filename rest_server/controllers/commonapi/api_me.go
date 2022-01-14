package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/labstack/echo"
)

// 지갑 정보 조회
func GetMeWallets(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 포인트 리스트 조회
func GetMePointList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 코인 리스트 조회
func GetMeCoinList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

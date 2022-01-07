package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/labstack/echo"
)

// 외부 지갑으로 코인 전송
func PostTransfer(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

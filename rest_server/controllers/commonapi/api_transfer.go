package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 외부 지갑으로 코인 전송
func PostTransfer(c echo.Context, reqCoinTransfer *context.ReqCoinTransfer) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// 필요한 정보를 모아서 point-manager "외부 지갑으로 토큰 전송" 요청

	return c.JSON(http.StatusOK, resp)
}

package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 외부 지갑으로 코인 전송
func (o *ExternalAPI) PostTransfer(c echo.Context) error {
	return commonapi.PostTransfer(c)
}

package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 공지 조회
func (o *ExternalAPI) GetNotice(c echo.Context) error {
	return commonapi.GetNotice(c)
}

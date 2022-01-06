package externalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

func (o *ExternalAPI) GetNotice(c echo.Context) error {
	return commonapi.GetNotice(c)
}

package internalapi

import (
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

func (o *InternalAPI) PostNotice(c echo.Context) error {
	return commonapi.PostNotice(c)
}

func (o *InternalAPI) PutNotice(c echo.Context) error {
	return commonapi.PutNotice(c)
}

func (o *InternalAPI) DeleteNotice(c echo.Context) error {
	return commonapi.DeleteNotice(c)
}

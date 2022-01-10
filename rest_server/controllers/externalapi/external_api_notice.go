package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 공지 조회
func (o *ExternalAPI) GetNotice(c echo.Context) error {
	pageInfo := new(context.PageInfo)

	// Request json 파싱
	if err := c.Bind(pageInfo); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := pageInfo.CheckValidate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetNotice(c)
}

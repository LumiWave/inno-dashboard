package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

//  현재 시세 조회
func (o *ExternalAPI) GetCoinPrice(c echo.Context) error {
	reqPriceInfo := new(context.ReqPriceInfo)

	// Request json 파싱
	if err := c.Bind(reqPriceInfo); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqPriceInfo.CheckValidate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.GetCoinPrice(c, reqPriceInfo)
}

// 시세 history 조회
func (o *ExternalAPI) GetCoinHistoryPrice(c echo.Context) error {
	return commonapi.GetCoinHistoryPrice(c)
}

// 유동량 조회
func (o *ExternalAPI) GetCoinHistoryLiquidity(c echo.Context) error {
	return commonapi.GetCoinHistoryLiquidity(c)
}

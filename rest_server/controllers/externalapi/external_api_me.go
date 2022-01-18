package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 지갑 정보 조회
func (o *ExternalAPI) GetMeWallets(c echo.Context) error {
	reqMeWallet := new(context.ReqMeWallet)

	// Request json 파싱
	if err := c.Bind(reqMeWallet); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqMeWallet.CheckValidate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMeWallets(c, reqMeWallet)
}

// App 별 총/금일 누적 포인트 리스트 조회
func (o *ExternalAPI) GetMePointList(c echo.Context) error {
	reqMePoint := new(context.ReqMePoint)

	// Request json 파싱
	if err := c.Bind(reqMePoint); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqMePoint.CheckValidate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMePointList(c, reqMePoint)
}

// App 별 총/금일 누적 코인 리스트 조회
func (o *ExternalAPI) GetMeCoinList(c echo.Context) error {
	reqMeCoin := new(context.ReqMeCoin)

	// Request json 파싱
	if err := c.Bind(reqMeCoin); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := reqMeCoin.CheckValidate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMeCoinList(c, reqMeCoin)
}

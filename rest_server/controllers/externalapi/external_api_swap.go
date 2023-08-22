package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 전체 포인트, 코인 정보 리스트 조회
func (o *ExternalAPI) GetSwapList(c echo.Context) error {
	return commonapi.GetSwapList(c)
}

// Swap 가능 정보 조회 (최소, 변동률, 수수료)
func (o *ExternalAPI) GetSwapEnable(c echo.Context) error {
	reqSwapEnable := new(context.ReqSwapEnable)

	// Request json 파싱
	if err := c.Bind(reqSwapEnable); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}
	return commonapi.GetSwapEnable(c, reqSwapEnable)
}

// Swap 처리
func (o *ExternalAPI) PostSwap(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqSwapInfo)

	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.PostSwap(ctx, params)
}

// Swap 을 위한 수수료 전송후 정보 수신
func (o *ExternalAPI) PutSwapGasFee(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqSwapGasFee)

	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.PutSwapGasFee(ctx, params)
}

// swap 진행 중인 정보가 있는지 확인
func (o *ExternalAPI) GetSwapInprogressNotExist(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewReqSwapIniprogress()

	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}
	return commonapi.GetSwapInprogressNotExist(ctx, params)
}

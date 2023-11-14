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
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqMeCoin)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMeWallets(c, params)
}

// App 별 총/금일 누적 포인트 리스트 조회
func (o *ExternalAPI) GetMePointList(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqMePoint)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMePointList(c, params)
}

// App 별 총/금일 누적 코인 리스트 조회
func (o *ExternalAPI) GetMeCoinList(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqMeCoin)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetMeCoinList(c, params)
}

// 내 sui 코인 보유 object id 리스트 조회
func (o *ExternalAPI) GetCoinObjects(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)

	params := context.NewReqCoinObjects()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	if err := params.CheckValidate(ctx); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinObjects(params, ctx)
}

// google otp : qrcode용 uri 조회
func (o *ExternalAPI) GetOtpUri(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	return commonapi.GetOtpUri(ctx)
}

func (o *ExternalAPI) GetOtpVerify(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)

	params := context.NewMeOtpVerify()
	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetOtpVerify(ctx, params)
}

func (o *ExternalAPI) PostCoinReload(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)

	params := context.NewCoinReload()
	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.PostCoinReload(ctx, params)
}

func (o *ExternalAPI) GetWalletRegist(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqGetWalletRegist)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetWalletRegist(ctx, params)
}

func (o *ExternalAPI) PostWalletRegist(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqPostWalletRegist)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.PostWalletRegist(ctx, params)
}

func (o *ExternalAPI) DeleteWalletRegist(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqDeleteWalletRegist)

	// Request json 파싱
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(ctx); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.DeleteWalletRegist(ctx, params)
}

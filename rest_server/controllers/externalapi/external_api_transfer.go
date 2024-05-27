package externalapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
)

// 외부 지갑으로 코인 전송
func (o *ExternalAPI) PostTransfer(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := new(context.ReqCoinTransfer)

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

	// 클라이언트로부터 전달받은 파라미터 정보 검증
	// 유저의 코인의 양이 가능한지,
	// 클라이언트가 전달해준 수수료가 맞는지 등.

	return commonapi.PostTransfer(ctx, params)
}

// 코인 외부 지갑으로 코인 전송 중인 상태 정보 요청
func (o *ExternalAPI) GetCoinTransferExistInProgress(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewGetCoinTransferExistInProgress()

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
	return commonapi.GetCoinTransferExistInProgress(ctx, params)
}

// 코인 외부 지갑 전송 중인 상태 정보 존재 하지 않는지 요청
func (o *ExternalAPI) GetCoinTransferNotExistInProgress(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewGetCoinTransferExistInProgress()

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
	return commonapi.GetCoinTransferNotExistInProgress(ctx, params)
}

func (o *ExternalAPI) GetCoinTransferFee(c echo.Context) error {
	params := context.NewGetCoinFee()
	// Request json 파싱
	if err := c.Bind(params); err != nil {
		log.Errorf("%v", err)
		return base.BaseJSONInternalServerError(c, err)
	}

	// 유효성 체크
	if err := params.CheckValidate(); err != nil {
		log.Errorf("%v", err)
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetCoinTransferFee(c, params)
}

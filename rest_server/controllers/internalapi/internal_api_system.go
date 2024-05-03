package internalapi

import (
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func (o *InternalAPI) GetNodeMetric(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)

	return commonapi.GetNodeMetric(ctx)
}

func (o *InternalAPI) PostPSMaintenance(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewPSMaintenance()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}
	return commonapi.PostPSMaintenance(ctx, params)
}

func (o *InternalAPI) PostPSSwap(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewPSSwap()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}
	return commonapi.PostPSSwap(ctx, params)
}

func (o *InternalAPI) PostPSCoinTransferExternal(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	params := context.NewPSCoinTransferExternal()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}
	return commonapi.PostPSCoinTransferExternal(ctx, params)
}

func (o *InternalAPI) PostPSMetaRefresh(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	return commonapi.PostPSMetaRefresh(ctx)
}

func (o *InternalAPI) GetPubsub(c echo.Context) error {
	ctx := base.GetContext(c).(*context.InnoDashboardContext)
	return commonapi.GetPubsub(ctx)
}

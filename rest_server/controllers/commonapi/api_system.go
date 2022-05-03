package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/schedule"
)

func GetNodeMetric(ctx *context.InnoDashboardContext) error {
	resp := new(base.BaseResponse)
	resp.Success()

	node := schedule.GetSystemMonitor().GetMetricInfo()
	resp.Value = node

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostPSMaintenance(ctx *context.InnoDashboardContext, req *context.PSMaintenance) error {
	resp := new(base.BaseResponse)
	resp.Success()

	msg := &model.PSMaintenance{
		PSHeader: model.PSHeader{
			Type: model.PubSub_type_maintenance,
		},
	}
	msg.Value.Enable = req.Enable

	if err := model.GetDB().PublishEvent(model.InternalCmd, msg); err != nil {
		log.Errorf("PublishEvent %v, type : %v, error : %v", model.InternalCmd, model.PubSub_type_maintenance, err)
		resp.SetReturn(resultcode.Result_PubSub_InternalErr)
	}
	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostPSSwap(ctx *context.InnoDashboardContext, req *context.PSSwap) error {
	resp := new(base.BaseResponse)
	resp.Success()

	msg := &model.PSSwap{
		PSHeader: model.PSHeader{
			Type: model.PubSub_type_Swap,
		},
	}
	msg.Value.ToCoinEnable = req.ToCoinEnable
	msg.Value.ToPointEnable = req.ToPointEnable

	if err := model.GetDB().PublishEvent(model.InternalCmd, msg); err != nil {
		log.Errorf("PublishEvent %v, type : %v, error : %v", model.InternalCmd, model.PubSub_type_Swap, err)
		resp.SetReturn(resultcode.Result_PubSub_InternalErr)
	}
	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostPSCoinTransferExternal(ctx *context.InnoDashboardContext, req *context.PSCoinTransferExternal) error {
	resp := new(base.BaseResponse)
	resp.Success()

	msg := &model.PSCoinTransferExternal{
		PSHeader: model.PSHeader{
			Type: model.PubSub_type_CoinTransferExternal,
		},
	}
	msg.Value.Enable = req.Enable

	if err := model.GetDB().PublishEvent(model.InternalCmd, msg); err != nil {
		log.Errorf("PublishEvent %v, type : %v, error : %v", model.InternalCmd, model.PubSub_type_CoinTransferExternal, err)
		resp.SetReturn(resultcode.Result_PubSub_InternalErr)
	}
	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostPSMetaRefresh(ctx *context.InnoDashboardContext) error {
	resp := new(base.BaseResponse)
	resp.Success()

	msg := &model.PSMetaRefresh{
		PSHeader: model.PSHeader{
			Type: model.PubSub_type_meta_refresh,
		},
	}

	if err := model.GetDB().PublishEvent(model.InternalCmd, msg); err != nil {
		log.Errorf("PublishEvent %v, type : %v, error : %v", model.InternalCmd, model.PubSub_type_meta_refresh, err)
		resp.SetReturn(resultcode.Result_PubSub_InternalErr)
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

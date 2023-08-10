package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
)

func GetWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletRegist, err := model.GetDB().USPAU_GetList_AccountWallets(params.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_AUID_Empty)
	} else {
		resp.Value = &context.ResWalletRegist{}
		//체크로직
		log.Debugf("%s", walletRegist)
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

package inner

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
)

func GetPointHistory(req *context.ReqPointLiquidity) *base.BaseResponse {
	resp := new(base.BaseResponse)
	resp.Success()

	// redis 에서 바로 읽어온다.
	key := model.MakeLogKeyOfPoint(req.AppID, req.PointID, req.Candle)
	if pointLiqs, err := model.GetDB().ZRangeLogOfPoint(key, 0, 100); err != nil {
		log.Errorf("ZRangeLogOfPoint error : %v", err)
		resp.SetReturn(resultcode.Result_Get_App_Point_Liquidity_Error)
	} else {
		resp.Value = pointLiqs
	}

	return resp
}

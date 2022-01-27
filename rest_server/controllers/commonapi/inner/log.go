package inner

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

func GetPointHistory(req *context.ReqPointLiquidity) *base.BaseResponse {
	resp := new(base.BaseResponse)
	resp.Success()

	// // redis 에 존재하는지 check

	// candleStr := context.CandleType[req.Candle]

	// key := model.MakeLogKeyOfPoint(req.AppID, req.Candle)
	// //_, _ = model.GetDB().ZRangeLogOfPoint(key, 0, 100)

	// if pointLiqs, err := model.GetDB().GetListPointLiquidity(candleStr, req); err != nil {
	// 	resp.SetReturn(resultcode.Result_Get_App_Point_Liquidity_Error)
	// } else {
	// 	resp.Value = pointLiqs

	// 	for _, pointLiq := range pointLiqs {
	// 		//model.GetDB().ZADDLogOfPoint(key, pointLiq.BaseDateToNumber, pointLiq)
	// 		model.GetDB().HSetLogOfPoint(key, strconv.FormatInt(pointLiq.BaseDateToNumber, 10), pointLiq)
	// 	}
	// 	_, _ = model.GetDB().ZRangeLogOfPoint(key, 0, 100)
	// }

	return resp
}

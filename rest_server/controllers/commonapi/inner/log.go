package inner

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/util"
)

func GetPointHistory(req *context.ReqPointLiquidity) *base.BaseResponse {
	resp := new(base.BaseResponse)
	resp.Success()

	// redis 에서 바로 읽어온다.
	key := model.MakeLogKeyOfPoint(req.AppID, req.PointID, req.Candle)
	if pointLiqs, err := model.GetDB().ZRevRangeLogOfPoint(
		key,
		util.MultiplyString(req.PageOffset, req.PageSize),
		util.MultiplyString(req.PageOffset, req.PageSize)+util.ParseInt(req.PageSize)-1); err != nil {
		log.Errorf("ZRevRangeLogOfPoint error : %v", err)
		resp.SetReturn(resultcode.Result_Get_App_Point_Liquidity_Error)
	} else {
		resp.Value = context.ResPointLiquidity{
			Count:           len(pointLiqs),
			PointLiquiditys: pointLiqs,
		}
	}

	return resp
}

func GetCoinHistory(req *context.ReqCoinLiquidity) *base.BaseResponse {
	resp := new(base.BaseResponse)
	resp.Success()

	// redis 에서 바로 읽어온다.
	key := model.MakeLogKeyOfCoin(req.CoinID, req.Candle)
	if coinLiqs, err := model.GetDB().ZRevRangeLogOfCoin(
		key,
		util.MultiplyString(req.PageOffset, req.PageSize),
		util.MultiplyString(req.PageOffset, req.PageSize)+util.ParseInt(req.PageSize)-1); err != nil {
		log.Errorf("ZRevRangeLogOfCoin error : %v", err)
		resp.SetReturn(resultcode.Result_Get_App_Coin_Liquidity_Error)
	} else {
		resp.Value = context.ResCoinLiquidity{
			Count:          len(coinLiqs),
			CoinLiquiditys: coinLiqs,
		}
	}

	return resp
}

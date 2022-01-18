package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 전체 포인트, 코인 정보 리스트 조회
func GetSwapList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	swapAble, err := model.GetDB().GetScanExchangeGoods()
	if err != nil {
		resp.SetReturn(resultcode.Result_Get_Swap_ExchangeGoods_Scan_Error)
		return c.JSON(http.StatusOK, resp)
	}

	swapList := context.SwapList{
		PointList: model.GetDB().ScanPoints,
		CoinList:  model.GetDB().Coins,
		Swapable:  swapAble,
	}

	resp.Value = swapList

	return c.JSON(http.StatusOK, resp)
}

// Swap 가능 정보 조회 (최소, 변동률, 수수료)
func GetSwapEnable(c echo.Context, reqSwapEnable *context.ReqSwapEnable) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// Swap 처리
func PostSwap(c echo.Context, swapInfo *context.SwapInfo) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 지갑 정보 조회
func GetMeWallets(c echo.Context, reqMeWallet *context.ReqMeWallet) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletList, err := model.GetDB().GetListAccountCoins(reqMeWallet.AUID); walletList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
	} else {
		resp.Value = walletList
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 포인트 리스트 조회
func GetMePointList(c echo.Context, reqMePoint *context.ReqMePoint) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if pointList, err := model.GetDB().GetListAccountPoints(reqMePoint.AUID, 0); pointList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_PointList_Scan_Error)
	} else {
		resp.Value = pointList
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 코인 리스트 조회
func GetMeCoinList(c echo.Context, reqMeCoin *context.ReqMeCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if coinList, err := model.GetDB().GetListAccountCoins(reqMeCoin.AUID); coinList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_CoinList_Scan_Error)
	} else {
		resp.Value = coinList
	}

	return c.JSON(http.StatusOK, resp)
}

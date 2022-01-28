package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

func GetMeta(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	swapList := context.Meta{
		PointList: model.GetDB().ScanPoints,
		AppPoints: model.GetDB().AppPoints,
		CoinList:  model.GetDB().Coins,
		Swapable:  model.GetDB().SwapAble,
	}

	resp.Value = swapList

	return c.JSON(http.StatusOK, resp)
}

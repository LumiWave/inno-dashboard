package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 전체 포인트 리스트 조회
func GetPointList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	resp.Value = model.GetDB().ScanPoints

	return c.JSON(http.StatusOK, resp)
}

// 전체 포인트 리스트 DB reload
func PostReLoadPointList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if err := model.GetDB().GetPointList(); err == nil {
		resp.Value = model.GetDB().ScanPoints
	} else {
		resp.SetReturn(resultcode.Result_DBError)
	}

	return c.JSON(http.StatusOK, resp)
}

// 전체 AppPoint 리스트 조회
func GetAppList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	resp.Value = model.GetDB().AppPoints.Apps

	return c.JSON(http.StatusOK, resp)
}

// 전체 App 리스트 DB reload
func PostReLoadAppList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if err := model.GetDB().GetApps(); err == nil {
		if err = model.GetDB().GetAppPoints(); err == nil {
			resp.Value = model.GetDB().AppPoints.Apps
		} else {
			resp.SetReturn(resultcode.Result_DBError)
		}
	} else {
		resp.SetReturn(resultcode.Result_DBError)
	}

	return c.JSON(http.StatusOK, resp)
}

// 전체 코인 리스트 조회
func GetCoinList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	resp.Value = model.GetDB().Coins.Coins

	return c.JSON(http.StatusOK, resp)
}

func PostReloadCoinList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if err := model.GetDB().GetAppCoins(); err == nil {
		if err = model.GetDB().GetCoins(); err == nil {
			resp.Value = model.GetDB().Coins.Coins
		} else {
			resp.SetReturn(resultcode.Result_DBError)
		}
	} else {
		resp.SetReturn(resultcode.Result_DBError)
	}

	return c.JSON(http.StatusOK, resp)
}

// App 포인트 별 당일 누적/전환량 정보 조회
func GetAppPoint(c echo.Context, reqAppPoint *context.ReqAppPoint) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if pointList, err := model.GetDB().GetListApplicationPoints(reqAppPoint.AppId); pointList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_App_Point_Scan_Error)
	} else {
		resp.Value = pointList
	}

	return c.JSON(http.StatusOK, resp)
}

// 코인 별 당일 누적/전환량 조회
func GetAppCoin(c echo.Context, reqAppCoin *context.ReqAppCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if coinList, err := model.GetDB().GetListApplicationCoins(reqAppCoin.AppId); coinList == nil || err != nil {
		resp.SetReturn(resultcode.Result_Get_App_Coin_Scan_Error)
	} else {
		resp.Value = coinList
	}

	return c.JSON(http.StatusOK, resp)
}

// App 포인트 별 유동량 history 조회
func GetAppPointHistory(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 코인별 유동량 history 조회
func GetAppCoinHistory(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 하루 최대 수집 포인트 양 조회
func GetAppPointMax(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

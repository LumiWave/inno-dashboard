package commonapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 공지 조회
func GetNotice(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 공지 등록
func PostNotice(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 공지 수정
func PutNotice(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// 공지 삭제
func DeleteNotice(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

func GetNews(c echo.Context, params *context.PageInfo) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if params.PageOffset == "0" {
		params.PageOffset = "1"
	}
	if resNews, err := model.GetDB().GetCacheNews(params.PageSize, params.PageOffset); err != nil {
		if newsList, totalCount, err := model.GetDB().USPAU_GetList_News(params); err != nil {
			resp.SetReturn(resultcode.ResultInternalServerError)
		} else {
			resNews := &context.ResNewsList{
				PageInfo:   *params,
				TotalCount: totalCount,
				List:       newsList,
			}
			resp.Value = resNews
			model.GetDB().SetCacheNews(params.PageSize, params.PageOffset, resNews)
		}
	} else {
		resp.Value = resNews
	}

	return c.JSON(http.StatusOK, resp)
}

func DeleteNewsCache(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if err := model.GetDB().DelCacheNews(); err != nil {
		resp.SetReturn(resultcode.ResultInternalServerError)
	}

	return c.JSON(http.StatusOK, resp)
}

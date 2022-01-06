package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
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

package internalapi

import (
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/labstack/echo"
)

// 공지 등록
func (o *InternalAPI) PostNotice(c echo.Context) error {
	return commonapi.PostNotice(c)
}

// 공지 수정
func (o *InternalAPI) PutNotice(c echo.Context) error {
	return commonapi.PutNotice(c)
}

// 공지 삭제
func (o *InternalAPI) DeleteNotice(c echo.Context) error {
	return commonapi.DeleteNotice(c)
}

// 뉴스 캐시 삭제
func (o *InternalAPI) DeleteNewsCache(c echo.Context) error {
	return commonapi.DeleteNewsCache(c)
}

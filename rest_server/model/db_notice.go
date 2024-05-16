package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_GetList_News = "[dbo].[USPAU_GetList_News]"
)

// 계정 코인 조회
func (o *DB) USPAU_GetList_News(pageInfo *context.PageInfo) ([]*context.News, int64, error) {
	var returnValue orginMssql.ReturnStatus
	Proc := USPAU_GetList_News
	var totalCount int64
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), Proc,
		sql.Named("PageID", pageInfo.PageOffset),
		sql.Named("PageSize", pageInfo.PageSize),
		sql.Named("TotalCount", sql.Out{Dest: &totalCount}),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Error(Proc+" QueryContext err : ", err)
		return nil, totalCount, err
	}

	newsList := []*context.News{}
	for rows.Next() {
		news := &context.News{}
		var isVisibled bool
		if err := rows.Scan(&news.NewsID, &news.Title, &news.StartSDT, &news.EndSDT, &news.BannerURL, &news.NewsURL, &isVisibled, &news.IsAlwaysVisibled); err != nil {
			log.Errorf(Proc+" Scan error %v", err)
			return nil, totalCount, err
		} else {
			newsList = append(newsList, news)
		}
	}

	if returnValue != 1 {
		log.Errorf(Proc+" returnvalue error : %v", returnValue)
		return nil, totalCount, errors.New(Proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return newsList, totalCount, nil
}

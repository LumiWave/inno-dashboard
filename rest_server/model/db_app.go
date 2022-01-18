package model

import (
	contextR "context"
	"database/sql"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_GetList_ApplicationPoints = "[dbo].[USPAU_GetList_ApplicationPoints]"
	USPAU_GetList_ApplicationCoins  = "[dbo].[USPAU_GetList_ApplicationCoins]"
)

// 앱 일일 포인트량 목록
func (o *DB) GetListApplicationPoints(AppId int64) ([]*context.AppPointDailyInfo, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationPoints,
		sql.Named("AppID", AppId),
		&returnValue)
	if err != nil {
		log.Errorf("%v", err)
		return nil, nil
	}

	var pointId int64
	var dailyQuantity, dailyExchangeQuantity int64
	var resetDate time.Time
	var appPointDailyInfoList []*context.AppPointDailyInfo

	for rows.Next() {
		if err := rows.Scan(&pointId, &dailyQuantity, &dailyExchangeQuantity, &resetDate); err != nil {
			log.Errorf("%v", err)
			return nil, err
		} else {
			appPointDailyInfo := &context.AppPointDailyInfo{
				PointID:               pointId,
				PointName:             o.ScanPointsMap[pointId].PointName,
				DailyQuantity:         dailyQuantity,
				DailyExchangeQuantity: dailyExchangeQuantity,
			}
			appPointDailyInfoList = append(appPointDailyInfoList, appPointDailyInfo)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return appPointDailyInfoList, nil

}

// 앱 일일 코인량 목록
func (o *DB) GetListApplicationCoins(AppId int64) ([]*context.AppCoinDailyInfo, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationCoins,
		sql.Named("AppID", AppId),
		&returnValue)
	if err != nil {
		log.Errorf("%v", err)
		return nil, nil
	}

	var coinId int64
	var dailyQuantity, dailyExchangeQuantity float64
	var resetDate time.Time
	var appCoinDailyInfoList []*context.AppCoinDailyInfo

	for rows.Next() {
		if err := rows.Scan(&coinId, &dailyQuantity, &dailyExchangeQuantity, &resetDate); err != nil {
			log.Errorf("%v", err)
			return nil, err
		} else {
			appCoinDailyInfo := &context.AppCoinDailyInfo{
				CoinID:                coinId,
				CoinSymbol:            o.CoinsMap[coinId].CoinSymbol,
				DailyQuantity:         dailyQuantity,
				DailyExchangeQuantity: dailyExchangeQuantity,
			}
			appCoinDailyInfoList = append(appCoinDailyInfoList, appCoinDailyInfo)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return appCoinDailyInfoList, nil
}

package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_GetList_ApplicationPoints = "[dbo].[USPAU_GetList_ApplicationPoints]"
	USPAU_GetList_ApplicationCoins  = "[dbo].[USPAU_GetList_ApplicationCoins]"
	USPW_GetList_DailyCoins         = "[dbo].[USPW_GetList_DailyCoins]"
	USPW_GetList_DailyPoints        = "[dbo].[USPW_GetList_DailyPoints]"
)

// 앱 일일 포인트량 목록
func (o *DB) GetListApplicationPoints(AppId int64) ([]*context.AppPointDailyInfo, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationPoints,
		sql.Named("AppID", AppId),
		&returnValue)
	if err != nil {
		log.Errorf("USPAU_GetList_ApplicationPoints QueryContext error: %v", err)
		return nil, nil
	}

	var pointId int64
	var dailyQuantity, dailyExchangeQuantity int64
	var resetDate time.Time
	var appPointDailyInfoList []*context.AppPointDailyInfo

	for rows.Next() {
		if err := rows.Scan(&pointId, &dailyQuantity, &dailyExchangeQuantity, &resetDate); err != nil {
			log.Errorf("USPAU_GetList_ApplicationPoints Scan error : %v", err)
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
		log.Errorf("USPAU_GetList_ApplicationPoints returnvalue error : %v", returnValue)
		return nil, errors.New("USPAU_GetList_ApplicationPoints returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return appPointDailyInfoList, nil

}

// 앱 일일 코인량 목록
func (o *DB) GetListApplicationCoins(AppId int64) ([]*context.AppCoinDailyInfo, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationCoins,
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
		log.Errorf("USPAU_GetList_ApplicationCoins returnvalue error : %v", returnValue)
		return nil, errors.New("USPAU_GetList_ApplicationCoins returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return appCoinDailyInfoList, nil
}

// 일일 코인 유동량 검색
func (o *DB) GetListDailyCoins(reqCoinLiquidity *context.ReqCoinLiquidity) (*context.CoinLiquidity, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), USPW_GetList_DailyCoins,
		sql.Named("BaseDate", reqCoinLiquidity.BaseDate),
		sql.Named("CoinID", reqCoinLiquidity.CoinID),
		sql.Named("Interval", reqCoinLiquidity.Interval),
		&returnValue)
	if err != nil {
		log.Errorf("%v", err)
		return nil, nil
	}

	coinLiquidity := new(context.CoinLiquidity)
	for rows.Next() {
		if err := rows.Scan(&coinLiquidity.BaseDate, &coinLiquidity.AcqQuantity, &coinLiquidity.AcqCount,
			&coinLiquidity.CnsmQuantity, coinLiquidity.CnsmCount, coinLiquidity.AcqExchangeQuantity,
			coinLiquidity.PointsToCoinsCount, coinLiquidity.CnsmExchangeQuantity, coinLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("USPW_GetList_DailyCoins returnvalue error : %v", returnValue)
		return nil, errors.New("USPW_GetList_DailyCoins returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return coinLiquidity, nil
}

// 일일 포인트 유동량 검색
func (o *DB) GetListPointCoins(reqPointLiquidity *context.ReqPointLiquidity) (*context.PointLiquidity, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), USPW_GetList_DailyPoints,
		sql.Named("BaseDate", reqPointLiquidity.BaseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		&returnValue)
	if err != nil {
		log.Errorf("%v", err)
		return nil, nil
	}

	pointLiquidity := new(context.PointLiquidity)
	for rows.Next() {
		if err := rows.Scan(&pointLiquidity.BaseDate, &pointLiquidity.AcqQuantity, &pointLiquidity.AcqCount,
			&pointLiquidity.CnsmQuantity, pointLiquidity.CnsmCount, pointLiquidity.AcqExchangeQuantity,
			pointLiquidity.PointsToCoinsCount, pointLiquidity.CnsmExchangeQuantity, pointLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("USPW_GetList_DailyPoints returnvalue error : %v", returnValue)
		return nil, errors.New("USPW_GetList_DailyPoints returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return pointLiquidity, nil
}

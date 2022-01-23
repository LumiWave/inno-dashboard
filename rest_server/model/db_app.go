package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPW_GetList_HourlyCoins  = "[dbo].[USPW_GetList_HourlyCoins]"
	USPW_GetList_DailyCoins   = "[dbo].[USPW_GetList_DailyCoins]"
	USPW_GetList_WeeklyCoins  = "[dbo].[USPW_GetList_WeeklyCoins]"
	USPW_GetList_MonthlyCoins = "[dbo].[USPW_GetList_MonthlyCoins]"

	USPW_GetList_HourlyPoints  = "[dbo].[USPW_GetList_HourlyPoints]"
	USPW_GetList_DailyPoints   = "[dbo].[USPW_GetList_DailyPoints]"
	USPW_GetList_WeeklyPoints  = "[dbo].[USPW_GetList_WeeklyPoints]"
	USPW_GetList_MonthlyPoints = "[dbo].[USPW_GetList_MonthlyPoints]"
)

// 일일 코인 유동량 검색
func (o *DB) GetListDailyCoins(reqCoinLiquidity *context.ReqCoinLiquidity) ([]*context.CoinLiquidity, error) {
	baseDate := &reqCoinLiquidity.BaseDate
	if reqCoinLiquidity.BaseDate.IsZero() {
		baseDate = nil
	}

	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), USPW_GetList_DailyCoins,
		sql.Named("BaseDate", baseDate),
		sql.Named("CoinID", reqCoinLiquidity.CoinID),
		sql.Named("Interval", reqCoinLiquidity.Interval),
		&returnValue)
	if err != nil {
		log.Errorf("USPW_GetList_DailyCoins QueryContext error : %v", err)
		return nil, nil
	}

	coinLiquiditys := []*context.CoinLiquidity{}
	for rows.Next() {
		coinLiquidity := new(context.CoinLiquidity)
		if err := rows.Scan(&coinLiquidity.BaseDate, &coinLiquidity.AcqQuantity, &coinLiquidity.AcqCount,
			&coinLiquidity.CnsmQuantity, &coinLiquidity.CnsmCount, &coinLiquidity.AcqExchangeQuantity,
			&coinLiquidity.PointsToCoinsCount, &coinLiquidity.CnsmExchangeQuantity, &coinLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("USPW_GetList_DailyCoins Scan error : %v", err)
			return nil, err
		} else {
			coinLiquiditys = append(coinLiquiditys, coinLiquidity)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("USPW_GetList_DailyCoins returnvalue error : %v", returnValue)
		return nil, errors.New("USPW_GetList_DailyCoins returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return coinLiquiditys, nil
}

// 일일 포인트 유동량 검색
func (o *DB) GetListDailyPoints(reqPointLiquidity *context.ReqPointLiquidity) ([]*context.PointLiquidity, error) {
	baseDate := &reqPointLiquidity.BaseDate
	if reqPointLiquidity.BaseDate.IsZero() {
		baseDate = nil
	}

	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), USPW_GetList_DailyPoints,
		sql.Named("BaseDate", baseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		&returnValue)
	if err != nil {
		log.Errorf("USPW_GetList_DailyPoints QueryContext error : %v", err)
		return nil, nil
	}

	pointLiquiditys := []*context.PointLiquidity{}
	for rows.Next() {
		pointLiquidity := new(context.PointLiquidity)
		if err := rows.Scan(&pointLiquidity.BaseDate, &pointLiquidity.AcqQuantity, &pointLiquidity.AcqCount,
			&pointLiquidity.CnsmQuantity, &pointLiquidity.CnsmCount, &pointLiquidity.AcqExchangeQuantity,
			&pointLiquidity.PointsToCoinsCount, &pointLiquidity.CnsmExchangeQuantity, &pointLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("USPW_GetList_DailyPoints Scan error : %v", err)
			return nil, err
		} else {
			pointLiquiditys = append(pointLiquiditys, pointLiquidity)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("USPW_GetList_DailyPoints returnvalue error : %v", returnValue)
		return nil, errors.New("USPW_GetList_DailyPoints returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return pointLiquiditys, nil
}

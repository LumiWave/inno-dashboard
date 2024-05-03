package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPW_GetList_HourlyPoints  = "[dbo].[USPW_GetList_HourlyPoints]"
	USPW_GetList_DailyPoints   = "[dbo].[USPW_GetList_DailyPoints]"
	USPW_GetList_WeeklyPoints  = "[dbo].[USPW_GetList_WeeklyPoints]"
	USPW_GetList_MonthlyPoints = "[dbo].[USPW_GetList_MonthlyPoints]"
)

// 포인트 유동량 검색
func (o *DB) GetListPointLiquidity(procedureType string, reqPointLiquidity *context.ReqPointLiquidity) ([]*context.PointLiquidity, error) {
	baseDate := ChangeBaseTime(reqPointLiquidity.BaseDate)

	baseDateStr := "BaseDate"
	if strings.EqualFold(procedureType, "USPW_GetList_HourlyPoints") {
		baseDateStr = "BaseSDT"
	}
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.QueryContext(contextR.Background(), procedureType,
		sql.Named(baseDateStr, baseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%v QueryContext error : %v", procedureType, err)
		return nil, err
	}

	pointLiquiditys := []*context.PointLiquidity{}
	for rows.Next() {
		pointLiquidity := new(context.PointLiquidity)
		if err := rows.Scan(&pointLiquidity.BaseDate, &pointLiquidity.AcqQuantity, &pointLiquidity.AcqCount,
			&pointLiquidity.CnsmQuantity, &pointLiquidity.CnsmCount, &pointLiquidity.AcqExchangeQuantity,
			&pointLiquidity.CoinsToPointsCount, &pointLiquidity.CnsmExchangeQuantity, &pointLiquidity.PointsToCoinsCount); err != nil {
			log.Errorf("%v Scan error : %v", procedureType, err)
			return nil, err
		} else {
			pointLiquidity.BaseDateToNumber = pointLiquidity.BaseDate.Unix()
			pointLiquiditys = append(pointLiquiditys, pointLiquidity)
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", procedureType, returnValue)
		return nil, errors.New(procedureType + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return pointLiquiditys, nil
}

// 일별 포인트 유동량 검색
func (o *DB) GetListDailyPoints(reqPointLiquidity *context.ReqPointLiquidity) ([]*context.PointLiquidity, error) {
	baseDate := ChangeTime(reqPointLiquidity.BaseDate)

	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.QueryContext(contextR.Background(), USPW_GetList_DailyPoints,
		sql.Named("BaseDate", baseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

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

	if returnValue != 1 {
		log.Errorf("USPW_GetList_DailyPoints returnvalue error : %v", returnValue)
		return nil, errors.New("USPW_GetList_DailyPoints returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return pointLiquiditys, nil
}

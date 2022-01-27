package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPW_GetList_HourlyPoints  = "[dbo].[USPW_GetList_HourlyPoints]"
	USPW_GetList_DailyPoints   = "[dbo].[USPW_GetList_DailyPoints]"
	USPW_GetList_WeeklyPoints  = "[dbo].[USPW_GetList_WeeklyPoints]"
	USPW_GetList_MonthlyPoints = "[dbo].[USPW_GetList_MonthlyPoints]"
)

// 포인트 유동량 검색
func (o *DB) GetListPointLiquidity(dateType string, reqPointLiquidity *context.ReqPointLiquidity) ([]*context.PointLiquidity, *time.Time, error) {
	baseDate := ChangeTime(reqPointLiquidity.BaseDate)

	baseDateStr := "BaseDate"
	firstDateStr := "FirstDate"
	if strings.EqualFold(dateType, "USPW_GetList_HourlyPoints") {
		baseDateStr = "BaseSDT"
		firstDateStr = "FirstSDT"
	}
	firstDate := &time.Time{}
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), dateType,
		sql.Named(baseDateStr, baseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		sql.Named(firstDateStr, sql.Out{Dest: &firstDate}),
		&returnValue)
	if err != nil {
		log.Errorf("%v QueryContext error : %v", dateType, err)
		return nil, nil, err
	}

	pointLiquiditys := []*context.PointLiquidity{}
	for rows.Next() {
		pointLiquidity := new(context.PointLiquidity)
		if err := rows.Scan(&pointLiquidity.BaseDate, &pointLiquidity.AcqQuantity, &pointLiquidity.AcqCount,
			&pointLiquidity.CnsmQuantity, &pointLiquidity.CnsmCount, &pointLiquidity.AcqExchangeQuantity,
			&pointLiquidity.PointsToCoinsCount, &pointLiquidity.CnsmExchangeQuantity, &pointLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("%v Scan error : %v", dateType, err)
			return nil, nil, err
		} else {
			pointLiquidity.BaseDateToNumber = pointLiquidity.BaseDate.Unix()
			pointLiquiditys = append(pointLiquiditys, pointLiquidity)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", dateType, returnValue)
		return nil, nil, errors.New(dateType + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return pointLiquiditys, firstDate, nil
}

// 일별 포인트 유동량 검색
func (o *DB) GetListDailyPoints(reqPointLiquidity *context.ReqPointLiquidity) ([]*context.PointLiquidity, error) {
	baseDate := ChangeTime(reqPointLiquidity.BaseDate)

	//firstDate := &time.Time{}
	firstDate := ""
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), USPW_GetList_DailyPoints,
		sql.Named("BaseDate", baseDate),
		sql.Named("AppID", reqPointLiquidity.AppID),
		sql.Named("PointID", reqPointLiquidity.PointID),
		sql.Named("Interval", reqPointLiquidity.Interval),
		sql.Named("FirstDate", sql.Out{Dest: &firstDate}),
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

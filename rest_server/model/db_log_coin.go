package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPW_GetList_HourlyCoins  = "[dbo].[USPW_GetList_HourlyCoins]"
	USPW_GetList_DailyCoins   = "[dbo].[USPW_GetList_DailyCoins]"
	USPW_GetList_WeeklyCoins  = "[dbo].[USPW_GetList_WeeklyCoins]"
	USPW_GetList_MonthlyCoins = "[dbo].[USPW_GetList_MonthlyCoins]"
)

// 코인 유동량 검색
func (o *DB) GetListCoinLiquidity(procedureType string, reqCoinLiquidity *context.ReqCoinLiquidity) ([]*context.CoinLiquidity, error) {
	baseDate := ChangeBaseTime(reqCoinLiquidity.BaseDate)

	baseDateStr := "BaseDate"
	if strings.EqualFold(procedureType, "USPW_GetList_HourlyCoins") {
		baseDateStr = "BaseSDT"
	}
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlLogRead.GetDB().QueryContext(contextR.Background(), procedureType,
		sql.Named(baseDateStr, baseDate),
		sql.Named("CoinID", reqCoinLiquidity.CoinID),
		sql.Named("Interval", reqCoinLiquidity.Interval),
		&returnValue)
	if err != nil {
		log.Errorf("%v QueryContext error : %v", procedureType, err)
		return nil, err
	}

	coinLiquiditys := []*context.CoinLiquidity{}
	for rows.Next() {
		coinLiquidity := new(context.CoinLiquidity)
		if err := rows.Scan(&coinLiquidity.BaseDate, &coinLiquidity.AcqQuantity, &coinLiquidity.AcqCount,
			&coinLiquidity.CnsmQuantity, &coinLiquidity.CnsmCount, &coinLiquidity.AcqExchangeQuantity,
			&coinLiquidity.PointsToCoinsCount, &coinLiquidity.CnsmExchangeQuantity, &coinLiquidity.CoinsToPointsCount); err != nil {
			log.Errorf("%v Scan error : %v", procedureType, err)
			return nil, err
		} else {
			coinLiquidity.BaseDateToNumber = coinLiquidity.BaseDate.Unix()
			coinLiquiditys = append(coinLiquiditys, coinLiquidity)
		}
	}

	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", procedureType, returnValue)
		return nil, errors.New(procedureType + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return coinLiquiditys, nil
}

// 일별 코인 유동량 검색
func (o *DB) GetListDailyCoins(reqCoinLiquidity *context.ReqCoinLiquidity) ([]*context.CoinLiquidity, error) {
	baseDate := ChangeTime(reqCoinLiquidity.BaseDate)

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

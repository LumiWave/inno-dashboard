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
	USPAU_GetList_AccountCoins      = "[dbo].[USPAU_GetList_AccountCoins]"
	USPAU_GetList_ApplicationCoins  = "[dbo].[USPAU_GetList_ApplicationCoins]"
	USPAU_GetList_ApplicationPoints = "[dbo].[USPAU_GetList_ApplicationPoints]"
)

// 계정 코인 조회
func (o *DB) GetListAccountCoins(auid int64) ([]*context.MeWalletInfo, error) {
	var returnValue orginMssql.ReturnStatus
	rows, _ := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_AccountCoins,
		sql.Named("AUID", auid),
		&returnValue)

	var coinId int64
	var walletAddress string
	var quantity, dailyQuantity string
	var resetDate time.Time

	var meWalletList []*context.MeWalletInfo

	for rows.Next() {
		if err := rows.Scan(&coinId, &walletAddress, &quantity, &dailyQuantity, &resetDate); err != nil {
			log.Error(err)
			return nil, err
		} else {
			meWallet := &context.MeWalletInfo{
				CoinID:        coinId,
				WalletAddress: walletAddress,
				Quantity:      quantity,
				DailyQuantity: dailyQuantity,
				ResetDate:     resetDate,
			}
			meWalletList = append(meWalletList, meWallet)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return meWalletList, nil
}

// 앱 일일 코인량 목록
func (o *DB) GetListApplicationCoins(AppId int64) ([]*context.MeCoin, error) {
	var returnValue orginMssql.ReturnStatus
	rows, _ := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationCoins,
		sql.Named("AppID", AppId),
		&returnValue)

	var coinId int64
	var dailyQuantity, dailyExchangeQuantity string
	var resetDate time.Time
	var meCoinList []*context.MeCoin

	for rows.Next() {
		if err := rows.Scan(&coinId, &dailyQuantity, &dailyExchangeQuantity, &resetDate); err != nil {
			log.Error(err)
			return nil, err
		} else {
			meCoin := &context.MeCoin{
				CoinID:                coinId,
				CoinSymbol:            o.CoinsMap[coinId].CoinSymbol,
				DailyQuantity:         dailyQuantity,
				DailyExchangeQuantity: dailyExchangeQuantity,
			}
			meCoinList = append(meCoinList, meCoin)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return meCoinList, nil
}

// 앱 일일 포인트량 목록
func (o *DB) GetListApplicationPoints(AppId int64) ([]*context.MePoint, error) {
	var returnValue orginMssql.ReturnStatus
	rows, _ := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_ApplicationPoints,
		sql.Named("AppID", AppId),
		&returnValue)

	var pointId int64
	var dailyQuantity, dailyExchangeQuantity int64
	var resetDate time.Time
	var mePointList []*context.MePoint

	for rows.Next() {
		if err := rows.Scan(&pointId, &dailyQuantity, &dailyExchangeQuantity, &resetDate); err != nil {
			log.Error(err)
			return nil, err
		} else {
			mePoint := &context.MePoint{
				PointID:               pointId,
				PointName:             o.ScanPointsMap[pointId].PointName,
				DailyQuantity:         dailyQuantity,
				DailyExchangeQuantity: dailyExchangeQuantity,
			}
			mePointList = append(mePointList, mePoint)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return mePointList, nil

}

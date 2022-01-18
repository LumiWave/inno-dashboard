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
	USPAU_GetList_AccountCoins  = "[dbo].[USPAU_GetList_AccountCoins]"
	USPAU_GetList_AccountPoints = "[dbo].[USPAU_GetList_AccountPoints]"
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
			log.Errorf("%v", err)
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

// 계정 포인트 조회
func (o *DB) GetListAccountPoints(auid, muid int64) ([]*context.MePoint, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_AccountPoints,
		sql.Named("AUID", auid),
		sql.Named("MUID", muid),
		&returnValue)

	if err != nil {
		log.Error("QueryContext err : ", err)
		return nil, err
	}

	var appID, pointID int64
	var dailyQuantity float64
	var resetDate time.Time

	var mePointList []*context.MePoint

	for rows.Next() {
		if err := rows.Scan(&appID, &pointID, &dailyQuantity, &resetDate); err != nil {
			log.Errorf("%v", err)
			return nil, err
		} else {
			mePoint := &context.MePoint{
				AppID:         appID,
				PointID:       pointID,
				DailyQuantity: dailyQuantity,
				ResetDate:     resetDate,
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

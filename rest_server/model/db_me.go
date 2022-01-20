package model

import (
	contextR "context"
	"database/sql"

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
	rows, err := o.MssqlAccount.GetDB().QueryContext(contextR.Background(), USPAU_GetList_AccountCoins,
		sql.Named("AUID", auid),
		&returnValue)

	if err != nil {
		log.Error("USPAU_GetList_AccountCoins QueryContext err : ", err)
		return nil, err
	}

	var meWalletList []*context.MeWalletInfo
	for rows.Next() {
		meWallet := &context.MeWalletInfo{}
		if err := rows.Scan(&meWallet.CoinID, &meWallet.WalletAddress, &meWallet.Quantity, &meWallet.TodayAcqQuantity, &meWallet.TodayCnsmQuantity, &meWallet.ResetDate); err != nil {
			log.Errorf("USPAU_GetList_AccountCoins Scan error %v", err)
			return nil, err
		} else {
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
		log.Error("USPAU_GetList_AccountPoints QueryContext err : %v", err)
		return nil, err
	}

	var mePointList []*context.MePoint

	for rows.Next() {
		mePoint := context.MePoint{}
		if err := rows.Scan(&mePoint.AppID, &mePoint.PointID, &mePoint.TodayLimitedQuantity, &mePoint.TodayAcqQuantity, &mePoint.TodayCnsmQuantity, &mePoint.ResetDate); err != nil {
			log.Errorf("USPAU_GetList_AccountPoints Scan error : %v", err)
			return nil, err
		} else {
			mePointList = append(mePointList, &mePoint)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		return nil, nil
	}
	return mePointList, nil
}

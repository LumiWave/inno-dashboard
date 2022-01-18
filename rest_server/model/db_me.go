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
	USPAU_GetList_AccountCoins = "[dbo].[USPAU_GetList_AccountCoins]"
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

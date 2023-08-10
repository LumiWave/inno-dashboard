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
	USPAU_GetList_AccountWallets = "[dbo].[USPAU_GetList_AccountWallets]"
)

func (o *DB) USPAU_GetList_AccountWallets(auid int64) (map[int64]*context.DBWalletRegist, error) {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_GetList_AccountWallets
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		sql.Named("AUID", auid),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return nil, err
	}

	DBWalletRegistMap := make(map[int64]*context.DBWalletRegist)

	for rows.Next() {
		data := &context.DBWalletRegist{}
		if err := rows.Scan(&data.BaseCoinID, &data.WalletID, &data.WalletAddress, &data.DisconnectedWalletAddress, &data.DisconnectedDT, &data.ModifiedDT); err != nil {
			log.Errorf("%s Scan error : %v", proc, err)
			return nil, err
		} else {
			if _, ok := DBWalletRegistMap[data.BaseCoinID]; !ok {
				DBWalletRegistMap[data.BaseCoinID] = data
			}
		}
	}

	if returnValue != 1 {
		log.Errorf("%s returnvalue error : %v", proc, returnValue)
		return nil, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return DBWalletRegistMap, nil
}

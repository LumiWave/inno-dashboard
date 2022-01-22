package model

import (
	contextR "context"
	"errors"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_Scan_ExchangeGoods = "[dbo].[USPAU_Scan_ExchangeGoods]"
)

// 교환 가능 코인, 포인트 정보
func (o *DB) GetScanExchangeGoods() ([]*context.Swapable, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.GetDB().QueryContext(contextR.Background(), USPAU_Scan_ExchangeGoods,
		&returnValue)
	if err != nil {
		log.Errorf("%v", err)
		return nil, nil
	}

	swapableList := []*context.Swapable{}

	for rows.Next() {
		swapAble := &context.Swapable{}
		if err := rows.Scan(&swapAble.AppID, &swapAble.CoinID, &swapAble.PointID); err != nil {
			log.Errorf("%v", err)
			return nil, err
		} else {
			swapableList = append(swapableList, swapAble)
		}
	}
	defer rows.Close()

	if returnValue != 1 {
		log.Errorf("USPAU_Scan_ExchangeGoods returnvalue error : %v", returnValue)
		return nil, errors.New("USPAU_Scan_ExchangeGoods returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return swapableList, nil

}

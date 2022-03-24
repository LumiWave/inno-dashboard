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
func (o *DB) GetScanExchangeGoods() error {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.GetDB().QueryContext(contextR.Background(), USPAU_Scan_ExchangeGoods,
		&returnValue)
	if err != nil {
		log.Errorf("USPAU_Scan_ExchangeGoods QueryContext error : %v", err)
		return nil
	}
	defer rows.Close()

	o.SwapAbleMap = make(map[int64]*context.Swapable)
	o.SwapAble = nil

	for rows.Next() {
		swapAble := &context.Swapable{}
		swapablePoint := &context.SwapablePoint{}
		if err := rows.Scan(&swapAble.AppID, &swapAble.CoinID, &swapAble.BaseCoinID, &swapablePoint.PointID, &swapAble.IsEnable); err != nil {
			log.Errorf("USPAU_Scan_ExchangeGoods Scan error : %v", err)
			return err
		} else {
			if swapItem, ok := o.SwapAbleMap[swapAble.AppID]; ok {
				swapItem.Points = append(swapItem.Points, swapablePoint)
			} else {
				swapAble.Points = append(swapAble.Points, swapablePoint)
				o.SwapAbleMap[swapAble.AppID] = swapAble
			}
		}
	}

	for _, swabAble := range o.SwapAbleMap {
		o.SwapAble = append(o.SwapAble, swabAble)
	}

	if returnValue != 1 {
		log.Errorf("USPAU_Scan_ExchangeGoods returnvalue error : %v", returnValue)
		return errors.New("USPAU_Scan_ExchangeGoods returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil

}

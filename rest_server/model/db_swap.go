package model

import (
	contextR "context"
	"errors"
	"strconv"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	//USPAU_Scan_ExchangeGoods        = "[dbo].[USPAU_Scan_ExchangeGoods]"
	USPAU_Scan_ExchangeCoinToCoins  = "[dbo].[USPAU_Scan_ExchangeCoinToCoins]"
	USPAU_Scan_ExchangePointToCoins = "[dbo].[USPAU_Scan_ExchangePointToCoins]"
	USPAU_Scan_ExchangeCoinToPoints = "[dbo].[USPAU_Scan_ExchangeCoinToPoints]"
)

// 교환 가능 코인, 포인트 정보
// func (o *DB) GetScanExchangeGoods() error {
// 	var returnValue orginMssql.ReturnStatus
// 	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), USPAU_Scan_ExchangeGoods,
// 		&returnValue)

// 	if rows != nil {
// 		defer rows.Close()
// 	}

// 	if err != nil {
// 		log.Errorf("USPAU_Scan_ExchangeGoods QueryContext error : %v", err)
// 		return nil
// 	}

// 	o.SwapAbleMap = make(map[int64]*context.Swapable)
// 	o.SwapAble = nil

// 	for rows.Next() {
// 		swapAble := &context.Swapable{}
// 		swapablePoint := &context.SwapablePoint{}
// 		if err := rows.Scan(&swapAble.AppID, &swapAble.CoinID, &swapAble.BaseCoinID, &swapablePoint.PointID, &swapAble.IsEnable); err != nil {
// 			log.Errorf("USPAU_Scan_ExchangeGoods Scan error : %v", err)
// 			return err
// 		} else {
// 			if swapItem, ok := o.SwapAbleMap[swapAble.AppID]; ok {
// 				swapItem.Points = append(swapItem.Points, swapablePoint)
// 			} else {
// 				swapAble.Points = append(swapAble.Points, swapablePoint)
// 				o.SwapAbleMap[swapAble.AppID] = swapAble
// 			}
// 		}
// 	}

// 	for _, swabAble := range o.SwapAbleMap {
// 		o.SwapAble = append(o.SwapAble, swabAble)
// 	}

// 	if returnValue != 1 {
// 		log.Errorf("USPAU_Scan_ExchangeGoods returnvalue error : %v", returnValue)
// 		return errors.New("USPAU_Scan_ExchangeGoods returnvalue error " + strconv.Itoa(int(returnValue)))
// 	}
// 	return nil

// }

func (o *DB) USPAU_Scan_ExchangeCoinToCoins() error {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Scan_ExchangeCoinToCoins
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%v QueryContext error : %v", proc, err)
		return nil
	}

	o.SwapAbleCoinToCoins = nil

	for rows.Next() {
		swapAble := &context.SwapCointoCon{}

		if err := rows.Scan(&swapAble.FromBaseCoinID,
			&swapAble.FromID,
			&swapAble.ToBaseCoinID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.SortOrder,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangeCoinToCoins Scan error : %v", err)
			return err
		} else {
			o.SwapAbleCoinToCoins = append(o.SwapAbleCoinToCoins, swapAble)
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil
}

func (o *DB) USPAU_Scan_ExchangePointToCoins() error {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Scan_ExchangePointToCoins
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%v QueryContext error : %v", proc, err)
		return nil
	}

	o.SwapAblePointToCoins = nil

	for rows.Next() {
		swapAble := &context.SwapPointToCoin{}

		if err := rows.Scan(&swapAble.FromID,
			&swapAble.ToBaseCoinID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.SortOrder,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangePointToCoins Scan error : %v", err)
			return err
		} else {
			o.SwapAblePointToCoins = append(o.SwapAblePointToCoins, swapAble)
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil
}

func (o *DB) USPAU_Scan_ExchangeCoinToPoints() error {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Scan_ExchangeCoinToPoints
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%v QueryContext error : %v", proc, err)
		return nil
	}

	o.SwapAbleCoinToPoints = nil

	for rows.Next() {
		swapAble := &context.SwapCoinToPoint{}

		if err := rows.Scan(&swapAble.FromBaseCoinID,
			&swapAble.FromID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.SortOrder,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangeCoinToPoints Scan error : %v", err)
			return err
		} else {
			o.SwapAbleCoinToPoints = append(o.SwapAbleCoinToPoints, swapAble)
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil
}

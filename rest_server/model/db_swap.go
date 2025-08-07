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
	USPAU_Scan_ExchangeCoinToCoins   = "[dbo].[USPAU_Scan_ExchangeCoinToCoins]"
	USPAU_Scan_ExchangePointToCoins  = "[dbo].[USPAU_Scan_ExchangePointToCoins]"
	USPAU_Scan_ExchangeCoinToPoints  = "[dbo].[USPAU_Scan_ExchangeCoinToPoints]"
	USPAU_Scan_ExchangePointToPoints = "[dbo].[USPAU_Scan_ExchangePointToPoints]"
)

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

	o.SwapAbleCoinToCoins = []*context.SwapCointoCoin{}
	o.SwapAbleC2CsMap = make(map[int64]map[int64]*context.SwapCointoCoin)

	for rows.Next() {
		swapAble := &context.SwapCointoCoin{}

		if err := rows.Scan(&swapAble.FromBaseCoinID,
			&swapAble.FromID,
			&swapAble.ToBaseCoinID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangeCoinToCoins Scan error : %v", err)
			return err
		} else {
			o.SwapAbleCoinToCoins = append(o.SwapAbleCoinToCoins, swapAble)

			if o.SwapAbleC2CsMap[swapAble.FromID] == nil {
				o.SwapAbleC2CsMap[swapAble.FromID] = make(map[int64]*context.SwapCointoCoin)
			}
			o.SwapAbleC2CsMap[swapAble.FromID][swapAble.ToID] = swapAble
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

	o.SwapAblePointToCoins = []*context.SwapPointToCoin{}
	o.SwapAbleP2CsMap = make(map[int64]map[int64]*context.SwapPointToCoin)

	for rows.Next() {
		swapAble := &context.SwapPointToCoin{}

		if err := rows.Scan(&swapAble.FromID,
			&swapAble.ToBaseCoinID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangePointToCoins Scan error : %v", err)
			return err
		} else {
			o.SwapAblePointToCoins = append(o.SwapAblePointToCoins, swapAble)

			if o.SwapAbleP2CsMap[swapAble.FromID] == nil {
				o.SwapAbleP2CsMap[swapAble.FromID] = make(map[int64]*context.SwapPointToCoin)
			}
			o.SwapAbleP2CsMap[swapAble.FromID][swapAble.ToID] = swapAble
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

	o.SwapAbleCoinToPoints = []*context.SwapCoinToPoint{}
	o.SwapAbleC2PsMap = make(map[int64]map[int64]*context.SwapCoinToPoint)

	for rows.Next() {
		swapAble := &context.SwapCoinToPoint{}

		if err := rows.Scan(&swapAble.FromBaseCoinID,
			&swapAble.FromID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("USPAU_Scan_ExchangeCoinToPoints Scan error : %v", err)
			return err
		} else {
			o.SwapAbleCoinToPoints = append(o.SwapAbleCoinToPoints, swapAble)

			if o.SwapAbleC2PsMap[swapAble.FromID] == nil {
				o.SwapAbleC2PsMap[swapAble.FromID] = make(map[int64]*context.SwapCoinToPoint)
			}
			o.SwapAbleC2PsMap[swapAble.FromID][swapAble.ToID] = swapAble
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil
}

func (o *DB) USPAU_Scan_ExchangePointToPoints() error {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Scan_ExchangePointToPoints
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%v QueryContext error : %v", proc, err)
		return nil
	}

	o.SwapAblePointToPoints = []*context.SwapPointToPoint{}
	o.SwapAbleP2PsMap = make(map[int64]map[int64]*context.SwapPointToPoint)

	for rows.Next() {
		swapAble := &context.SwapPointToPoint{}

		if err := rows.Scan(
			&swapAble.FromID,
			&swapAble.ToID,
			&swapAble.IsEnabled,
			&swapAble.IsVisible,
			&swapAble.MinimumExchangeQuantity,
			&swapAble.ExchangeRatio); err != nil {
			log.Errorf("%v Scan error : %v", proc, err)
			return err
		} else {
			o.SwapAblePointToPoints = append(o.SwapAblePointToPoints, swapAble)

			if o.SwapAbleP2PsMap[swapAble.FromID] == nil {
				o.SwapAbleP2PsMap[swapAble.FromID] = make(map[int64]*context.SwapPointToPoint)
			}
			o.SwapAbleP2PsMap[swapAble.FromID][swapAble.ToID] = swapAble
		}
	}

	if returnValue != 1 {
		log.Errorf("%v returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return nil
}

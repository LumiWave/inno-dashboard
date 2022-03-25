package model

import (
	"math"
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
)

func MakeCoinFeeKey(baseSymbol string) string {
	return config.GetInstance().DBPrefix + ":COIN-:" + baseSymbol
}

func (o *DB) SetCacheCoinFee(key string, data *context.ResGetCoinFee) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}

	return o.Cache.Set(key, data, -1)
}

func (o *DB) GetCacheCoinFee(key string) (*context.ResGetCoinFee, error) {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}

	res := new(context.ResGetCoinFee)
	err := o.Cache.Get(key, res)

	return res, err
}

func (o *DB) UpdateCoinFee() {
	go func() {
		for {
			for _, baseCoin := range o.BaseCoinMapByCoinID {
				req := &point_manager_server.ReqCoinFee{
					Symbol: baseCoin.BaseCoinSymbol,
				}
				if fee, err := point_manager_server.GetInstance().GetCoinFee(req); err != nil {
					log.Errorf("GetCoinFee err : %v", err)
				} else {
					gasPrice, _ := strconv.ParseFloat(fee.ResCoinFeeInfoValue.Fast, 64)
					gasPrice = toFixed(gasPrice*0.000000001, 18)
					transactionFee := gasPrice * 52346
					newFee := &context.ResGetCoinFee{
						BaseCoinID:     baseCoin.BaseCoinID,
						BaseCoinSymbol: baseCoin.BaseCoinSymbol,
						GasPrice:       gasPrice,
						TransactionFee: transactionFee,
					}

					key := MakeCoinFeeKey(baseCoin.BaseCoinSymbol)
					o.SetCacheCoinFee(key, newFee)
				}
			}

			timer := time.NewTimer(15 * time.Second)
			//timer := time.NewTimer(5 * time.Second)
			<-timer.C
			timer.Stop()
		}
	}()
}

func round(num float64) int {
	return int(num + math.Copysign(0, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

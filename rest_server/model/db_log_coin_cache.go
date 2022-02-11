package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

func MakeLogKeyOfCoin(CoinID int64, candleType string) string {
	return config.GetInstance().DBPrefix + ":LOG-APP-COIN:" + strconv.FormatInt(CoinID, 10) + "-" + candleType
}

func (o *DB) HSetLogOfCoin(key string, field string, value interface{}) error {
	return o.Cache.HSet(key, field, value)
}

func (o *DB) HGetLogOfCoin(key string, field string, value interface{}) error {
	return o.Cache.HGet(key, field, value)
}

func (o *DB) ZADDLogOfCoin(key string, score int64, value interface{}) error {
	return o.Cache.ZAdd(key, basedb.Z{Score: float64(score), Member: value})
}

func (o *DB) ZADDLogOfCoinSlice(key string, liqs []*context.CoinLiquidity) error {
	z := []basedb.Z{}
	for _, liq := range liqs {
		z = append(z, basedb.Z{
			Score:  float64(liq.BaseDateToNumber),
			Member: liq,
		})
	}
	return o.Cache.ZAdd(key, z...)
}

func (o *DB) ZRemRangeByScoreOfCoin(key, min, max string) (int64, error) {
	return o.Cache.ZRemRangeByScore(key, min, max)
}

func (o *DB) ZRangeLogOfCoin(key string, start, stop int64) ([]*context.CoinLiquidity, error) {
	coinLiqs := []*context.CoinLiquidity{}

	list, err := o.Cache.ZRange(key, start, stop)
	for _, member := range list {
		coinLiq := &context.CoinLiquidity{}
		json.Unmarshal([]byte(fmt.Sprintf("%v", member.Member)), coinLiq)
		coinLiqs = append(coinLiqs, coinLiq)
	}
	return coinLiqs, err
}

func (o *DB) ZRevRangeLogOfCoin(key string, start, stop int64) ([]*context.CoinLiquidity, error) {

	coinLiqs := []*context.CoinLiquidity{}

	list, err := o.Cache.ZRevRange(key, start, stop)
	for _, member := range list {
		coinLiq := &context.CoinLiquidity{}
		json.Unmarshal([]byte(fmt.Sprintf("%v", member.Member)), coinLiq)
		coinLiqs = append(coinLiqs, coinLiq)
	}
	return coinLiqs, err
}

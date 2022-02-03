package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

func MakeLogKeyOfPoint(appID, pointID int64, candleType string) string {
	return config.GetInstance().DBPrefix + ":LOG-APP-POINT:" + strconv.FormatInt(appID, 10) + "-" + strconv.FormatInt(pointID, 10) + "-" + candleType
}

func (o *DB) HSetLogOfPoint(key string, field string, value interface{}) error {
	return o.Cache.HSet(key, field, value)
}

func (o *DB) HGetLogOfPoint(key string, field string, value interface{}) error {
	return o.Cache.HGet(key, field, value)
}

func (o *DB) ZADDLogOfPoint(key string, score int64, value interface{}) error {
	return o.Cache.ZAdd(key, basedb.Z{Score: float64(score), Member: value})
}

func (o *DB) ZRemRangeByScore(key, min, max string) (int64, error) {
	return o.Cache.ZRemRangeByScore(key, min, max)
}

func (o *DB) ZRangeLogOfPoint(key string, start, stop int64) ([]*context.PointLiquidity, error) {

	pointLiqs := []*context.PointLiquidity{}

	list, err := o.Cache.ZRange(key, start, stop)
	for _, member := range list {
		pointLiq := &context.PointLiquidity{}
		json.Unmarshal([]byte(fmt.Sprintf("%v", member.Member)), pointLiq)
		pointLiqs = append(pointLiqs, pointLiq)
	}
	return pointLiqs, err
}

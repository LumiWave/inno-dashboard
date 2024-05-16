package model

import (
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
)

const (
	News = "NEWS"
)

func MakeKey(key string) string {
	return config.GetInstance().DBPrefix + ":" + key
}

func MakeNewsKey(pageSize string, pageOffset string) string {
	return MakeKey(News) + ":" + pageSize + ":" + pageOffset
}

func (o *DB) SetCacheNews(pageSize string, pageOffset string, newsList *context.ResNewsList) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}

	key := MakeNewsKey(pageSize, pageOffset)
	return o.Cache.Set(key, newsList, -1)
}

func (o *DB) GetCacheNews(pageSize string, pageOffset string) (*context.ResNewsList, error) {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}
	key := MakeNewsKey(pageSize, pageOffset)

	res := &context.ResNewsList{}
	err := o.Cache.Get(key, res)
	if err != nil {
		return nil, err
	} else {
		return res, err
	}
}
func (o *DB) DelCacheNews() error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}
	key := MakeKey(News) + ":*"
	return o.Cache.Truncate(key)
}

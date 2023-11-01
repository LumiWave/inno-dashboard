package model

import (
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

// 코인 마이그레이션 관련
func MakeMIGCoinTransferKey(auid, coinID int64) string {
	return config.GetInstance().DBPrefix + ":IMG:" + strconv.FormatInt(auid, 10) + ":COIN:" + strconv.FormatInt(coinID, 10)
}

func (o *DB) SetCacheIMGCoinTransfer(auid int64, coin *context.MIGCoin) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}

	key := MakeMIGCoinTransferKey(auid, coin.CoinID)
	return o.Cache.Set(key, coin, -1)
}

func (o *DB) DelCacheIMGCoinTransfer(auid int64, coin *context.MIGCoin) error {
	key := MakeMIGCoinTransferKey(auid, coin.CoinID)
	return o.Cache.Del(key)
}

///////////////////////

// NFT 마이그레이션 관련
func MakeMIGNFTTransferKey(auid, coinID, nftID int64, objectID string) string {
	key := config.GetInstance().DBPrefix + ":IMG:" + strconv.FormatInt(auid, 10) + ":NFT:" + strconv.FormatInt(coinID, 10) + ":"
	if nftID != 0 {
		key += strconv.FormatInt(nftID, 10)
	} else {
		key += objectID
	}
	return key
}

func (o *DB) SetCacheIMGNFTransfer(auid int64, coin *context.MIGNFT) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}

	key := MakeMIGNFTTransferKey(auid, coin.CoinID, coin.NFTID, "")
	return o.Cache.Set(key, coin, -1)
}

func (o *DB) DelCacheIMGNFTTransfer(auid int64, coin *context.MIGNFT) error {
	key := MakeMIGNFTTransferKey(auid, coin.CoinID, coin.NFTID, "")
	return o.Cache.Del(key)
}

///////////////////////

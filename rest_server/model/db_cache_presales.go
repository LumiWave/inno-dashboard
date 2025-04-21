package model

import (
	"strconv"
	"time"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
)

const (
	PreSales       = "PRESALES"
	PreSalesClaims = "PRESALESCLAIMS"
	PreSalesKeys   = "PRE"
	PreSalesRank   = "PRESALESRANK"
)

func MakePreSalesLockKey(salesID int64, claimID int64) string {
	return config.GetInstance().DBPrefix + "-PRESALES-" + strconv.FormatInt(salesID, 10) + "-" + strconv.FormatInt(claimID, 10) + "-lock"
}

func (o *DB) GetCachePreSales() ([]*context.PreSales, error) {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}
	fullkey := MakeKey(PreSales)
	cacheData := []*context.PreSales{}
	err := o.Cache.Get(fullkey, &cacheData)

	return cacheData, err
}

func (o *DB) SetCachePreSales(params []*context.PreSales) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}
	fullkey := MakeKey(PreSales)
	return o.Cache.Set(fullkey, params, time.Duration(10)*time.Minute)
}

func (o *DB) DelCacheAllPreSales() error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
	}
	fullkey := MakeKey(PreSales)
	o.Cache.Del(fullkey)
	fullkey = MakeKey(PreSalesClaims)
	o.Cache.Del(fullkey)
	return nil
}

// func (o *DB) GetCachePreSalesClaims() ([]*context.PreSalesClaim, error) {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	fullkey := MakeKey(PreSalesClaims)
// 	cacheData := []*context.PreSalesClaim{}
// 	err := o.Cache.Get(fullkey, &cacheData)

// 	return cacheData, err
// }

// func (o *DB) SetCachePreSalesClaims(params []*context.PreSalesClaim) error {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	fullkey := MakeKey(PreSalesClaims)
// 	return o.Cache.Set(fullkey, params, time.Duration(10)*time.Minute)
// }

// func (o *DB) SetCacheUserPreSalesInfo(walletAddress string, params *context.ProductPurchase) error {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	key := MakeKey(PreSales + ":" + walletAddress)
// 	o.SetCacheUserSaleInfoKey(PreSalesKeys+":"+walletAddress, context.PurchaseType_PreSales)
// 	//address에서 사전구매는 PRE:를 붙여 하나의 캐시데이터에서 관리를한다. 스케쥴러를 이거때문에 하나 더돌리기엔 오바인거같다

// 	return o.Cache.Set(key, params, -1)
// }

// func (o *DB) GetCacheUserPreSalesInfo(walletAddress string) (*context.ProductPurchase, error) {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	key := MakeMarketKey(PreSales + ":" + walletAddress)

// 	res := &context.ProductPurchase{}
// 	err := o.Cache.Get(key, res)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return res, err
// 	}
// }

// func (o *DB) DeleteCacheUserPreSalesInfo(walletAddress string) error {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	key := MakeMarketKey(PreSales + ":" + walletAddress)

// 	o.DeleteCacheUserSaleInfoKey(PreSalesKeys + ":" + walletAddress)

// 	return o.Cache.Del(key)
// }

// func (o *DB) DeleteCacheUserPreSalesInfoKey(walletAddress string) error {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}

// 	return o.Cache.HDel(MakeMarketKey(ProductSaleKeys), PreSalesKeys+":"+walletAddress)
// }

// func (o *DB) GetCachePreSalesLeaderBoard() (*context.PreSalesLeaderBoard, error) {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	fullkey := MakeMarketKey(PreSalesRank)
// 	cacheData := &context.PreSalesLeaderBoard{}
// 	err := o.Cache.Get(fullkey, &cacheData)

// 	return cacheData, err
// }

// func (o *DB) SetCachePreSalesLeaderBoard(params *context.PreSalesLeaderBoard) error {
// 	if !o.Cache.Enable() {
// 		log.Warnf("redis disable")
// 	}
// 	fullkey := MakeMarketKey(PreSalesRank)
// 	//todo::일단 ttl 조절해야함 일단 1분
// 	return o.Cache.Set(fullkey, params, time.Duration(1*time.Minute))
// }

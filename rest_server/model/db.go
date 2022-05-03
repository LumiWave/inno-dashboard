package model

import (
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type PointDB struct {
	DatabaseID   int64
	DatabaseName string
	ServerName   string
}

type Point struct {
	PointIds []int64
}

type AppCoin struct {
	AppID int64 `json:"app_id"`
	context.CoinInfo
}

type AppInfo struct {
	AppId   int64  `json:"app_id,omitempty"`
	AppName string `json:"app_name,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type DB struct {
	MssqlAccountAll  *basedb.Mssql
	MssqlAccountRead *basedb.Mssql
	MssqlLogRead     *basedb.Mssql
	Cache            *basedb.CacheV8

	MssqlPoints map[int64]*basedb.Mssql

	ScanPointsMap map[int64]*context.PointInfo // 전체 포인트 종류 1 : key PointId
	ScanPoints    context.PointList            // 전체 포인트 종류 2

	AppPointsMap map[int64]*context.AppPointInfo // 전체 app과 포인트 1 : key appId
	AppPoints    context.AppPoints               // 전체 app과 포인트 2

	AppCoins map[int64][]*AppCoin // 전체 app에 속한 CoinID 정보 : key AppId

	CoinsMap map[int64]*context.CoinInfo // 전체 coin 정보 1 : key CoinId
	Coins    context.CoinList            // 전체 coin 정보 2

	BaseCoinMapByCoinID map[int64]*context.BaseCoinInfo  // 전체 base coin 정보 : key coin symbol
	BaseCoinMapBySymbol map[string]*context.BaseCoinInfo // 전체 base coin 정보 : key coin symbol
	BaseCoins           context.BaseCoinList

	SwapAbleMap map[int64]*context.Swapable // 전체 스왑 가능한 정보 1 : key appID
	SwapAble    []*context.Swapable         // 전체 스왑 가능한 정보 2

	RedSync *redsync.Redsync
}

var gDB *DB

func GetDB() *DB {
	return gDB
}

func InitDB(conf *config.ServerConfig) error {
	cache := basedb.GetCacheV8(&conf.Cache)
	gDB = &DB{
		Cache: cache,
	}

	pool := goredis.NewPool(cache.GetDB().RedisClient())
	gDB.RedSync = redsync.New(pool)

	if err := ConnectAllDB(conf); err != nil {
		return err
	}

	go func() {
		for {
			timer := time.NewTimer(5 * time.Second)
			<-timer.C
			timer.Stop()

			// DB ping 체크 후 오류 시 재 연결
			if db := CheckPingDB(gDB.MssqlAccountAll, conf.MssqlDBAccountAll); db != nil {
				gDB.MssqlAccountAll = db
			}

			if db := CheckPingDB(gDB.MssqlAccountRead, conf.MssqlDBAccountRead); db != nil {
				gDB.MssqlAccountRead = db
			}

			if db := CheckPingDB(gDB.MssqlLogRead, conf.MssqlDBLogRead); db != nil {
				gDB.MssqlLogRead = db
			}
		}
	}()

	LoadDBPoint(conf)

	go gDB.ListenSubscribeEvent()

	return nil
}

func LoadDBPoint(conf *config.ServerConfig) {
	gDB.ScanPointsMap = make(map[int64]*context.PointInfo)
	gDB.AppCoins = make(map[int64][]*AppCoin)
	gDB.AppPointsMap = make(map[int64]*context.AppPointInfo)
	gDB.CoinsMap = make(map[int64]*context.CoinInfo)
	gDB.SwapAbleMap = make(map[int64]*context.Swapable)
	gDB.BaseCoinMapByCoinID = make(map[int64]*context.BaseCoinInfo)
	gDB.BaseCoinMapBySymbol = make(map[string]*context.BaseCoinInfo)

	gDB.GetPointList()
	gDB.GetBaseCoins()
	gDB.GetAppCoins()
	gDB.GetCoins()
	gDB.GetApps()
	gDB.GetAppPoints()
	gDB.GetScanExchangeGoods()

	if conf.App.LiquidityUpdate {
		gDB.LoadFullPointLiquidity(1000, true)
		gDB.LoadFullCoinLiquidity(1000, true)
		gDB.UpdateLiquidity()
		gDB.UpdateCoinFee()
	}
}

func MakeDbError(resp *base.BaseResponse, errCode int, err error) {
	resp.Return = errCode
	resp.Message = resultcode.ResultCodeText[errCode] + " : " + err.Error()
}

func (o *DB) ConnectDB(conf *baseconf.DBAuth) (*basedb.Mssql, error) {
	port, _ := strconv.ParseInt(conf.Port, 10, 32)
	mssqlDB, err := basedb.NewMssql(conf.Database, "", conf.ID, conf.Password, conf.Host, int(port),
		conf.ApplicationIntent, conf.Timeout, conf.ConnectRetryCount, conf.ConnectRetryInterval)

	if err != nil {
		log.Errorf("err: %v, val: %v, %v, %v, %v, %v, %v",
			err, conf.Host, conf.ID, conf.Password, conf.Database, conf.PoolSize, conf.IdleSize)
		return nil, err
	}
	idleSize, _ := strconv.ParseInt(conf.IdleSize, 10, 32)
	mssqlDB.GetDB().SetMaxOpenConns(int(idleSize))
	mssqlDB.GetDB().SetMaxIdleConns(int(idleSize))

	return mssqlDB, nil
}

func ConnectAllDB(conf *config.ServerConfig) error {
	var err error
	gDB.MssqlAccountAll, err = gDB.ConnectDB(&conf.MssqlDBAccountAll)
	if err != nil {
		return err
	}

	gDB.MssqlAccountRead, err = gDB.ConnectDB(&conf.MssqlDBAccountRead)
	if err != nil {
		return err
	}

	gDB.MssqlLogRead, err = gDB.ConnectDB(&conf.MssqlDBLogRead)
	if err != nil {
		return err
	}
	return nil
}

func ChangeTime(strTime string) *time.Time {
	if len(strTime) == 0 {
		return nil
	}
	var baseDate *time.Time
	t, err := time.Parse("2006-01-02T15:04:05Z", strTime)
	if err != nil {
		log.Errorf("time.Parse [err%v]", err)
		return nil
	} else {
		baseDate = &t
	}
	if t.IsZero() {
		baseDate = nil
	}
	return baseDate
}

func ChangeHourTime(strTime string) *time.Time {
	if len(strTime) == 0 {
		return nil
	}
	var baseDate *time.Time
	t, err := time.Parse("Jan 02 2006  3:04PM", strTime)
	if err != nil {
		log.Errorf("time.Parse [err%v]", err)
		return nil
	} else {
		baseDate = &t
	}
	if t.IsZero() {
		baseDate = nil
	}
	return baseDate
}

func ChangeDayTime(strTime string) *time.Time {
	if len(strTime) == 0 {
		return nil
	}
	var baseDate *time.Time
	t, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		log.Errorf("time.Parse [err%v]", err)
		return nil
	} else {
		baseDate = &t
	}
	if t.IsZero() {
		baseDate = nil
	}
	return baseDate
}

func ChangeBaseTime(strTime string) *time.Time {
	if len(strTime) == 0 {
		return nil
	}
	var baseDate *time.Time
	t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strTime)
	if err != nil {
		log.Errorf("time.Parse [err%v]", err)
		return nil
	} else {
		baseDate = &t
	}
	if t.IsZero() {
		baseDate = nil
	}
	return baseDate
}

func CheckPingDB(db *basedb.Mssql, conf baseconf.DBAuth) *basedb.Mssql {
	// 연결이 안되어있거나, DB Connection이 끊어진 경우에는 재연결 시도
	if db == nil || !db.Connection.IsConnect {
		var err error
		newDB, err := gDB.ConnectDB(&conf)
		if err == nil {
			log.Errorf("connect DB OK")
		}
		return newDB
	}

	// 연결이 되어있는 상태면 ping
	if db.Connection.IsConnect {
		if err := db.GetDB().Ping(); err != nil {
			// 재시도 횟수
			db.Connection.RetryCount += 1
			log.Errorf("%v DB Ping err RetryCount(%v)", conf.Database, db.Connection.RetryCount)
			// ping 2회 시도해도 안되면 close
			if db.Connection.RetryCount >= 2 {
				db.Connection.IsConnect = false
				// DB Close
				if err = db.GetDB().Close(); err == nil {
					log.Errorf("DB Closed (RetryCount >=2)")
				}
			}
		}
	}
	return nil
}

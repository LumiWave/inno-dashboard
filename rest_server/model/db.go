package model

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
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
	Mysql        *basedb.Mysql
	MssqlAccount *basedb.Mssql
	Cache        *basedb.Cache

	MssqlPoints map[int64]*basedb.Mssql

	ScanPointsMap map[int64]context.PointInfo // 전체 포인트 종류 1 : key PointId
	ScanPoints    context.PointList           // 전체 포인트 종류 2

	AppPointsMap map[int64]*context.AppPointInfo // 전체 app과 포인트 1
	AppPoints    context.AppPoints               // 전체 app과 포인트 2

	AppCoins map[int64][]*AppCoin // 전체 app에 속한 CoinID 정보 : key AppId

	CoinsMap map[int64]*context.CoinInfo // 전체 coin 정보 1 : key CoinId
	Coins    context.CoinList            // 전체 coin 정보 2
}

var gDB *DB

func SetDB(db *basedb.Mssql, cache *basedb.Cache, pointdbs map[int64]*basedb.Mssql) {
	gDB = &DB{
		MssqlAccount: db,
		Cache:        cache,
		MssqlPoints:  pointdbs,
	}
}

func SetDBPoint(pointdbs map[int64]*basedb.Mssql) {
	gDB.ScanPointsMap = make(map[int64]context.PointInfo)
	gDB.AppCoins = make(map[int64][]*AppCoin)
	gDB.AppPointsMap = make(map[int64]*context.AppPointInfo)
	gDB.CoinsMap = make(map[int64]*context.CoinInfo)
	gDB.MssqlPoints = pointdbs

	gDB.GetPointList()
	gDB.GetAppCoins()
	gDB.GetCoins()
	gDB.GetApps()
	gDB.GetAppPoints()
}

func GetDB() *DB {
	return gDB
}

func MakeDbError(resp *base.BaseResponse, errCode int, err error) {
	resp.Return = errCode
	resp.Message = resultcode.ResultCodeText[errCode] + " : " + err.Error()
}

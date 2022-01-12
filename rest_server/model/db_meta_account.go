package model

import (
	originCtx "context"
	"database/sql"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_Scan_DatabaseServers   = "[dbo].[USPAU_Scan_DatabaseServers]"
	USPAU_Scan_Points            = "[dbo].[USPAU_Scan_Points]"
	USPAU_Scan_ApplicationCoins  = "[dbo].[USPAU_Scan_ApplicationCoins]"
	USPAU_Scan_ApplicationPoints = "[dbo].[USPAU_Scan_ApplicationPoints]"
	USPAU_Scan_Applications      = "[dbo].[USPAU_Scan_Applications]"
	USPAU_Scan_Coins             = "[dbo].[USPAU_Scan_Coins]"
)

// point database 리스트 요청
func (o *DB) GetPointDatabases() (map[int64]*PointDB, error) {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_DatabaseServers, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return nil, err
	}

	defer rows.Close()

	pointdbs := make(map[int64]*PointDB)

	pointdb := new(PointDB)
	for rows.Next() {
		rows.Scan(&pointdb.DatabaseID, &pointdb.DatabaseName, &pointdb.ServerName)
		pointdbs[pointdb.DatabaseID] = pointdb
	}

	return pointdbs, nil
}

// point 전체 list
func (o *DB) GetPointList() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_Points, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.ScanPointsMap = make(map[int64]context.PointInfo)
	o.ScanPoints.Points = nil

	var pointId int64
	var pointName, iconPath string
	for rows.Next() {
		if err := rows.Scan(&pointId, &pointName, &iconPath); err == nil {
			info := context.PointInfo{
				PointId:   pointId,
				PointName: pointName,
				IconUrl:   iconPath,
			}
			o.ScanPointsMap[pointId] = info
			o.ScanPoints.Points = append(o.ScanPoints.Points, info)
		} else {
			log.Warn("GetPointList err :", err)
		}
	}

	return nil
}

// 전체 app coinid list
func (o *DB) GetAppCoins() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_ApplicationCoins, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.AppCoins = make(map[int64][]*AppCoin)
	for rows.Next() {
		appCoin := &AppCoin{}
		if err := rows.Scan(&appCoin.AppID, &appCoin.CoinId); err == nil {
			o.AppCoins[appCoin.AppID] = append(o.AppCoins[appCoin.AppID], appCoin)
		}
	}

	return nil
}

// 전체 coin info list
func (o *DB) GetCoins() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_Coins, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.CoinsMap = make(map[int64]*context.CoinInfo)
	o.Coins.Coins = nil

	for rows.Next() {
		coin := &context.CoinInfo{}
		if err := rows.Scan(&coin.CoinId, &coin.CoinSymbol, &coin.ContractAddress, &coin.IconUrl); err == nil {
			o.Coins.Coins = append(o.Coins.Coins, coin)
			o.CoinsMap[coin.CoinId] = coin
		}
	}

	for _, appCoins := range o.AppCoins {
		for _, appCoin := range appCoins {
			for _, coin := range o.Coins.Coins {
				if appCoin.CoinId == coin.CoinId {
					appCoin.CoinSymbol = coin.CoinSymbol
					appCoin.ContractAddress = coin.ContractAddress
					appCoin.IconUrl = coin.IconUrl
					break
				}
			}
		}
	}

	return nil
}

// 전체 app list 조회
func (o *DB) GetApps() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_Applications, &rs)
	if err != nil {
		log.Error("GetApps QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.AppPointsMap = make(map[int64]*context.AppPointInfo)
	for rows.Next() {
		appInfo := &context.AppPointInfo{}
		if err := rows.Scan(&appInfo.AppId, &appInfo.AppName, &appInfo.IconUrl); err == nil {
			o.AppPointsMap[appInfo.AppId] = appInfo
		}
	}

	return nil
}

// 전체 app 과 포인트 list 조회
func (o *DB) GetAppPoints() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccount.GetDB().QueryContext(originCtx.Background(), USPAU_Scan_ApplicationPoints, &rs)
	if err != nil {
		log.Error("GetAppPoints QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	var daliyLimiteQuantity sql.NullInt64
	appPointInfo := &context.AppPointInfo{}
	for rows.Next() {
		if err := rows.Scan(&appPointInfo.AppId, &appPointInfo.PointId, &daliyLimiteQuantity); err == nil {
			o.AppPointsMap[appPointInfo.AppId].PointId = appPointInfo.PointId
			temp := o.ScanPointsMap[appPointInfo.PointId]
			temp.DaliyLimitedQuantity = daliyLimiteQuantity.Int64
			o.AppPointsMap[appPointInfo.AppId].Points = append(o.AppPointsMap[appPointInfo.AppId].Points, &temp)
		}
	}

	o.AppPoints.Apps = nil

	for _, appPoint := range o.AppPointsMap {
		o.AppPoints.Apps = append(o.AppPoints.Apps, appPoint)
	}

	return nil
}

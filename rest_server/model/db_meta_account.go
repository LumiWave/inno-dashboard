package model

import (
	originCtx "context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_Scan_DatabaseServers     = "[dbo].[USPAU_Scan_DatabaseServers]"
	USPAU_Scan_Points              = "[dbo].[USPAU_Scan_Points]"
	USPAU_Scan_ApplicationCoins    = "[dbo].[USPAU_Scan_ApplicationCoins]"
	USPAU_Scan_ApplicationPoints   = "[dbo].[USPAU_Scan_ApplicationPoints]"
	USPAU_Scan_Applications        = "[dbo].[USPAU_Scan_Applications]"
	USPAU_Scan_Coins               = "[dbo].[USPAU_Scan_Coins]"
	USPAU_Scan_BaseCoins           = "[dbo].[USPAU_Scan_BaseCoins]"
	USPAU_Scan_WalletTypes         = "[dbo].[USPAU_Scan_WalletTypes]"
	USPAU_Scan_BaseCoinWalletTypes = "[dbo].[USPAU_Scan_BaseCoinWalletTypes]"
)

// point database 리스트 요청
func (o *DB) GetPointDatabases() (map[int64]*PointDB, error) {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_DatabaseServers, &rs)
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
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_Points, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.ScanPointsMap = make(map[int64]*context.PointInfo)
	o.ScanPoints.Points = nil

	var pointId, dailyLimitExchangeAcqQuantity int64
	var pointName, iconPath string
	for rows.Next() {
		if err := rows.Scan(&pointId, &pointName, &iconPath, &dailyLimitExchangeAcqQuantity); err == nil {
			info := &context.PointInfo{
				PointId:                       pointId,
				PointName:                     pointName,
				IconUrl:                       iconPath,
				DailyLimitExchangeAcqQuantity: dailyLimitExchangeAcqQuantity,
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
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_ApplicationCoins, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.AppCoins = make(map[int64][]*AppCoin)
	for rows.Next() {
		appCoin := &AppCoin{}
		if err := rows.Scan(&appCoin.AppID, &appCoin.CoinId, &appCoin.BaseCoinID); err == nil {
			o.AppCoins[appCoin.AppID] = append(o.AppCoins[appCoin.AppID], appCoin)
		}
	}

	return nil
}

// 전체 coin info list
func (o *DB) GetCoins() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_Coins, &rs)
	if err != nil {
		log.Error("USPAU_Scan_Coins QueryContext err : ", err)
		return err
	}

	if rows != nil {
		defer rows.Close()
	}

	o.CoinsMap = make(map[int64]*context.CoinInfo)
	o.Coins.Coins = nil

	for rows.Next() {
		coin := &context.CoinInfo{}
		var customProperties string
		if err := rows.Scan(&coin.CoinId,
			&coin.BaseCoinID,
			&coin.CoinName,
			&coin.CoinSymbol,
			&coin.ContractAddress,
			&coin.Decimal,
			&coin.ExplorePath,
			&coin.IconUrl,
			&coin.DailyLimitExchangeAcqQuantity,
			&coin.ExchangeFees,
			&coin.IsRechargeable,
			&coin.RechargeURL,
			&customProperties); err == nil {
			if err := json.Unmarshal([]byte(customProperties), &coin.CustomProperties); err != nil {
				log.Error("USPAU_Scan_Coins customProperties Unmarshal err : ", err)
			}
			o.Coins.Coins = append(o.Coins.Coins, coin)
			o.CoinsMap[coin.CoinId] = coin
		} else {
			log.Errorf("USPAU_Scan_Coins Scan error : %v", err)
		}
	}

	for _, appCoins := range o.AppCoins {
		for _, appCoin := range appCoins {
			for _, coin := range o.Coins.Coins {
				if appCoin.CoinId == coin.CoinId {
					appCoin.CoinName = coin.CoinName
					appCoin.CoinSymbol = coin.CoinSymbol
					appCoin.ContractAddress = coin.ContractAddress
					appCoin.IconUrl = coin.IconUrl
					appCoin.ExchangeFees = coin.ExchangeFees
					break
				}
			}
		}
	}

	return nil
}

// 전체 base coin list 조회
func (o *DB) GetBaseCoins() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_BaseCoins, &rs)
	if err != nil {
		log.Error("QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.BaseCoinMapByCoinID = make(map[int64]*context.BaseCoinInfo)
	o.BaseCoinMapBySymbol = make(map[string]*context.BaseCoinInfo)
	o.BaseCoins.Coins = nil
	for rows.Next() {
		baseCoin := &context.BaseCoinInfo{}
		if err := rows.Scan(&baseCoin.BaseCoinID, &baseCoin.BaseCoinName, &baseCoin.BaseCoinSymbol, &baseCoin.IsUsedParentWallet); err == nil {
			o.BaseCoinMapByCoinID[baseCoin.BaseCoinID] = baseCoin
			o.BaseCoinMapBySymbol[baseCoin.BaseCoinSymbol] = baseCoin
			o.BaseCoins.Coins = append(o.BaseCoins.Coins, baseCoin)
		}
	}

	return nil
}

// 전체 app list 조회
func (o *DB) GetApps() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_Applications, &rs)
	if err != nil {
		log.Error("GetApps QueryContext err : ", err)
		return err
	}

	defer rows.Close()

	o.AppPointsMap = make(map[int64]*context.AppPointInfo)
	for rows.Next() {
		appInfo := &context.AppPointInfo{}
		if err := rows.Scan(&appInfo.AppId, &appInfo.AppName, &appInfo.IconUrl,
			&appInfo.GooglePlayPath, &appInfo.AppleStorePath, &appInfo.BrandingPagePath); err == nil {
			o.AppPointsMap[appInfo.AppId] = appInfo
		}
	}

	return nil
}

// 전체 app 과 포인트 list 조회
func (o *DB) GetAppPoints() error {
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(originCtx.Background(), USPAU_Scan_ApplicationPoints, &rs)
	if err != nil {
		log.Errorf("USPAU_Scan_ApplicationPoints QueryContext error : %v", err)
		return err
	}

	defer rows.Close()

	o.AppPoints.Apps = nil

	var appId, pointId, daliyLimitAcqQuantity sql.NullInt64
	//var exchangeRatio sql.NullFloat64
	for rows.Next() {
		if err := rows.Scan(&appId, &pointId, &daliyLimitAcqQuantity); err == nil {
			temp := o.ScanPointsMap[pointId.Int64]
			//temp.ExchangeRatio = exchangeRatio.Float64
			//temp.MinExchangeQuantity = minExchangeQuantity.Int64
			temp.DaliyLimitAcqQuantity = daliyLimitAcqQuantity.Int64
			//temp.DailyLimitedAcqExchangeQuantity = dailyLimitedAcqExchangeQuantity.Int64

			o.AppPointsMap[appId.Int64].Points = append(o.AppPointsMap[appId.Int64].Points, temp)
			o.AppPoints.Apps = append(o.AppPoints.Apps, o.AppPointsMap[appId.Int64])
		}
	}

	if rs != 1 {
		log.Errorf("%v returnvalue error : %v", USPAU_Scan_ApplicationPoints, rs)
		return errors.New(USPAU_Scan_ApplicationPoints + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return nil
}

func (o *DB) USPAU_Scan_WalletTypes() error {
	ProcName := USPAU_Scan_WalletTypes
	var rs orginMssql.ReturnStatus

	rows, err := o.MssqlAccountRead.GetDB().QueryContext(originCtx.Background(), ProcName, &rs)
	if err != nil {
		log.Errorf(ProcName+"QueryContext err : %v", err)
		return err
	}

	defer rows.Close()

	o.WalletTypeMap = make(map[int64]*context.WalletType)
	o.WalletTypes = context.WalletTypeList{
		WalletTypes: make([]*context.WalletType, 0),
	}
	for rows.Next() {
		var WalletTypeID int64
		var WalletTypeName string
		if err := rows.Scan(&WalletTypeID, &WalletTypeName); err == nil {
			wallet := &context.WalletType{WalletName: WalletTypeName, WalletTypeID: WalletTypeID}
			o.WalletTypeMap[WalletTypeID] = wallet
			o.WalletTypes.WalletTypes = append(o.WalletTypes.WalletTypes, wallet)
		} else {
			log.Errorf("Scan error : %v", err)
		}
	}

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return nil
}
func (o *DB) USPAU_Scan_BaseCoinWalletTypes() error {
	ProcName := USPAU_Scan_BaseCoinWalletTypes
	var rs orginMssql.ReturnStatus

	rows, err := o.MssqlAccountRead.GetDB().QueryContext(originCtx.Background(), ProcName, &rs)
	if err != nil {
		log.Errorf(ProcName+"QueryContext err : %v", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var BaseCoinID int64
		var WalletTypeID int64
		if err := rows.Scan(&BaseCoinID, &WalletTypeID); err == nil {
			if _, ok := o.BaseCoinMapByCoinID[BaseCoinID]; ok {
				if o.BaseCoinMapByCoinID[BaseCoinID].AllowWalletTypes == nil {
					o.BaseCoinMapByCoinID[BaseCoinID].AllowWalletTypes = make([]int64, 0)
				}
				o.BaseCoinMapByCoinID[BaseCoinID].AllowWalletTypes = append(o.BaseCoinMapByCoinID[BaseCoinID].AllowWalletTypes, WalletTypeID)
			}
		} else {
			log.Errorf("Scan error : %v", err)
		}
	}

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return nil
}

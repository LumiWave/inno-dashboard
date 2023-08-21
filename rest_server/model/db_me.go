package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_GetList_AccountCoins   = "[dbo].[USPAU_GetList_AccountCoins]"
	USPAU_GetList_AccountPoints  = "[dbo].[USPAU_GetList_AccountPoints]"
	USPAU_GetList_Members        = "[dbo].[USPAU_GetList_Members]"
	USPAU_GetList_AccountWallets = "[dbo].[USPAU_GetList_AccountWallets]"
	USPAU_Cnct_AccountWallets    = "[dbo].[USPAU_Cnct_AccountWallets]"
	USPAU_Dscnct_AccountWallets  = "[dbo].[USPAU_Dscnct_AccountWallets]"
)

// 계정 코인 조회
func (o *DB) USPAU_GetList_AccountCoins(auid int64) ([]*context.MeCoin, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), USPAU_GetList_AccountCoins,
		sql.Named("AUID", auid),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Error("USPAU_GetList_AccountCoins QueryContext err : ", err)
		return nil, err
	}

	meCoinList := []*context.MeCoin{}
	for rows.Next() {
		meCoin := &context.MeCoin{}
		if err := rows.Scan(&meCoin.CoinID,
			&meCoin.BaseCoinID,
			&meCoin.WalletAddress,
			&meCoin.Quantity,
			&meCoin.TodayAcqQuantity,
			&meCoin.TodayCnsmQuantity,
			&meCoin.TodayAcqExchangeQuantity,
			&meCoin.TodayCnsmExchangeQuantity,
			&meCoin.ResetDate); err != nil {
			log.Errorf("USPAU_GetList_AccountCoins Scan error %v", err)
			return nil, err
		} else {
			meCoin.CoinSymbol = o.CoinsMap[meCoin.CoinID].CoinSymbol
			meCoinList = append(meCoinList, meCoin)
		}
	}

	if returnValue != 1 {
		log.Errorf("USPAU_GetList_AccountCoins returnvalue error : %v", returnValue)
		return nil, errors.New("USPAU_GetList_AccountCoins returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return meCoinList, nil
}

// 계정 포인트 조회
func (o *DB) USPAU_GetList_AccountPoints(auid, muid int64) ([]*context.MePoint, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), USPAU_GetList_AccountPoints,
		sql.Named("AUID", auid),
		sql.Named("MUID", muid),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Error("USPAU_GetList_AccountPoints QueryContext err : %v", err)
		return nil, err
	}

	var mePointList []*context.MePoint

	for rows.Next() {
		mePoint := context.MePoint{}
		if err := rows.Scan(&mePoint.AppID,
			&mePoint.PointID,
			&mePoint.TodayAcqQuantity,
			&mePoint.TodayCnsmQuantity,
			&mePoint.TodayAcqExchangeQuantity,
			&mePoint.TodayCnsmExchangeQuantity,
			&mePoint.ResetDate); err != nil {
			log.Errorf("USPAU_GetList_AccountPoints Scan error : %v", err)
			return nil, err
		} else {
			mePointList = append(mePointList, &mePoint)
		}
	}

	if returnValue != 1 {
		log.Errorf("USPAU_GetList_AccountPoints returnvalue error : %v", returnValue)
		return nil, errors.New("USPAU_GetList_AccountPoints returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return mePointList, nil
}

// 계정 앱 회원 조회
func (o *DB) USPAU_GetList_Members(auid int64) ([]*context.Member, map[int64]*context.Member, error) {
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), USPAU_GetList_Members,
		sql.Named("AUID", auid),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Error("USPAU_GetList_Members QueryContext err : %v", err)
		return nil, nil, err
	}

	var memberList []*context.Member
	memberMap := make(map[int64]*context.Member)

	for rows.Next() {
		member := context.Member{}
		if err := rows.Scan(&member.MUID, &member.AppID, &member.DatabaseID); err != nil {
			log.Errorf("USPAU_GetList_Members Scan error : %v", err)
			return nil, nil, err
		} else {
			memberMap[member.AppID] = &member
			memberList = append(memberList, &member)
		}
	}

	if returnValue != 1 {
		log.Errorf("USPAU_GetList_Members returnvalue error : %v", returnValue)
		return nil, nil, errors.New("USPAU_GetList_Members returnvalue error " + strconv.Itoa(int(returnValue)))
	}
	return memberList, memberMap, nil
}

// 등록된 지갑정보
func (o *DB) USPAU_GetList_AccountWallets(auid int64) ([]*context.DBWalletRegist, error) {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_GetList_AccountWallets
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		sql.Named("AUID", auid),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return nil, err
	}

	walletRegists := make([]*context.DBWalletRegist, 0)

	for rows.Next() {
		data := &context.DBWalletRegist{}
		if err := rows.Scan(&data.WalletID, &data.BaseCoinID, &data.WalletAddress, &data.ConnectionStatus, &data.ModifiedDT); err != nil {
			log.Errorf("%s Scan error : %v", proc, err)
			return nil, err
		} else {
			walletRegists = append(walletRegists, data)
		}
	}

	if returnValue != 1 {
		log.Errorf("%s returnvalue error : %v", proc, returnValue)
		return nil, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return walletRegists, nil
}

// 지갑등록
func (o *DB) USPAU_Cnct_AccountWallets(auid int64, baseCoinID int64, walletAddress string) (int, error) {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Cnct_AccountWallets
	rows, err := o.MssqlAccountAll.QueryContext(contextR.Background(), proc,
		sql.Named("AUID", auid),
		sql.Named("BaseCoinID", baseCoinID),
		sql.Named("WalletAddress", walletAddress),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return 1, err
	}

	if returnValue != 1 {
		log.Errorf("%s returnvalue error : %v", proc, returnValue)
		switch returnValue {
		case 50106:
			//이미 다른지갑에 연결된 지갑주소
			return 2, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
		case 50107:
			//다른 사용자에 의해 연결된 지갑주소
			return 3, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
		default:
			return 1, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
		}
	}

	return 0, nil
}

// 지갑삭제
func (o *DB) USPAU_Dscnct_AccountWallets(auid int64, baseCoinID int64, walletAddress string) error {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Dscnct_AccountWallets
	rows, err := o.MssqlAccountAll.QueryContext(contextR.Background(), proc,
		sql.Named("AUID", auid),
		sql.Named("BaseCoinID", baseCoinID),
		sql.Named("WalletAddress", walletAddress),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return err
	}

	if returnValue != 1 {
		log.Errorf("%s returnvalue error : %v", proc, returnValue)
		return errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return nil
}

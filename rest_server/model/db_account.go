package model

import (
	contextR "context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/LumiWave/baseutil/log"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPAU_Verify_Accounts         = "[dbo].[USPAU_Verify_Accounts]"
	USPAU_Auth_Members            = "[dbo].[USPAU_Auth_Members]"
	USPAU_Get_Accounts_By_InnoUID = "[dbo].[USPAU_Get_Accounts_By_InnoUID]"
)

func (o *DB) USPAU_Verify_Accounts(innoUID string) (bool, error) {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Verify_Accounts
	isExist := false
	rows, err := o.MssqlAccountAll.QueryContext(contextR.Background(), proc,
		sql.Named("InnoUID", innoUID),
		sql.Named("IsExist", sql.Out{Dest: &isExist}),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return isExist, err
	}

	if returnValue != 1 {
		log.Errorf(USPAU_Verify_Accounts+" returnvalue error : %v", returnValue)
		return isExist, errors.New(USPAU_Verify_Accounts + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return isExist, nil
}

func (o *DB) USPAU_Auth_Members(innoUID string, appID int64) (int64, int64, error) { // muid, databaseid
	isJoined := false
	auid := int64(0)
	muid := int64(0)
	databaseID := int64(0)
	var returnValue orginMssql.ReturnStatus
	rows, err := o.MssqlAccountAll.QueryContext(contextR.Background(), USPAU_Auth_Members,
		sql.Named("InnoUID", innoUID),
		sql.Named("AppID", appID),
		sql.Named("IsJoined", sql.Out{Dest: &isJoined}),
		sql.Named("AUID", sql.Out{Dest: &auid}),
		sql.Named("MUID", sql.Out{Dest: &muid}),
		sql.Named("DatabaseID", sql.Out{Dest: &databaseID}),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if returnValue != 1 {
		return muid, databaseID, err
	}

	return muid, databaseID, err
}

func (o *DB) USPAU_Get_Accounts_By_InnoUID(innoUID string) (int64, string, string, error) {
	var returnValue orginMssql.ReturnStatus
	proc := USPAU_Get_Accounts_By_InnoUID
	auid := int64(0)
	kycResultStatus := ""
	submissionID := ""
	rows, err := o.MssqlAccountRead.QueryContext(contextR.Background(), proc,
		sql.Named("InnoUID", innoUID),
		&returnValue)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		log.Errorf("%s QueryContext error : %v", proc, err)
		return auid, kycResultStatus, submissionID, err
	}

	for rows.Next() {
		if err := rows.Scan(&auid, &kycResultStatus, &submissionID); err != nil {
			log.Errorf(proc+" Get error : %v", err)
			return auid, kycResultStatus, submissionID, err
		}
	}

	if returnValue != 1 {
		log.Errorf(proc+" returnvalue error : %v", returnValue)
		return auid, kycResultStatus, submissionID, errors.New(proc + " returnvalue error " + strconv.Itoa(int(returnValue)))
	}

	return auid, kycResultStatus, submissionID, nil
}

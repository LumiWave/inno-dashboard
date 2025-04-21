package model

import (
	originCtx "context"
	"errors"
	"strconv"

	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPPR_Scan_PreSales                         = "[dbo].[USPPR_Scan_PreSales]"
	USPPR_Scan_PreSalesClaims                   = "[dbo].[USPPR_Scan_PreSalesClaims]"
	USPPR_PrchsStrt_PreSalesClaims              = "[dbo].[USPPR_PrchsStrt_PreSalesClaims]"
	USPPR_PrchsCmplt_PreSalesClaims             = "[dbo].[USPPR_PrchsCmplt_PreSalesClaims]"
	USPPR_Mod_TransactPreSalesClaims_Coin       = "[dbo].[USPPR_Mod_TransactPreSalesClaims_Coin]"
	USPPR_Mod_TransactPreSalesClaims_NFT        = "[dbo].[USPPR_Mod_TransactPreSalesClaims_NFT]"
	USPPR_Mod_TransactPreSalesClaims_TxStatus   = "[dbo].[USPPR_Mod_TransactPreSalesClaims_TxStatus]"
	USPPR_Mod_TransactPreSalesClaims_DynamicNFT = "[dbo].[USPPR_Mod_TransactPreSalesClaims_DynamicNFT]"
)

func (o *DB) USPPR_Scan_PreSales() ([]*context.PreSales, error) {
	ProcName := USPPR_Scan_PreSales
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return nil, err
	}

	defer rows.Close()

	preSalesList := make([]*context.PreSales, 0)
	for rows.Next() {
		item := &context.PreSales{}
		if err := rows.Scan(
			&item.SalesID,
			&item.OpenStartSDT,
			&item.OpenEndSDT,
			&item.TotalExchangePointQuantity,
			&item.CurrExchangePointQuantity,
			&item.ReferralLimitCount,
			&item.ReferralRewardAppID,
			&item.ReferralRewardPointID,
			&item.ReferrerPointQuantity,
			&item.RefereePointQuantity); err != nil {
			log.Errorf(ProcName+" Get error : %v", err)
			return nil, err
		} else {
			preSalesList = append(preSalesList, item)
		}
	}

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return nil, errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}
	return preSalesList, nil
}

// func (o *DB) USPPR_Scan_PreSalesClaims() ([]*context.PreSalesClaim, error) {
// 	ProcName := USPPR_Scan_PreSalesClaims
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	preSalesClaimList := make([]*context.PreSalesClaim, 0)
// 	for rows.Next() {
// 		item := &context.PreSalesClaim{}
// 		if err := rows.Scan(&item.SalesID, &item.ClaimID, &item.DisplayArea, &item.DisplaySortOrder, &item.IsEnabled, &item.ClaimStartSDT, &item.ClaimEndSDT,
// 			&item.PaidPointQuantity, &item.RewardType, &item.NFTPackID, &item.ImageRewradID, &item.LWAPrice, &item.LWADiscountPrice, &item.SUIPrice); err != nil {
// 			log.Errorf(ProcName+" Get error : %v", err)
// 			return nil, err
// 		} else {
// 			preSalesClaimList = append(preSalesClaimList, item)
// 		}
// 	}

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return nil, errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}
// 	return preSalesClaimList, nil
// }

// // 상품 구매 시작
// func (o *DB) USPPR_PrchsStrt_PreSalesClaims(req *context.ProductPurchase) (int64 /* txID */, int64 /* nftID */, string, int64, int64, string, error) {
// 	ProcName := USPPR_PrchsStrt_PreSalesClaims
// 	var rs orginMssql.ReturnStatus
// 	var txID, nftID, dynamicNFTID, pointQuantity int64
// 	var objectID, dynamicObjectID string
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("AUID", req.AUID),
// 		sql.Named("SalesID", req.SalesID),
// 		sql.Named("ClaimID", req.ClaimID),
// 		sql.Named("BaseCoinID", req.BaseCoinID),
// 		sql.Named("WalletTypeID", req.WalletTypeID),
// 		sql.Named("WalletAddress", req.WalletAddress),
// 		sql.Named("CoinID", req.CoinID),
// 		sql.Named("AdjQuantity", req.AdjQuantity),

// 		sql.Named("NFTID", sql.Out{Dest: &nftID}),
// 		sql.Named("ObjectID", sql.Out{Dest: &objectID}),
// 		sql.Named("PointQuantity", sql.Out{Dest: &pointQuantity}),
// 		sql.Named("TxID", sql.Out{Dest: &txID}),
// 		sql.Named("DynamicNFTID", sql.Out{Dest: &dynamicNFTID}),
// 		sql.Named("DynamicObjectID", sql.Out{Dest: &dynamicObjectID}),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return txID, nftID, objectID, pointQuantity, dynamicNFTID, dynamicObjectID, err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)

// 		if int(rs) == context.DB_Return_Enough_NFT {
// 			return txID, nftID, objectID, pointQuantity, dynamicNFTID, dynamicObjectID, errors.New(strconv.Itoa(int(rs)))
// 		} else {
// 			return 0, 0, "", 0, 0, "", errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 		}
// 	}

// 	return txID, nftID, objectID, pointQuantity, dynamicNFTID, dynamicObjectID, nil
// }

// // 상품 구매 완료
// func (o *DB) USPPR_PrchsCmplt_PreSalesClaims(txID int64, completeDT string, dynamicNFTID int64, dynamicNFTWalletID int64) error {
// 	ProcName := USPPR_PrchsCmplt_PreSalesClaims
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("TxID", txID),
// 		sql.Named("CompletedDT", completeDT),
// 		sql.Named("DynamicNFTID", dynamicNFTID),
// 		sql.Named("DynamicNFTWalletID", dynamicNFTWalletID),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}

// 	return nil
// }

// func (o *DB) USPPR_Mod_TransactPreSalesClaims_Coin(txID int64, txStatus int64, txHash string, BaseCoinID int64, gasFee string) error {
// 	ProcName := USPPR_Mod_TransactPreSalesClaims_Coin
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("TxID", txID),
// 		sql.Named("TxStatus", txStatus),
// 		sql.Named("TxHash", txHash),
// 		sql.Named("BaseCoinID", BaseCoinID),
// 		sql.Named("GasFee", gasFee),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}

// 	return nil
// }

// func (o *DB) USPPR_Mod_TransactPreSalesClaims_NFT(txID int64, txStatus int64, txHash string, BaseCoinID int64, gasFee string) error {
// 	ProcName := "USPPR_Mod_TransactPreSalesClaims_NFT"
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("TxID", txID),
// 		sql.Named("TxStatus", txStatus),
// 		sql.Named("TxHash", txHash),
// 		sql.Named("BaseCoinID", BaseCoinID),
// 		sql.Named("GasFee", gasFee),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}

// 	return nil
// }

// func (o *DB) USPPR_Mod_TransactPreSalesClaims_TxStatus(txID int64, txStatus int64, BaseCoinID int64, gasFee string) error {
// 	ProcName := USPPR_Mod_TransactPreSalesClaims_TxStatus
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("TxID", txID),
// 		sql.Named("TxStatus", txStatus),
// 		sql.Named("BaseCoinID", BaseCoinID),
// 		sql.Named("GasFee", gasFee),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}

// 	return nil
// }

// func (o *DB) USPPR_Mod_TransactPreSalesClaims_DynamicNFT(txID int64, txStatus int64, txHash string, BaseCoinID int64, gasFee string) error {
// 	ProcName := USPPR_Mod_TransactPreSalesClaims_DynamicNFT
// 	var rs orginMssql.ReturnStatus
// 	rows, err := o.MssqlPreSales.GetDB().QueryContext(originCtx.Background(), ProcName,
// 		sql.Named("TxID", txID),
// 		sql.Named("TxStatus", txStatus),
// 		sql.Named("TxHash", txHash),
// 		sql.Named("BaseCoinID", BaseCoinID),
// 		sql.Named("GasFee", gasFee),
// 		&rs)
// 	if err != nil {
// 		log.Errorf(ProcName+" QueryContext err : %v", err)
// 		return err
// 	}

// 	defer rows.Close()

// 	if rs != 1 {
// 		log.Errorf(ProcName+" returnvalue error : %v", rs)
// 		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
// 	}

// 	return nil
// }

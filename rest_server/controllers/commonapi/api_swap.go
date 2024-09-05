package commonapi

import (
	"net/http"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/baseutil/otp_google"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi/inner"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/point_manager_server"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

// 전체 포인트, 코인 정보 리스트 조회
func GetSwapList(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	swapList := context.SwapList{
		PointList: model.GetDB().ScanPoints,
		AppPoints: model.GetDB().AppPoints,
		CoinList:  model.GetDB().Coins,
		SwapAble: context.SwapAble{
			SwapAbleP2C: model.GetDB().SwapAblePointToCoins,
			SwapAbleC2P: model.GetDB().SwapAbleCoinToPoints,
			SwapAbleC2C: model.GetDB().SwapAbleCoinToCoins,
			SwapAbleP2P: model.GetDB().SwapAblePointToPoints,
		},
	}

	resp.Value = swapList

	return c.JSON(http.StatusOK, resp)
}

// Swap 가능 정보 조회 (최소, 변동률, 수수료)
func GetSwapEnable(c echo.Context, reqSwapEnable *context.ReqSwapEnable) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return c.JSON(http.StatusOK, resp)
}

// Swap 처리
func PostSwap(ctx *context.InnoDashboardContext, reqSwapInfo *context.ReqSwapInfo) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// swap 사용 가능 상태 체크
	if !inner.IsEnableSwap(reqSwapInfo) {
		log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Support_Swap_Error])
		resp.SetReturn(resultcode.Result_Not_Support_Swap_Error)
		resp.Message = "This swap is not currently supported."
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	// otp check
	if config.GetInstance().Otp.EnableSwap {
		if !otp_google.VerifyTimebase(ctx.GetValue().InnoUID, reqSwapInfo.OtpCode) {
			resp.SetReturn(resultcode.Result_Get_Me_Verify_otp_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	Lockkey := model.MakeMemberSwapLockKey(ctx.GetValue().AUID)
	mutex := model.GetDB().RedSync.NewMutex(Lockkey)
	if err := mutex.Lock(); err != nil {
		log.Error(err)
		return err
	}

	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			if err != nil {
				log.Errorf("unlock err : %v", err)
			}
		}
	}()

	// 기초 정보 생성
	swapInfo := inner.MakeSwapInfo(ctx.GetValue(), reqSwapInfo)

	// SwapPoint 정보 추가 : P2C, C2P 인경우에만 처리
	if reqSwapInfo.EventID == context.EventID_toC2P ||
		reqSwapInfo.EventID == context.EventID_toP2C ||
		reqSwapInfo.EventID == context.EventID_toP2P {
		if _, membersMap, err := model.GetDB().USPAU_GetList_Members(swapInfo.AUID); err != nil {
			log.Errorf(resultcode.ResultCodeText[resultcode.Result_Get_MemberList_Scan_Error])
			resp.SetReturn(resultcode.Result_Get_MemberList_Scan_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		} else {
			if reqSwapInfo.EventID == context.EventID_toC2P {
				if member, ok := membersMap[swapInfo.SwapToPoint.AppID]; ok {
					swapInfo.SwapToPoint.MUID = member.MUID
					swapInfo.SwapToPoint.DatabaseID = member.DatabaseID
				} else {
					// swap 하려는 app point 정보가 없다.
					log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
					resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			} else if reqSwapInfo.EventID == context.EventID_toP2C {
				if member, ok := membersMap[swapInfo.SwapFromPoint.AppID]; ok {
					swapInfo.SwapFromPoint.MUID = member.MUID
					swapInfo.SwapFromPoint.DatabaseID = member.DatabaseID
				} else {
					// swap 하려는 app point 정보가 없다.
					log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
					resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			} else if reqSwapInfo.EventID == context.EventID_toP2P {
				if member, ok := membersMap[swapInfo.SwapFromPoint.AppID]; ok {
					swapInfo.SwapFromPoint.MUID = member.MUID
					swapInfo.SwapFromPoint.DatabaseID = member.DatabaseID
				} else {
					// swap 하려는 app point 정보가 없다.
					log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
					resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}

				if member, ok := membersMap[swapInfo.SwapToPoint.AppID]; ok {
					swapInfo.SwapToPoint.MUID = member.MUID
					swapInfo.SwapToPoint.DatabaseID = member.DatabaseID
				} else {
					// swap 하려는 app point 정보가 없다.
					log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
					resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			}
		}
	}

	if reqSwapInfo.EventID != context.EventID_toP2P {
		// SwapCoin 정보 추가
		// 내가 보유하고 있는 지갑이 있는지 검증한다.
		if wallets, err := model.GetDB().USPAU_GetList_AccountWallets(ctx.GetValue().AUID); err != nil {
			log.Errorf("USPAU_GetList_AccountWallets err : %v, auid:%v", err, ctx.GetValue().AUID)
			resp.SetReturn(resultcode.Result_Error_Db_GetAccountWallets)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		} else if len(wallets) == 0 {
			log.Errorf("USPAU_GetList_AccountWallets not exist wallet by auid : %v", ctx.GetValue().AUID)
			resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		} else {
			// 수수료로 지불할 지갑이 존재 하는지 찾기
			bFind := false
			for _, wallet := range wallets {
				// to coin 처리를 위해 정보 수집
				if swapInfo.TxType == context.EventID_toP2C || swapInfo.TxType == context.EventID_toC2C {
					if swapInfo.SwapToCoin.CoinID != 0 && wallet.BaseCoinID == model.GetDB().CoinsMap[swapInfo.SwapToCoin.CoinID].BaseCoinID && wallet.ConnectionStatus == 1 {
						swapInfo.SwapToCoin.WalletAddress = wallet.WalletAddress
						swapInfo.SwapToCoin.WalletTypeID = wallet.WalletTypeID
						swapInfo.SwapToCoin.WalletID = wallet.WalletID
						swapInfo.SwapToCoin.CoinSymbol = model.GetDB().CoinsMap[swapInfo.SwapToCoin.CoinID].CoinSymbol
						swapInfo.SwapToCoin.BaseCoinID = wallet.BaseCoinID
						swapInfo.SwapToCoin.BaseCoinSymbol = model.GetDB().BaseCoinMapByCoinID[wallet.BaseCoinID].BaseCoinSymbol

						bFind = true
					}
				}
				// from coin 처리를 위해 정보 수집
				if swapInfo.TxType == context.EventID_toC2P || swapInfo.TxType == context.EventID_toC2C {
					if swapInfo.SwapFromCoin.CoinID != 0 && wallet.BaseCoinID == model.GetDB().CoinsMap[swapInfo.SwapFromCoin.CoinID].BaseCoinID && wallet.ConnectionStatus == 1 {
						swapInfo.SwapFromCoin.WalletAddress = wallet.WalletAddress
						swapInfo.SwapFromCoin.WalletTypeID = wallet.WalletTypeID
						swapInfo.SwapFromCoin.WalletID = wallet.WalletID
						swapInfo.SwapFromCoin.CoinSymbol = model.GetDB().CoinsMap[swapInfo.SwapFromCoin.CoinID].CoinSymbol
						swapInfo.SwapFromCoin.BaseCoinID = wallet.BaseCoinID
						swapInfo.SwapFromCoin.BaseCoinSymbol = model.GetDB().BaseCoinMapByCoinID[wallet.BaseCoinID].BaseCoinSymbol

						bFind = true
					}
				}
			}
			if !bFind { // 수수료 지불한 지갑이 존재하지 않다면 에러
				log.Errorf("Not Find swap fee wallet / auid : %v", ctx.GetValue().AUID)
				resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			}

			// p2c, c2c 에는 수수료로 지불할 지갑이 존재하는 지 체크
			if swapInfo.TxType == context.EventID_toP2C {
				// 지갑에 수수료 확인을 위해서 보유량 가져오기
				inner.CheckSwapFee(swapInfo, resp)
				if resp.Return != 0 {
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			} else if swapInfo.TxType == context.EventID_toC2P { // coin -> point 에는 유저 지갑에 전환할 balance 존재하는지 만 체크
				// 스왑에 사용될 수량 존재 확인
				inner.CheckSwapCoinBalance(swapInfo, resp)
				if resp.Return != 0 {
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			} else if swapInfo.TxType == context.EventID_toC2C { // c2c는 보유 balance, fee 모두 체크 확인
				// 스왑에 사용될 수량 존재 확인
				inner.CheckSwapCoinBalance(swapInfo, resp)
				if resp.Return != 0 {
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
				// 스왑에 필요한 수수료 보유량 확인
				inner.CheckSwapFee(swapInfo, resp)
				if resp.Return != 0 {
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			}
		}
	}

	// 아래 체크 사항은 point manager server에서 처리한다.
	// 최소 변환 비율에 맞는지 체크
	// 전환 비율 계산 후 타당성 확인
	if resSwap, err := point_manager_server.GetInstance().PostPointCoinSwap(swapInfo); err != nil {
		log.Errorf("PostPointCoinSwap error : %v", err)
		resp.SetReturn(resultcode.Result_Unknown_Swap_Error)
	} else {
		if resSwap.Common.Return != 0 {
			log.Errorf("PostPointCoinSwap response return:%v, msg:%v", resSwap.Return, resSwap.Message)
			resp.Return = resSwap.Return
			resp.Message = resSwap.Message
		} else {
			resp.Value = resSwap.ReqSwapInfo
		}
	}
	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PutSwapStatus(ctx *context.InnoDashboardContext, params *context.PutSwapStatus) error {
	resp := new(base.BaseResponse)
	resp.Success()

	req := &point_manager_server.ReqSwapStatus{
		TxID:              params.TxID,
		TxStatus:          params.TxStatus,
		TxHash:            params.TxHash,
		FromWalletAddress: params.FromWalletAddress,
	}

	if resSwap, err := point_manager_server.GetInstance().PutSwapStatus(req); err != nil {
		log.Errorf("PutSwapStatus error : %v", err)
		resp.SetReturn(resultcode.Result_Unknown_Swap_Error)
	} else {
		if resSwap.Common.Return != 0 {
			resp.Return = resSwap.Return
			resp.Message = resSwap.Message
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetSwapInprogressNotExist(ctx *context.InnoDashboardContext, params *context.ReqSwapInprogress) error {
	resp := new(base.BaseResponse)
	resp.Success()

	req := &point_manager_server.ReqSwapInprogress{
		AUID: params.AUID,
	}

	if resSwap, err := point_manager_server.GetInstance().GetSwapInprogressNotExist(req); err != nil {
		log.Errorf("GetSwapInprogressNotExist error : %v", err)
		resp.SetReturn(resultcode.Result_Unknown_Swap_Error)
	} else {
		if resSwap.Common.Return != 0 {
			resp.Return = resSwap.Return
			resp.Message = resSwap.Message
			//resp.Value = resSwap.Value
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

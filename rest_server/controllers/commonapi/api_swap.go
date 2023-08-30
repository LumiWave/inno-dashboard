package commonapi

import (
	"math"
	"net/http"
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/baseutil/otp_google"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/util"
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
		Swapable:  model.GetDB().SwapAble,
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

	isToCoinEnable, isToPointEnable := model.GetSwapEnable()

	if reqSwapInfo.EventID == context.EventID_toPoint && !isToPointEnable {
		log.Errorf(resultcode.ResultCodeText[resultcode.Result_Invalid_CoinID_Error])
		resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
		resp.Message = "This swap is not currently supported."
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	if reqSwapInfo.EventID == context.EventID_toCoin && !isToCoinEnable {
		log.Errorf(resultcode.ResultCodeText[resultcode.Result_Invalid_CoinID_Error])
		resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
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

	swapInfo := &point_manager_server.ReqSwapInfo{
		AUID: ctx.GetValue().AUID,
		SwapPoint: point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.AdjustPointQuantity,
			PointQuantity:         0,
		},
		SwapCoin: point_manager_server.SwapCoin{
			CoinID:             reqSwapInfo.CoinID,
			WalletAddress:      "",
			AdjustCoinQuantity: reqSwapInfo.AdjustCoinQuantity,
		},
		TxType: reqSwapInfo.EventID,
	}

	// SwapPoint 정보 추가
	if _, membersMap, err := model.GetDB().USPAU_GetList_Members(swapInfo.AUID); err != nil {
		log.Errorf(resultcode.ResultCodeText[resultcode.Result_Get_MemberList_Scan_Error])
		resp.SetReturn(resultcode.Result_Get_MemberList_Scan_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		if member, ok := membersMap[swapInfo.AppID]; ok {
			swapInfo.MUID = member.MUID
			swapInfo.DatabaseID = member.DatabaseID
		} else {
			// swap 하려는 app point 정보가 없다.
			log.Errorf(resultcode.ResultCodeText[resultcode.Result_Not_Exist_AppPointInfo_Error])
			resp.SetReturn(resultcode.Result_Not_Exist_AppPointInfo_Error)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

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
			if wallet.BaseCoinID == model.GetDB().CoinsMap[swapInfo.CoinID].BaseCoinID {
				swapInfo.WalletAddress = wallet.WalletAddress
				swapInfo.CoinSymbol = model.GetDB().CoinsMap[swapInfo.CoinID].CoinSymbol
				swapInfo.BaseCoinID = wallet.BaseCoinID
				swapInfo.BaseCoinSymbol = model.GetDB().BaseCoinMapByCoinID[wallet.BaseCoinID].BaseCoinSymbol
				bFind = true
				break
			}
		}
		if !bFind { // 수수료 지불한 지갑이 존재하지 않다면 에러
			log.Errorf("Not Find swap fee wallet / auid : %v", ctx.GetValue().AUID)
			resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}

		// point -> coin 에는 수수료로 지불할 지갑이 존재하는 지 체크
		if swapInfo.TxType == context.EventID_toCoin {
			// 지갑에 수수료 확인을 위해서 보유량 가져오기
			req := &point_manager_server.ReqBalance{
				Symbol:  swapInfo.BaseCoinSymbol,
				Address: swapInfo.WalletAddress,
			}
			res, err := point_manager_server.GetInstance().GetBalance(req)
			if err != nil {
				log.Errorf("GetBalance err : %v, wallet:%v", err, swapInfo.WalletAddress)
				resp.SetReturn(resultcode.ResultInternalServerError)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			} else if res.Return != 0 {
				log.Errorf("GetBalance return : %v, msg:%v", res.Return, res.Message)
				resp.SetReturn(resultcode.ResultInternalServerError)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			} else {
				log.Debugf("GetBalance symbol:%v, addr:%v, bal:%v", req.Symbol, res.Value.Address, res.Value.Balance)
				// 스왑에 필요한 가스비 가지고 있는지 체크
				coinInfo := model.GetDB().CoinsMap[swapInfo.CoinID]
				redisKey := model.MakeCoinFeeKey(coinInfo.CoinSymbol)
				if coinFee, err := model.GetDB().GetCacheCoinFee(redisKey); err != nil {
					log.Errorf("GetCacheCoinFee err : %v", err)
					resp.SetReturn(resultcode.Result_CoinFee_NotExist)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				} else {
					balance := util.ToDecimalEncf(res.Value.Balance, res.Value.Decimal)

					basecoinRedisKey := model.MakeCoinFeeKey(swapInfo.BaseCoinSymbol)
					basecoinFee, err := model.GetDB().GetCacheCoinFee(basecoinRedisKey)
					if err != nil {
						log.Errorf("GetCacheCoinFee err : %v", err)
						resp.SetReturn(resultcode.Result_CoinFee_NotExist)
						return ctx.EchoContext.JSON(http.StatusOK, resp)
					}
					//// 수수료로 사용될 코인 balance를 가져와서 보유량 확인
					if balance <= coinFee.TransactionFee+basecoinFee.TransactionFee { // 부모지갑에 보낼 전송 수수료 + 부모가 보내줄 수수료만큼 있어야함
						resp.SetReturn(resultcode.Result_CoinFee_LackOfGas)
						return ctx.EchoContext.JSON(http.StatusOK, resp)
					}
					swapInfo.SwapFeeCoinID = coinFee.BaseCoinID
					swapInfo.SwapFeeCoinSymbol = coinFee.BaseCoinSymbol
					swapInfo.SwapFee = coinFee.TransactionFee

					swapInfo.SwapFeeT = util.ToDecimalDecStr(coinFee.TransactionFee, model.GetDB().CoinsMap[swapInfo.SwapFeeCoinID].Decimal)
					swapInfo.SwapFeeD = strconv.FormatFloat(coinFee.TransactionFee, 'f', -1, 64)

				}
			}
		} else if swapInfo.TxType == context.EventID_toPoint { // coin -> point 에는 유저 지갑에 전환할 balance가 존재하는지 체크
			// if balance <= coinFee.TransactionFee { // 부모지갑에 보낼 전송 수수료만 있으면 됨
			// 	resp.SetReturn(resultcode.Result_CoinFee_LackOfGas)
			// 	return ctx.EchoContext.JSON(http.StatusOK, resp)
			// }
			req := &point_manager_server.ReqBalance{
				Symbol:  swapInfo.CoinSymbol,
				Address: swapInfo.WalletAddress,
			}
			res, err := point_manager_server.GetInstance().GetBalance(req)
			if err != nil {
				log.Errorf("GetBalance err : %v, wallet:%v", err, swapInfo.WalletAddress)
				resp.SetReturn(resultcode.ResultInternalServerError)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			} else if res.Return != 0 {
				log.Errorf("GetBalance return : %v, msg:%v", res.Return, res.Message)
				resp.SetReturn(resultcode.ResultInternalServerError)
				return ctx.EchoContext.JSON(http.StatusOK, resp)
			} else {
				balance := util.ToDecimalEncf(res.Value.Balance, res.Value.Decimal)

				// 보유 수량 부족
				if math.Abs(swapInfo.SwapCoin.AdjustCoinQuantity) > balance {
					log.Errorf("not enough swap coin cur : %v, wallet:%v, symbol:%v", balance, swapInfo.WalletAddress, swapInfo.CoinSymbol)
					resp.SetReturn(resultcode.Result_CoinTransfer_NotEnough_Coin)
					return ctx.EchoContext.JSON(http.StatusOK, resp)
				}
			}

			swapInfo.SwapFee = 0
		}

	}
	swapInfo.InnoUID = ctx.GetValue().InnoUID

	// 아래 체크 사항은 point manager server에서 처리한다.
	// 최소 변환 비율에 맞는지 체크
	// 전환 비율 계산 후 타당성 확인
	if resSwap, err := point_manager_server.GetInstance().PostPointCoinSwap(swapInfo); err != nil {
		log.Errorf("PostPointCoinSwap error : %v", err)
		resp.SetReturn(resultcode.Result_Unknown_Swap_Error)
	} else {
		if resSwap.Common.Return != 0 {
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

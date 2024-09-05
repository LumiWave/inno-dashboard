package inner

import (
	"math"
	"strconv"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/auth"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/point_manager_server"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/LumiWave/inno-dashboard/rest_server/util"
)

func IsEnableSwap(reqSwapInfo *context.ReqSwapInfo) bool {
	isToCoinEnable, isToPointEnable, isToC2CEnable := model.GetSwapEnable()

	if reqSwapInfo.EventID == context.EventID_toC2P && !isToPointEnable {
		return false
	}

	if reqSwapInfo.EventID == context.EventID_toP2C && !isToCoinEnable {
		return false
	}

	if reqSwapInfo.EventID == context.EventID_toC2C && !isToC2CEnable {
		return false
	}

	return true
}

func MakeSwapInfo(token *auth.VerifyAuthToken, reqSwapInfo *context.ReqSwapInfo) *point_manager_server.ReqSwapInfo {
	swapInfo := &point_manager_server.ReqSwapInfo{
		AUID:    token.AUID,
		InnoUID: token.InnoUID,
		TxType:  reqSwapInfo.EventID,
	}

	switch reqSwapInfo.EventID {
	case context.EventID_toC2P:
		swapInfo.SwapFromCoin = point_manager_server.SwapCoin{
			CoinID:             reqSwapInfo.SwapFromCoin.CoinID,
			WalletAddress:      "",
			AdjustCoinQuantity: reqSwapInfo.SwapFromCoin.AdjustCoinQuantity,
		}
		swapInfo.SwapToPoint = point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.SwapToPoint.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.SwapToPoint.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.SwapToPoint.AdjustPointQuantity,
			PointQuantity:         0,
		}
	case context.EventID_toP2C:
		swapInfo.SwapFromPoint = point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.SwapFromPoint.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.SwapFromPoint.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.SwapFromPoint.AdjustPointQuantity,
			PointQuantity:         0,
		}
		swapInfo.SwapToCoin = point_manager_server.SwapCoin{
			CoinID:             reqSwapInfo.SwapToCoin.CoinID,
			WalletAddress:      "",
			AdjustCoinQuantity: reqSwapInfo.SwapToCoin.AdjustCoinQuantity,
		}
	case context.EventID_toC2C:
		swapInfo.SwapFromCoin = point_manager_server.SwapCoin{
			CoinID:             reqSwapInfo.SwapFromCoin.CoinID,
			WalletAddress:      "",
			AdjustCoinQuantity: reqSwapInfo.SwapFromCoin.AdjustCoinQuantity,
		}
		swapInfo.SwapToCoin = point_manager_server.SwapCoin{
			CoinID:             reqSwapInfo.SwapToCoin.CoinID,
			WalletAddress:      "",
			AdjustCoinQuantity: reqSwapInfo.SwapToCoin.AdjustCoinQuantity,
		}
	case context.EventID_toP2P:
		swapInfo.SwapFromPoint = point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.SwapFromPoint.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.SwapFromPoint.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.SwapFromPoint.AdjustPointQuantity,
			PointQuantity:         0,
		}
		swapInfo.SwapToPoint = point_manager_server.SwapPoint{
			MUID:                  0,
			AppID:                 reqSwapInfo.SwapToPoint.AppID,
			DatabaseID:            0,
			PointID:               reqSwapInfo.SwapToPoint.PointID,
			PreviousPointQuantity: 0,
			AdjustPointQuantity:   reqSwapInfo.SwapToPoint.AdjustPointQuantity,
			PointQuantity:         0,
		}
	}

	return swapInfo
}

// swap에 필요한 수수료 보유 수량 체크
func CheckSwapFee(swapInfo *point_manager_server.ReqSwapInfo, resp *base.BaseResponse) {

	swapCoin := &point_manager_server.SwapCoin{}
	if swapInfo.TxType == context.EventID_toP2C {
		swapCoin = &swapInfo.SwapToCoin
	} else if swapInfo.TxType == context.EventID_toC2P {
		return
	} else if swapInfo.TxType == context.EventID_toC2C {
		swapCoin = &swapInfo.SwapToCoin
	}

	// 지갑에 수수료 확인을 위해서 보유량 가져오기
	req := &point_manager_server.ReqBalance{
		Symbol:  swapCoin.BaseCoinSymbol,
		Address: swapCoin.WalletAddress,
	}
	res, err := point_manager_server.GetInstance().GetBalance(req)
	if err != nil {
		log.Errorf("GetBalance err : %v, wallet:%v", err, swapCoin.WalletAddress)
		resp.SetReturn(resultcode.ResultInternalServerError)
		return
	} else if res.Return != 0 {
		log.Errorf("GetBalance return : %v, msg:%v", res.Return, res.Message)
		resp.SetReturn(resultcode.ResultInternalServerError)
		return
	} else {
		log.Debugf("GetBalance symbol:%v, addr:%v, bal:%v", req.Symbol, res.Value.Address, res.Value.Balance)
		// 스왑에 필요한 가스비 가지고 있는지 체크
		coinInfo := model.GetDB().CoinsMap[swapCoin.CoinID]
		redisKey := model.MakeCoinFeeKey(coinInfo.CoinSymbol)
		if coinFee, err := model.GetDB().GetCacheCoinFee(redisKey); err != nil {
			log.Errorf("GetCacheCoinFee err : %v", err)
			resp.SetReturn(resultcode.Result_CoinFee_NotExist)
			return
		} else {
			balance := util.ToDecimalEncf(res.Value.Balance, res.Value.Decimal)

			basecoinRedisKey := model.MakeCoinFeeKey(swapCoin.BaseCoinSymbol)
			basecoinFee, err := model.GetDB().GetCacheCoinFee(basecoinRedisKey)
			if err != nil {
				log.Errorf("GetCacheCoinFee err : %v", err)
				resp.SetReturn(resultcode.Result_CoinFee_NotExist)
				return
			}
			//// 수수료로 사용될 코인 balance를 가져와서 보유량 확인
			if balance <= coinFee.TransactionFee+basecoinFee.TransactionFee { // 부모지갑에 보낼 전송 수수료 + 부모가 보내줄 수수료만큼 있어야함
				log.Errorf("lock of gas for swap audi : %v, coin_id:%v, symbol:%v, balance:%v", swapInfo.AUID, swapCoin.CoinID, swapCoin.CoinSymbol, balance)
				resp.SetReturn(resultcode.Result_CoinFee_LackOfGas)
				return
			}

			for _, coin := range model.GetDB().CoinsMap {
				if coin.CoinSymbol == coinFee.BaseCoinSymbol {
					swapInfo.SwapFeeCoinID = coin.CoinId
					swapInfo.SwapFeeCoinSymbol = coin.CoinSymbol
					break
				}
			}
			swapInfo.SwapFee = coinFee.TransactionFee

			swapInfo.SwapFeeT = util.ToDecimalDecStr(coinFee.TransactionFee, model.GetDB().CoinsMap[swapInfo.SwapFeeCoinID].Decimal)
			swapInfo.SwapFeeD = strconv.FormatFloat(coinFee.TransactionFee, 'f', -1, 64)

		}
	}
}

// swap에 필요한 코인 보유량 체크
func CheckSwapCoinBalance(swapInfo *point_manager_server.ReqSwapInfo, resp *base.BaseResponse) {

	swapCoin := &point_manager_server.SwapCoin{}
	if swapInfo.TxType == context.EventID_toP2C {
		return
	} else if swapInfo.TxType == context.EventID_toC2P || swapInfo.TxType == context.EventID_toC2C {
		swapCoin = &swapInfo.SwapFromCoin
	}

	req := &point_manager_server.ReqBalance{
		Symbol:  swapCoin.CoinSymbol,
		Address: swapCoin.WalletAddress,
	}
	res, err := point_manager_server.GetInstance().GetBalance(req)
	if err != nil {
		log.Errorf("GetBalance err : %v, wallet:%v", err, swapCoin.WalletAddress)
		resp.SetReturn(resultcode.ResultInternalServerError)
	} else if res.Return != 0 {
		log.Errorf("GetBalance return : %v, msg:%v", res.Return, res.Message)
		resp.SetReturn(resultcode.ResultInternalServerError)
	} else {
		balance := util.ToDecimalEncf(res.Value.Balance, res.Value.Decimal)

		// 보유 수량 부족
		if math.Abs(swapCoin.AdjustCoinQuantity) > balance {
			log.Errorf("not enough swap coin cur : %v, wallet:%v, symbol:%v", balance, swapCoin.WalletAddress, swapCoin.CoinSymbol)
			resp.SetReturn(resultcode.Result_CoinTransfer_NotEnough_Coin)
		}
	}
}

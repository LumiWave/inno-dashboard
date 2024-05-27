package commonapi

import (
	"net/http"
	"time"

	"github.com/LumiWave/baseInnoClient/market"
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/baseutil/otp_google"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/inno_market"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/point_manager_server"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/LumiWave/inno-dashboard/rest_server/util"

	"github.com/labstack/echo"
)

// 지갑 정보 조회
func GetMeWallets(c echo.Context, reqMeCoin *context.ReqMeCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if reqMeCoin.CoinID == 0 {
		balances, err := point_manager_server.GetInstance().GetBalanceAll(&point_manager_server.ReqBalanceAll{AUID: reqMeCoin.AUID})
		if err != nil {
			log.Errorf("GetBalanceAll Error : %v", err)
			resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
		} else {
			if balances.Return == 0 {
				if len(balances.Value.Balances) == 0 {
					resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
				} else {
					resp.Value = balances.Value.Balances
				}
			} else {
				resp.Message = balances.Message
				resp.Return = balances.Return
			}
		}
	} else {
		userWallets, err := model.GetDB().USPAU_GetList_AccountWallets(reqMeCoin.AUID)
		if err != nil {
			resp.SetReturn(resultcode.Result_Get_Me_Wallets_Regiest_Error)
		} else {
			meCoin, ok := model.GetDB().CoinsMap[reqMeCoin.CoinID]
			if !ok {
				resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
			} else {
				var targetWallet *context.DBWalletRegist
				for _, userwallet := range userWallets {
					if userwallet.ConnectionStatus == 1 && userwallet.BaseCoinID == meCoin.BaseCoinID {
						targetWallet = userwallet
					}
				}
				if targetWallet == nil {
					resp.SetReturn(resultcode.Result_Get_Me_Wallets_Regiest_Error)
				} else {
					balance, err := point_manager_server.GetInstance().GetBalance(&point_manager_server.ReqBalance{Symbol: meCoin.CoinSymbol, Address: targetWallet.WalletAddress})
					if err != nil {
						log.Errorf("GetBalanceAll Error : %v", err)
						resp.SetReturn(resultcode.Result_Get_Me_WalletList_Scan_Error)
					} else {
						if balance.Return == 0 {
							if len(balance.Value.Balance) == 0 {
								resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
							} else {
								res := []*point_manager_server.Balance{
									{
										CoinID:     reqMeCoin.CoinID,
										BaseCoinID: meCoin.BaseCoinID,
										Symbol:     meCoin.CoinSymbol,
										Balance:    balance.Value.Balance,
										Address:    balance.Value.Address,
										Decimal:    balance.Value.Decimal,
									},
								}
								resp.Value = res
							}
						} else {
							resp.Message = balance.Message
							resp.Return = balance.Return
						}
					}
				}
			}
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 포인트 리스트 조회
func GetMePointList(c echo.Context, reqMePoint *context.ReqMePoint) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if pointList, err := model.GetDB().USPAU_GetList_AccountPoints(reqMePoint.AUID, 0); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_PointList_Scan_Error)
	} else {
		if _, membersMap, err := model.GetDB().USPAU_GetList_Members(reqMePoint.AUID); err != nil {
			resp.SetReturn(resultcode.Result_Get_MemberList_Scan_Error)
		} else {
			for _, member := range membersMap {
				// 포인트 서버에서 현재 실제 정보 가져와서 merge
				if memberInfo, err := point_manager_server.GetInstance().GetPointAppList(member.MUID, member.DatabaseID); err == nil {
					for _, point := range memberInfo.Points {
						for _, mePoint := range pointList {
							if point.PointID == mePoint.PointID {
								mePoint.Quantity = point.Quantity
							}
						}
					}
				} else {
					log.Errorf("point_manager_server GetPointAppList error : %v", err)
				}
			}

			resp.Value = []*context.MePoint{}
			if pointList != nil {
				resp.Value = pointList
			}
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// App 별 총/금일 누적 코인 리스트 조회
func GetMeCoinList(c echo.Context, reqMeCoin *context.ReqMeCoin) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if coinList, err := model.GetDB().USPAU_GetList_AccountCoins(reqMeCoin.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Get_Me_CoinList_Scan_Error)
	} else {
		resp.Value = []*context.MeCoin{}
		if coinList != nil {
			resp.Value = coinList
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func GetCoinObjects(req *context.ReqCoinObjects, ctx *context.InnoDashboardContext) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if _, ok := model.GetDB().CoinsMap[req.CoinID]; !ok {
		resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	}

	if wallets, err := model.GetDB().USPAU_GetList_AccountWallets(ctx.GetValue().AUID); err != nil {
		log.Errorf("USPAU_GetList_AccountWallets err : %v, auid:%v", err, ctx.GetValue().AUID)
		resp.SetReturn(resultcode.Result_Error_Db_GetAccountWallets)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else if len(wallets) == 0 {
		log.Errorf("USPAU_GetList_AccountWallets not exist wallet by auid : %v", ctx.GetValue().AUID)
		resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
		return ctx.EchoContext.JSON(http.StatusOK, resp)
	} else {
		bFind := false
		params := &point_manager_server.ReqCoinObjects{}
		for _, wallet := range wallets {
			if wallet.BaseCoinID == model.GetDB().CoinsMap[req.CoinID].BaseCoinID {
				params.WalletAddress = wallet.WalletAddress
				params.ContractAddress = model.GetDB().CoinsMap[req.CoinID].ContractAddress
				bFind = true

				if res, err := point_manager_server.GetInstance().GetCoinObjectIDS(params); err != nil {
					log.Errorf("GetCoinObjectIDS err : %v,  wallet:%v, contract:%v ", err, params.WalletAddress, params.ContractAddress)
					resp.SetReturn(resultcode.ResultInternalServerError)
				} else {
					if res.Common.Return == 0 {
						resValue := new(context.ResCoinObjects)
						resValue.ObjectIDs = res.Value.ObjectIDs
						resp.Value = resValue
					} else {
						log.Errorf("GetCoinObjectIDS error return : %v, %s, wallet:%v, contract:%v", res.Common.Return, res.Common.Message, params.WalletAddress, params.ContractAddress)
						resp.SetReturn(resultcode.ResultInternalServerError)
					}
				}

				break
			}
		}
		if !bFind { // 수수료 지불한 지갑이 존재하지 않다면 에러
			log.Errorf("Not Find swap fee wallet / auid : %v", ctx.GetValue().AUID)
			resp.SetReturn(resultcode.Result_Error_Db_NotExistWallets)
			return ctx.EchoContext.JSON(http.StatusOK, resp)
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

// google otp : qrcode용 uri 조회
func GetOtpUri(ctx *context.InnoDashboardContext) error {
	resp := new(base.BaseResponse)
	resp.Success()

	conf := config.GetInstance().Otp

	resp.Value = context.MeOtpUri{
		OtpUri: otp_google.MakeTimebaseUri(ctx.GetValue().InnoUID, ctx.GetValue().InnoUID, conf.IssueName),
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetOtpVerify(ctx *context.InnoDashboardContext, params *context.MeOtpVerify) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if !otp_google.VerifyTimebase(ctx.GetValue().InnoUID, params.OtpCode) {
		resp.SetReturn(resultcode.Result_Get_Me_Verify_otp_Error)
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostCoinReload(ctx *context.InnoDashboardContext, params *context.CoinReload) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// req := &point_manager_server.CoinReload{
	// 	AUID: params.AUID,
	// }

	// if res, err := point_manager_server.GetInstance().PostCoinReload(req); err == nil {
	// 	resp.Value = res.Value
	// } else {
	// 	log.Errorf("point_manager_server GetPointAppList error : %v", err)
	// 	resp.SetReturn(resultcode.ResultInternalServerError)
	// }

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqGetWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletRegistMap, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		res := &context.ResGetWalletRegist{
			Wallets: make([]*context.WalletRegistInfo, 0),
		}
		for _, walletRegist := range walletRegistMap {
			res.Wallets = append(res.Wallets, walletRegist)
		}
		resp.Value = res
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqPostWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletRegistMap, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		if walletData, ok := walletRegistMap[params.BaseCoinID]; !ok {
			resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_UnsupportWallet_Error)
		} else {
			if walletData.IsRegistered {
				resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegistered_Error)
			} else {
				if allowWallets, ok := model.GetDB().AllowWalletTypeMap[params.BaseCoinID]; !ok {
					resp.SetReturn(resultcode.Result_Invalid_CoinID_Error)
				} else {
					if !util.ContainsInt(params.WalletTypeID, allowWallets) {
						resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_NotAllowedWalletType)
					} else {
						isMigration := true
						if errType, isMigrated, err := model.GetDB().USPAU_Cnct_AccountWallets(params.AUID, params.BaseCoinID, params.WalletAddress, params.WalletTypeID); err != nil {
							switch errType {
							case 2:
								resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegisteredDB_Error)
							case 3:
								resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_AreadyRegistered_AnotherAccount_Error)
							default:
								resp.SetReturn(resultcode.Result_DBError)
							}
						} else {
							isMigration = isMigrated
						}

						// 등록하기 전에 마이그래션을 위한 정보를 수집해준다.
						// 마이그레이션은 비동기로 처리해준다.
						if !isMigration {
							go func(auid int64, userWalletAddress string) {
								if migCoins, migNFT, err := model.GetDB().USPAU_GetList_MigrationData(auid); err != nil {
									log.Errorf("USPAU_GetList_MigrationData err : %v, auid:%v", err, auid)
								} else {
									// 시작 전에 redis에 남기도록 해서 문제가 발생시 레디스를 삭제하지 않고 그대로 두고 추후 수동으로 처리해준다.
									// 코인 전송 시작
									for _, coin := range migCoins {
										coin.WalletAddress = userWalletAddress
										coin.Ts = time.Now().Unix()

										if coin.Quantity != 0 {
											// 레디스에 전송 시작 기록
											if err := model.GetDB().SetCacheIMGCoinTransfer(auid, coin); err != nil {
												log.Errorf("[IMG] SetCacheIMGCoinTransfer err:%v, auid:%v", err, auid)
											}

											req := &point_manager_server.ReqCoinTransferFromParentWallet{
												AUID:             auid,
												CoinID:           coin.CoinID,
												CoinSymbol:       model.GetDB().CoinsMap[coin.CoinID].CoinSymbol,
												ToAddress:        userWalletAddress,
												Quantity:         coin.Quantity,
												IsNormalTransfer: true,
											}
											if res, err := point_manager_server.GetInstance().PostCoinTransferFromParentWallet(req); err != nil {
												log.Errorf("[IMG] PostCoinTransferFromParentWallet err : %v, auid:%v, coinid:%v, toAddress:%v, quantity:%v", err, auid, coin.CoinID, userWalletAddress, coin.Quantity)
											} else {
												if res.Common.Return != 0 {
													log.Errorf("[IMG] PostCoinTransferFromParentWallet return err : %v, auid:%v, coinid:%v, toAddress:%v, quantity:%v", res.Common.Return, auid, coin.CoinID, userWalletAddress, coin.Quantity)
												} else {
													log.Infof("[IMG] coin send success tx:%v, auid:%v, coinid:%v, toAddress:%v, quantity:%v", res.Value.TransactionId, auid, coin.CoinID, userWalletAddress, coin.Quantity)
													// 전송 완료라고 생각하고 레디스 삭제 만약에 삭제가 안되어 있으면 문제가 있어서 전송 안된것으로 간주한다.
													if err := model.GetDB().DelCacheIMGCoinTransfer(auid, coin); err != nil {
														log.Errorf("[IMG] DelCacheIMGCoinTransfer:%v, auid:%v", err, auid)
													}
												}
											}
										}
									}
									// nft 전송 시작
									for _, nft := range migNFT {
										if err := model.GetDB().SetCacheIMGNFTransfer(params.AUID, nft); err != nil {
											log.Errorf("[IMG] SetCacheIMGNFTransfer err:%v, auid:%v, nftid:%v ", err, auid, nft.NFTID)
										}
										req := &market.ReqPostNFTTransferFromParent{
											AUID:            auid,
											NFTPackID:       nft.NFTPackID,
											CoinID:          nft.CoinID,
											NFTID:           nft.NFTID,
											ToWalletAddress: userWalletAddress,
										}
										if res, err := inno_market.GetInstance().PostNFTTransferFromParent(req); err != nil {
											log.Errorf("[IMG] PostNFTTransferFromParent err : %v, auid:%v, nftid:%v, toAddress:%v", err, auid, nft.NFTID, userWalletAddress)
										} else {
											if res.Common.Return != 0 {
												log.Errorf("[IMG] PostNFTTransferFromParent return err : %v, auid:%v, nftid:%v, toAddress:%v", res.Common.Return, auid, nft.NFTID, userWalletAddress)
											} else {
												log.Infof("[IMG] nft send success tx:%v, auid:%v, nftid:%v, toAddress:%v", res.Value.TxHash, auid, nft.NFTID, userWalletAddress)
												// 전송 완료라고 생각하고 레디스 삭제 만약에 삭제가 안되어 있으면 문제가 있어서 전송 안된것으로 간주한다.
												if err := model.GetDB().DelCacheIMGNFTTransfer(auid, nft); err != nil {
													log.Errorf("[IMG] DelCacheIMGNFTTransfer:%v, auid:%v", err, auid)
												}
											}
										}
									}
									// 마이그레이션 완료 처리
								}
							}(params.AUID, params.WalletAddress)
						}
					}
				}
			}
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func DeleteWalletRegist(ctx *context.InnoDashboardContext, params *context.ReqDeleteWalletRegist) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if walletRegistMap, errCode := GetWalletRegistInfo(params.AUID); errCode > 0 {
		resp.SetReturn(errCode)
	} else {
		if walletData, ok := walletRegistMap[params.BaseCoinID]; !ok {
			resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_UnsupportWallet_Error)
		} else {
			if walletData.IsRegistered {
				if registDT, err := time.Parse(time.RFC3339, walletData.RegistDT); err != nil {
					log.Errorf("wallet registDT parse error : %v", err)
					resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_System_Error)
				} else {
					limitDT := registDT.Add(time.Hour * context.DeleteWalletHour)
					//limitDT := registDT.Add(0)
					cmp := limitDT.Compare(time.Now())
					if cmp > 0 {
						resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_DeleteTime_Error) //24시간이 안됨
					} else {
						if walletData.WalletAddress != params.WalletAddress || walletData.WalletTypeID != params.WalletTypeID {
							resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_Diffrent_Wallet_Error) //현재 등록되어있는 지갑주소or지갑종류가 다름
						} else {
							if err := model.GetDB().USPAU_Dscnct_AccountWallets(params.AUID, walletData.BaseCoinID, params.WalletAddress, params.WalletTypeID); err != nil {
								resp.SetReturn(resultcode.Result_DBError)
							}
						}
					}
				}
			} else {
				resp.SetReturn(resultcode.Result_Post_Me_WalletRegist_NoRegistered_Wallet_Error)
			}
		}
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func GetWalletRegistInfo(auid int64) (map[int64]*context.WalletRegistInfo, int) {
	if userWallets, err := model.GetDB().USPAU_GetList_AccountWallets(auid); err != nil {
		return nil, resultcode.Result_Get_Me_AUID_Empty
	} else {
		res := make(map[int64]*context.WalletRegistInfo)

		for _, baseCoin := range model.GetDB().BaseCoins.Coins {
			res[baseCoin.BaseCoinID] = &context.WalletRegistInfo{
				BaseCoinID:     baseCoin.BaseCoinID,
				BaseCoinSymbol: baseCoin.BaseCoinSymbol,
			}
		}
		for baseCoinID, walletRegistInfo := range res {
			for _, userWallet := range userWallets {
				if baseCoinID == userWallet.BaseCoinID {
					walletType := model.GetDB().WalletTypeMap[userWallet.WalletTypeID]

					if auid > context.UserTypeLimit {
						walletRegistInfo.UserType = 2
					} else {
						walletRegistInfo.UserType = 1
					}
					walletRegistInfo.WalletID = userWallet.WalletID
					switch userWallet.ConnectionStatus {
					case 1:
						walletRegistInfo.IsRegistered = true
						walletRegistInfo.WalletAddress = userWallet.WalletAddress
						walletRegistInfo.RegistDT = userWallet.ModifiedDT
						walletRegistInfo.WalletTypeID = walletType.WalletTypeID
						walletRegistInfo.WalletName = walletType.WalletName
					case 2:
						walletRegistInfo.LastDeleteWalletAddress = userWallet.WalletAddress
						walletRegistInfo.LastDeleteDT = userWallet.ModifiedDT
						walletRegistInfo.LastDeleteWalletTypeID = walletType.WalletTypeID
						walletRegistInfo.LastDeleteWalletName = walletType.WalletName
					default:
					}
				}
			}
		}

		//미등록상태이면 마지막 등록 주소는 내보내지않는다
		for _, userWallet := range res {
			if !userWallet.IsRegistered {
				userWallet.LastDeleteWalletAddress = ""
				userWallet.LastDeleteDT = ""
				userWallet.LastDeleteWalletTypeID = 0
				userWallet.LastDeleteWalletName = ""
			}
		}

		return res, 0
	}
}

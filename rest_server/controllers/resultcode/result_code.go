package resultcode

const (
	Result_Success            = 0
	ResultInternalServerError = 500

	Result_Require_PageInfo = 30001 // 유효한 페이지 정보 필요

	Result_Upbit_EmptyCoinSymbol = 32001 // 코인 심볼이 비어있음
	Result_Upbit_EmptyUnit       = 32002 // Unit 이 비어있음
	Result_Upbit_InvalidUnit     = 32003 // Unit 이 유효하지 않음
	Result_Upbit_EmptyCount      = 32004 // Count가 비어있음
	Result_Upbit_EmptyTo         = 32005 // To가  비어있음
	Result_Upbit_TickerMarkets   = 32006 // 업비트 시세 Ticker 조회 API 에러
	Result_Upbit_CandleMinutes   = 32007 // 업비트 CandleMinutes 조회 API 에러
	Result_Upbit_CandleDays      = 32008 // 업비트 CandleDays 조회 API 에러
	Result_Upbit_CandleWeeks     = 32009 // 업비트 CandleWeeks 조회 API 에러
	Result_Upbit_CandleMonths    = 32010 // 업비트 CandleMonths 조회 API 에러

	Result_Invalid_EventID_Error        = 32100 // eventid 유효성 에러
	Result_Invalid_AppID_Error          = 32101 // appid 유효성 에러
	Result_Invalid_PointID_Error        = 32102 // pointid 유효성 에러
	Result_Invalid_PointQuantity_Error  = 32103 // point 수량 유효성 에러
	Result_Invalid_CoinID_Error         = 32104 // coinid 유효성 에러
	Result_Invalid_CoinQuantity_Error   = 32105 // coin 수량 유효성 에러
	Result_Invalid_AdjustQuantity_Error = 32106 // adjust 수량 유효성 에러
	Result_Invalid_WalletPlatform_Error = 32107 // wallet platform 이름
	Result_Invalid_WalletAddress_Error  = 32108 // wallet address
	Result_Invalid_TxStatus             = 32109 // invalid tx status
	Result_Invalid_WalletTypeID_Error   = 32110 // wallet platform 이름
	Result_Invalid_BaseCoinID_Error     = 32111 // basecoinid 유효성 에러

	Result_Not_Exist_AppPointInfo_Error = 32201 // 앱 포인트 정보가 존재하지 않는다.
	Result_Unknown_Swap_Error           = 32202 // unknown swap 에러

	Result_Get_App_AppID_Empty                = 33001 // Get App AppID Empty
	Result_Get_App_PointID_Empty              = 33002 // Get App PointID Empty
	Result_Get_App_CoinID_Empty               = 33003 // Get App CoinID Empty
	Result_Get_App_Point_DailyLiquidity_Error = 33004 // Get App Point DailyLiquidity DB Error
	Result_Get_App_Coin_DailyLiquidity_Error  = 33005 // Get App Coin DailyLiquidity DB Error
	Result_Get_App_Point_Liquidity_Error      = 33006 // Get App Point Liquidity DB Error
	Result_Get_App_Coin_Liquidity_Error       = 33007 // Get App Coin Liquidity DB Error

	Result_Get_Swap_ExchangeGoods_Scan_Error = 33100 // USPAU_GetList_AccountPoints DB Scan Error

	Result_Get_Me_AUID_Empty            = 34001 // GetMeWallets AUID Empty
	Result_Get_Me_WalletList_Scan_Error = 34002 // GetMeWalletList DB Scan Error
	Result_Get_Me_PointList_Scan_Error  = 34003 // GetMePointList DB Scan Error
	Result_Get_Me_CoinList_Scan_Error   = 34004 // GetMeCoinList DB Scan Error
	Result_Get_MemberList_Scan_Error    = 34005 // USPAU_GetList_Members DB Scan Error
	Result_Get_Me_Verify_otp_Error      = 34006 // otp verify error
	Result_Error_Db_GetAccountWallets   = 34007 // db 지갑 리스트 조회 에러
	Result_Error_Db_NotExistWallets     = 34008 // db 보유한 지갑 없음 에러
	Result_Get_Me_Wallets_Regiest_Error = 34009 // 지갑등록안됨

	Result_Post_Me_WalletRegist_UnsupportWallet_Error                 = 34101 // 지원하지 않는 지갑 종류
	Result_Post_Me_WalletRegist_AreadyRegistered_Error                = 34102 // 이미 지갑이 등록된 계정
	Result_Post_Me_WalletRegist_AreadyRegisteredDB_Error              = 34103 // 이미 지갑이 등록된 계정. 서버는 통과했는데 디비만 정보가있따.
	Result_Post_Me_WalletRegist_AreadyRegistered_AnotherAccount_Error = 34104 // 이미 지갑이 다른계정에 등록되어있다
	Result_Post_Me_WalletRegist_NoRegistered_Wallet_Error             = 34105 // 지갑을 찾을 수 없습니다
	Result_Post_Me_WalletRegist_System_Error                          = 34106 // 공통오류
	Result_Post_Me_WalletRegist_DeleteTime_Error                      = 34107 // 해제시간이 안됨
	Result_Post_Me_WalletRegist_Diffrent_Wallet_Error                 = 34108 // 등록된 지갑주소와 정보가 다름
	Result_Post_Me_WalletRegist_NotAllowedWalletType                  = 34109 // 허용된 지갑종류가 아님

	Result_CoinTransfer_CoinSymbol_Empty = 36001 // Coin Symbol is Empty
	Result_CoinTransfer_ToAddress_Empty  = 36002 // To Address is Empty
	Result_CoinTransfer_Quantity_Empty   = 36003 // Amount is Empty
	Result_CoinTransfer_NotEnough_Coin   = 36004 // 전송할 코인량이 충분하지 않다.
	Result_CoinFee_BaseSymbol_Empty      = 36005 // base coin symbol 정보가 필요하다.
	Result_CoinFee_NotExist              = 36006 // GetCacheCoinFee Error
	Result_CoinFee_LackOfGas             = 36007 // loack of gas

	Result_PubSub_InternalErr                      = 36100 // pubsub error
	Result_Error_IsMaintenance                     = 36101 // 정검 중
	Result_Error_IsSwapMaintenance                 = 36102 // swap 정검중
	Result_Error_IsCoinTransferExternalMaintenance = 36103 // 외부 코인 전송 정검중

	Result_DBError              = 19000 // db 에러
	Result_Invalid_DBID         = 19001 // 유효하지 못한 database index
	Result_DBError_Unknown      = 19002 // 알려지지 않은 db 에러
	Result_RedisError_Lock_fail = 19003 // redis lock error

	Result_Auth_RequireMessage    = 20000
	Result_Auth_RequireSign       = 20001
	Result_Auth_InvalidLoginInfo  = 20002
	Result_Auth_DontEncryptJwt    = 20003
	Result_Auth_InvalidJwt        = 20004
	Result_Auth_InvalidWalletType = 20005
)

var ResultCodeText = map[int]string{
	Result_Success:            "success",
	ResultInternalServerError: "internal server error",

	Result_Require_PageInfo: "require page info",

	Result_Upbit_EmptyCoinSymbol: "Empty coin symbol",
	Result_Upbit_EmptyUnit:       "Empty Unit",
	Result_Upbit_InvalidUnit:     "Invalid Unit",
	Result_Upbit_EmptyCount:      "Empty Count",
	Result_Upbit_EmptyTo:         "Empty To",

	Result_Upbit_TickerMarkets: "Upbit Ticker Markets API Error",
	Result_Upbit_CandleMinutes: "Upbit Candle Minutes API Error",
	Result_Upbit_CandleDays:    "Upbit Candle Days API Error",
	Result_Upbit_CandleWeeks:   "Upbit Candle Weeks API Error",
	Result_Upbit_CandleMonths:  "Upbit Candle Months API Error",

	Result_Invalid_EventID_Error:        "TxType is invalid",
	Result_Invalid_AppID_Error:          "AppID is invalid",
	Result_Invalid_PointID_Error:        "PointID is invalid",
	Result_Invalid_PointQuantity_Error:  "point quantity is invalid",
	Result_Invalid_CoinID_Error:         "CoinID is invalid",
	Result_Invalid_CoinQuantity_Error:   "Coin quantity is invalid",
	Result_Invalid_AdjustQuantity_Error: "Adjust quantity is invalid",
	Result_Invalid_WalletPlatform_Error: "Wallet Platform is invalid",
	Result_Invalid_WalletAddress_Error:  "Wallet Address is invalid",
	Result_Invalid_TxStatus:             "TxStatus is invalid",
	Result_Invalid_WalletTypeID_Error:   "wallet_type_id is invalid",
	Result_Invalid_BaseCoinID_Error:     "base_coin_id is invalid",

	Result_Not_Exist_AppPointInfo_Error: "App point information does not exist",
	Result_Unknown_Swap_Error:           "Unknown swap error",

	Result_Get_App_AppID_Empty:                "AppID is empty",
	Result_Get_App_PointID_Empty:              "PointID is empty",
	Result_Get_App_CoinID_Empty:               "CoinID is empty",
	Result_Get_App_Point_DailyLiquidity_Error: "Get app point daily liquidity db error",
	Result_Get_App_Coin_DailyLiquidity_Error:  "Get app coin daily liquidity db error",
	Result_Get_App_Point_Liquidity_Error:      "Get app point liquidity db error",
	Result_Get_App_Coin_Liquidity_Error:       "Get app coin liquidity db error",

	Result_Get_Swap_ExchangeGoods_Scan_Error: "USPAU_GetList_AccountPoints DB Scan Error",

	Result_Get_Me_AUID_Empty:            "AUID is empty",
	Result_Get_Me_WalletList_Scan_Error: "GetMeWalletList DB Scan Error",
	Result_Get_Me_PointList_Scan_Error:  "GetMePointList DB Scan Error",
	Result_Get_Me_CoinList_Scan_Error:   "GetMeCoinList DB Scan Error",
	Result_Get_MemberList_Scan_Error:    "USPAU_GetList_Members DB Scan Error",
	Result_Get_Me_Verify_otp_Error:      "Otp Verify Error",
	Result_Error_Db_GetAccountWallets:   "db get account wallet list error",
	Result_Error_Db_NotExistWallets:     "db not exist wallet",
	Result_Get_Me_Wallets_Regiest_Error: "not regist wallet",

	Result_Post_Me_WalletRegist_UnsupportWallet_Error:                 "Unsupported wallet.",
	Result_Post_Me_WalletRegist_AreadyRegistered_Error:                "already have a registered wallet",
	Result_Post_Me_WalletRegist_AreadyRegisteredDB_Error:              "already have a registered wallet",
	Result_Post_Me_WalletRegist_AreadyRegistered_AnotherAccount_Error: "Wallet already registered on another account",
	Result_Post_Me_WalletRegist_NoRegistered_Wallet_Error:             "No registered wallet",
	Result_Post_Me_WalletRegist_System_Error:                          "system error.",
	Result_Post_Me_WalletRegist_DeleteTime_Error:                      "can delete wallet in 24 hours.",
	Result_Post_Me_WalletRegist_Diffrent_Wallet_Error:                 "Not a registered wallet address",
	Result_Post_Me_WalletRegist_NotAllowedWalletType:                  "Not Allowed Wallet Type",

	Result_CoinTransfer_CoinSymbol_Empty: "CoinSymbol is empty",
	Result_CoinTransfer_ToAddress_Empty:  "ToAddress is empty",
	Result_CoinTransfer_Quantity_Empty:   "Quantity is empty",
	Result_CoinTransfer_NotEnough_Coin:   "Not enough Coin Quantity",
	Result_CoinFee_BaseSymbol_Empty:      "Base coin symbol information is empty",
	Result_CoinFee_NotExist:              "not exist coin fee",
	Result_CoinFee_LackOfGas:             "lack of gas",

	Result_PubSub_InternalErr:                      "Internal pubsub error",
	Result_Error_IsMaintenance:                     "It is under maintainence",
	Result_Error_IsSwapMaintenance:                 "Swap is under maintainence",
	Result_Error_IsCoinTransferExternalMaintenance: "External coin transfer is under maintenance",

	Result_DBError:              "Internal DB error",
	Result_Invalid_DBID:         "Invalid DB ID",
	Result_DBError_Unknown:      "Unknown DB error",
	Result_RedisError_Lock_fail: "Redis lock error.",

	Result_Auth_RequireMessage:    "Message is required",
	Result_Auth_RequireSign:       "Sign info is required",
	Result_Auth_InvalidLoginInfo:  "Invalid login info",
	Result_Auth_DontEncryptJwt:    "Auth token create fail",
	Result_Auth_InvalidJwt:        "Invalid jwt token",
	Result_Auth_InvalidWalletType: "Invalid wallet type",
}

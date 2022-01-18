package resultcode

const (
	Result_Success = 0

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

	Result_Get_App_AppID_Empty      = 33001 // Get App AppID Empty
	Result_Get_App_Point_Scan_Error = 33002 // Get App Point DB Scan Error
	Result_Get_App_Coin_Scan_Error  = 33003 // Get App Coin DB Scan Error

	Result_Get_Me_AUID_Empty            = 34001 // GetMeWallets AUID Empty
	Result_Get_Me_WalletList_Scan_Error = 34002 // GetMeWalletList DB Scan Error
	Result_Get_Me_PointList_Scan_Error  = 34003 // GetMePointList DB Scan Error
	Result_Get_Me_CoinList_Scan_Error   = 34004 // GetMeCoinList DB Scan Error

	Result_DBError         = 19000 // db 에러
	Result_Invalid_DBID    = 19001 // 유효하지 못한 database index
	Result_DBError_Unknown = 19002 // 알려지지 않은 db 에러

	Result_Auth_RequireMessage    = 20000
	Result_Auth_RequireSign       = 20001
	Result_Auth_InvalidLoginInfo  = 20002
	Result_Auth_DontEncryptJwt    = 20003
	Result_Auth_InvalidJwt        = 20004
	Result_Auth_InvalidWalletType = 20005
)

var ResultCodeText = map[int]string{
	Result_Success: "success",

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

	Result_Get_App_AppID_Empty:      "AppID is empty",
	Result_Get_App_Point_Scan_Error: "Get App Point DB Scan Error",
	Result_Get_App_Coin_Scan_Error:  "Get App Coin DB Scan Error",

	Result_Get_Me_AUID_Empty:            "AUID is empty",
	Result_Get_Me_WalletList_Scan_Error: "GetMeWalletList DB Scan Error",
	Result_Get_Me_PointList_Scan_Error:  "GetMePointList DB Scan Error",
	Result_Get_Me_CoinList_Scan_Error:   "GetMeCoinList DB Scan Error",

	Result_DBError:         "Internal DB error",
	Result_Invalid_DBID:    "Invalid DB ID",
	Result_DBError_Unknown: "Unknown DB error",

	Result_Auth_RequireMessage:    "Message is required",
	Result_Auth_RequireSign:       "Sign info is required",
	Result_Auth_InvalidLoginInfo:  "Invalid login info",
	Result_Auth_DontEncryptJwt:    "Auth token create fail",
	Result_Auth_InvalidJwt:        "Invalid jwt token",
	Result_Auth_InvalidWalletType: "Invalid wallet type",
}

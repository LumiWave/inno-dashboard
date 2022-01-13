package resultcode

const (
	Result_Success = 0

	Result_Require_PageInfo = 10001 // 유효한 페이지 정보 필요

	Result_Upbit_EmptyCoinSymbol = 12001 // 코인 심볼이 비어있음
	Result_Upbit_EmptyUnit       = 12002 // Unit 이 비어있음
	Result_Upbit_InvalidUnit     = 12003 // Unit 이 유효하지 않음
	Result_Upbit_EmptyCount      = 12004 // Count가 비어있음
	Result_Upbit_EmptyTo         = 12005 // To가  비어있음

	Result_Upbit_TickerMarkets = 13001 // 업비트 시세 Ticker 조회 API 에러
	Result_Upbit_CandleMinutes = 13002 // 업비트 CandleMinutes 조회 API 에러
	Result_Upbit_CandleDays    = 13003 // 업비트 CandleDays 조회 API 에러
	Result_Upbit_CandleWeeks   = 13004 // 업비트 CandleWeeks 조회 API 에러
	Result_Upbit_CandleMonths  = 13005 // 업비트 CandleMonths 조회 API 에러

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

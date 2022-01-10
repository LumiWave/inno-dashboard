package resultcode

const (
	Result_Success = 0

	Result_Require_PageInfo = 10001 // 유효한 페이지 정보 필요

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

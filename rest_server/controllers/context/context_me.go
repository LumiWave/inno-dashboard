package context

import (
	"time"

	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
)

// 지갑 등록 후 해제 가능시간
const DeleteWalletHour = 24 //todo::test후 주석해제
//const DeleteWalletHour = 0

const UserTypeLimit int64 = 2000000

////////////////////////////////////////

// /////// Me Point List
type ReqMePoint struct {
	AUID int64 `json:"au_id" query:"au_id"`
	MUID int64 `json:"mu_id" query:"mu_id"`
}

func (o *ReqMePoint) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}

	return nil
}

type MePoint struct {
	AppID                     int64  `json:"app_id" query:"app_id"`
	PointID                   int64  `json:"point_id" query:"point_id"`
	Quantity                  int64  `json:"quantity"`
	TodayAcqQuantity          int64  `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity         int64  `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	TodayAcqExchangeQuantity  int64  `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
	TodayCnsmExchangeQuantity int64  `json:"today_cnsm_exchange_quantity" query:"today_cnsm_exchange_quantity"`
	ResetDate                 string `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

// sui coin object id 리스트 조회
type ReqCoinObjects struct {
	AUID   int64 `json:"au_id" query:"au_id"`
	CoinID int64 `query:"coin_id"`
}

func NewReqCoinObjects() *ReqCoinObjects {
	return new(ReqCoinObjects)
}

func (o *ReqCoinObjects) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	if o.CoinID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_CoinID_Error)
	}
	return nil
}

type ResCoinObjects struct {
	ObjectIDs []string `json:"object_ids"`
}

////////////////////////////////////////

// /////// Me Coin List
type ReqMeCoin struct {
	AUID   int64 `json:"au_id" query:"au_id"`
	CoinID int64 `json:"coin_id" query:"coin_id"`
}

func (o *ReqMeCoin) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

type MeCoin struct {
	CoinID                    int64     `json:"coin_id" query:"coin_id"`
	BaseCoinID                int64     `json:"base_coin_id" query:"base_coin_id"`
	CoinSymbol                string    `json:"coin_symbol" query:"coin_symbol"`
	WalletAddress             string    `json:"wallet_address" query:"wallet_address"`
	Quantity                  float64   `json:"quantity" query:"quantity"`
	TodayAcqQuantity          float64   `json:"today_acq_quantity" query:"today_acq_quantity"`
	TodayCnsmQuantity         float64   `json:"today_cnsm_quantity" query:"today_cnsm_quantity"`
	TodayAcqExchangeQuantity  float64   `json:"today_acq_exchange_quantity" query:"today_acq_exchange_quantity"`
	TodayCnsmExchangeQuantity float64   `json:"today_cnsm_exchange_quantity" query:"today_cnsm_exchange_quantity"`
	ResetDate                 time.Time `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

// /////// Member
type Member struct {
	MUID       int64 `json:"mu_id"`
	AppID      int64 `json:"app_id"`
	DatabaseID int64 `json:"database_id"`
}

////////////////////////////////////////

// /////// otp : qrcode 용 uri 조회
type MeOtpUri struct {
	OtpUri string `json:"otp_uri"`
}

////////////////////////////////////////

// /////// otp : qrcode 용 uri 조회
type MeOtpVerify struct {
	OtpCode string `json:"otp_code" query:"otp_code"`
}

func NewMeOtpVerify() *MeOtpVerify {
	return new(MeOtpVerify)
}

func (o *MeOtpVerify) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	return nil
}

////////////////////////////////////////

// /////// coin mainnet 보정
type CoinReload struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func NewCoinReload() *CoinReload {
	return new(CoinReload)
}

func (o *CoinReload) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

////////////////////////////////////////

// ///// 등록된 지갑정보요청
type ReqGetWalletRegist struct {
	AUID int64 `json:"au_id" url:"au_id"` // 계정의 UID (Access Token에서 가져옴)
}

func (o *ReqGetWalletRegist) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

type ResGetWalletRegist struct {
	Wallets []*WalletRegistInfo `json:"wallets"`
}

type WalletRegistInfo struct {
	BaseCoinID              int64  `json:"base_coin_id"`
	BaseCoinSymbol          string `json:"base_coin_symbol"`
	WalletTypeID            int64  `json:"wallet_type_id"`
	WalletName              string `json:"wallet_name"`
	IsRegistered            bool   `json:"is_registered"` //등록여부 true:등록되어있음, false:등록안됨
	WalletAddress           string `json:"wallet_address"`
	WalletID                int64  `json:"wallet_id"`
	RegistDT                string `json:"regist_dt"` //등록시간(해제가능시간 24시간체크용)
	LastDeleteWalletAddress string `json:"last_delete_wallet_address"`
	LastDeleteWalletTypeID  int64  `json:"last_delete_wallet_type_id"`
	LastDeleteWalletName    string `json:"last_delete_wallet_name"`
	LastDeleteDT            string `json:"last_delete_dt"`
	UserType                int    `json:"user_type"` //1:구유저(지갑바로등록), 2:신유저(필요할떄등록) 체크용
}

type DBWalletRegist struct {
	WalletID         int64
	BaseCoinID       int64
	WalletAddress    string
	WalletTypeID     int64
	ConnectionStatus int64
	ModifiedDT       string
}

// //// 지갑등록요청
type ReqPostWalletRegist struct {
	AUID          int64  `json:"au_id" url:"au_id"` // 계정의 UID (Access Token에서 가져옴)
	BaseCoinID    int64  `json:"base_coin_id"`
	WalletTypeID  int64  `json:"wallet_type_id"`
	WalletAddress string `json:"wallet_address"`
}

func (o *ReqPostWalletRegist) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	if o.BaseCoinID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_BaseCoinID_Error)
	}
	if o.WalletTypeID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_WalletTypeID_Error)
	}
	if o.WalletAddress == "" {
		return base.MakeBaseResponse(resultcode.Result_Invalid_WalletAddress_Error)
	}
	return nil
}

// //// 지갑삭제요청
type ReqDeleteWalletRegist struct {
	AUID          int64  `json:"au_id" url:"au_id"` // 계정의 UID (Access Token에서 가져옴)
	BaseCoinID    int64  `json:"base_coin_id" query:"base_coin_id"`
	WalletTypeID  int64  `json:"wallet_platform" query:"wallet_type_id"`
	WalletAddress string `json:"wallet_address" query:"wallet_address"`
}

func (o *ReqDeleteWalletRegist) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	if o.BaseCoinID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_CoinID_Error)
	}
	if o.WalletTypeID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_WalletTypeID_Error)
	}
	if o.WalletAddress == "" {
		return base.MakeBaseResponse(resultcode.Result_Invalid_WalletAddress_Error)
	}
	return nil
}

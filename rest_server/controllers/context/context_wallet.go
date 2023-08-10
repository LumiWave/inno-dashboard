package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
)

// /////// Coin Transfer
type ReqWalletRegist struct {
	AUID int64 `json:"au_id" url:"au_id"` // 계정의 UID (Access Token에서 가져옴)
}

func (o *ReqWalletRegist) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	}
	return nil
}

type ResWalletRegist struct {
	IsRegistered bool   `json:"is_registered"` //등록여부 true:등록되어있음, false:등록안됨
	RegistDT     string `json:"regist_dt"`     //등록시간(실시간 24시간체크용)
	UserType     int    `json:"user_type"`     //1:구유저(바로등록), 2:신유저(필요할떄등록)
}

type DBWalletRegist struct {
	BaseCoinID                int64
	WalletID                  int64
	WalletAddress             string
	DisconnectedWalletAddress string
	DisconnectedDT            string
	ModifiedDT                string
}

////////////////////////////////////////

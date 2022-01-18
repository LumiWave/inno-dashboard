package context

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Wallet Info
type ReqMeWallet struct {
	AUID int64 `json:"au_id" query:"au_id"`
}

func (o *ReqMeWallet) CheckValidate() *base.BaseResponse {
	if o.AUID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AUID_Empty)
	}
	return nil
}

type MeWalletInfo struct {
	CoinID        int64     `json:"coin_id" query:"coin_id"`
	WalletAddress string    `json:"wallet_address" query:"wallet_address"`
	Quantity      string    `json:"quantity" query:"quantity"`
	DailyQuantity string    `json:"daily_quantity" query:"daily_quantity"`
	ResetDate     time.Time `json:"reset_date" query:"reset_date"`
}

////////////////////////////////////////

package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
)

///////// Swap Info 전체 포인트, 코인 정보 리스트, swap 가능 정보 조회
type SwapablePoint struct {
	PointID int64 `json:"point_id"`
}

type Swapable struct {
	AppID      int64            `json:"app_id"`
	CoinID     int64            `json:"coin_id"`
	BaseCoinID int64            `json:"base_coin_id"`
	IsEnable   bool             `json:"is_enabled"`
	Points     []*SwapablePoint `json:"points"`
}

type SwapList struct {
	PointList
	AppPoints
	CoinList

	Swapable []*Swapable `json:"swapable"`
}

////////////////////////////////////////

///////// Swap 가능 정보 조회

type ReqSwapEnable struct {
	FromType     string `json:"from_type" query:"from_type"`         // 출발지 타입 "point" | "coin"
	FromID       string `json:"from_id" query:"from_id"`             // 출발지 id
	FromQuantity string `json:"from_quantity" query:"from_quantity"` // 양
}

type RespSwapEnable struct {
	MinimumReceived      int64   `json:"minimum_received,omitempty"`       // 최소 스왑 가능 금액
	PriceImpact          float64 `json:"price_impact,omitempty"`           // 가격 변동률
	LiquidityProviderFee float64 `json:"liquidity_provider_fee,omitempty"` // 수수료

	ToType     string `json:"to_type,omitempty"`     // 도착지 타입 point or coin
	ToID       int64  `json:"to_id,omitempty"`       // 도착지 id
	ToQuantity string `json:"to_quantity,omitempty"` // 양
}

////////////////////////////////////////

///////// Swap 처리
type SwapPoint struct {
	AppID               int64 `json:"app_id"`
	PointID             int64 `json:"point_id"`
	AdjustPointQuantity int64 `json:"adjust_point_quantity"`
}

type SwapCoin struct {
	CoinID             int64   `json:"coin_id"`
	AdjustCoinQuantity float64 `json:"adjust_coin_quantity"`
}

type ReqSwapInfo struct {
	EventID int64  `json:"event_id"` // 3: point->coin,  4: coin->point
	OtpCode string `json:"otp_code"`

	SwapPoint `json:"point"`
	SwapCoin  `json:"coin"`
}

func (o *ReqSwapInfo) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if o.EventID != EventID_toCoin && o.EventID != EventID_toPoint {
		return base.MakeBaseResponse(resultcode.Result_Invalid_EventID_Error)
	}
	if o.AppID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_AppID_Error)
	}
	if o.PointID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_PointID_Error)
	}
	if o.AdjustPointQuantity == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_PointQuantity_Error)
	}
	if o.CoinID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_CoinID_Error)
	}
	if o.AdjustCoinQuantity == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_CoinQuantity_Error)
	}
	// event id에 따라  AdjustPointQuantity AdjustCoinQuantity두 정보 양수 음수 체크
	if o.EventID == EventID_toCoin {
		if !(o.AdjustPointQuantity < 0 && o.AdjustCoinQuantity > 0) {
			return base.MakeBaseResponse(resultcode.Result_Invalid_AdjustQuantity_Error)
		}
	} else if o.EventID == EventID_toPoint {
		if !(o.AdjustCoinQuantity < 0 && o.AdjustPointQuantity > 0) {
			return base.MakeBaseResponse(resultcode.Result_Invalid_AdjustQuantity_Error)
		}
	}

	return nil
}

////////////////////////////////////////

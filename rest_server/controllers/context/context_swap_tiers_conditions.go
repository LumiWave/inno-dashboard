package context

import (
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
)

// / 스왑 가능 메타 데이터 정보

// point to coin
type SwapP2CTier struct {
	// FromID는 전환할 재료의 포인트 ID입니다.
	FromID int64 `json:"from_id"`

	// ToID는 받을 코인의 ID입니다.
	ToID int64 `json:"to_id"`

	// IsEnabled는 해당 전환이 활성화 되어있는지 여부를 나타냅니다.
	TierID int64 `json:"tier_id"`

	// MinimumExchangeQuantity는 최소 전환량을 나타냅니다.
	MinimumExchangeQuantity string `json:"minimum_exchange_quantity"`

	// ExchangeRatio는 받을 전환 비율을 나타냅니다.
	ExchangeRatio float64 `json:"exchange_ratio"`
}

type SwapP2CTierCondition struct {
	FromID        int64  `json:"from_id"`
	ToID          int64  `json:"to_id"`
	TierID        int64  `json:"tier_id"`
	ConditionType int64  `json:"condition_type"` // 2: 토큰, 3:NFT
	ConditionID   int64  `json:"condition_id"`   // type이 2면 CoinID, 3이면 NFTPackID
	Quantity      string `json:"quantity"`
}

////////////////////////////////////////

// swap 조건 체크
type ReqSwapTierCheck struct {
	AUID    int64 `query:"au_id"`
	EventID int64 `query:"tx_type"` // EventID_toP2C ~ EventID_toP2P

	FromID int64 `query:"from_id"`
	ToID   int64 `query:"to_id"`
}

func NewReqSwapTierCheck() *ReqSwapTierCheck {
	return new(ReqSwapTierCheck)
}

func (o *ReqSwapTierCheck) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	} else if ctx.GetValue() == nil && o.AUID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AUID_Empty)
	}

	if o.EventID < EventID_toP2C || o.EventID > EventID_toP2P {
		return base.MakeBaseResponse(resultcode.Result_Invalid_EventID_Error)
	}
	if o.FromID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_FromID_Error)
	}
	if o.ToID <= 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_ToID_Error)
	}

	return nil
}

type SwapTiercheck struct {
	TierID        int64  `json:"tier_id"`
	ConditionType int64  `json:"condition_type"` // 2: 토큰, 3:NFT
	ConditionID   int64  `json:"condition_id"`   // type이 2면 CoinID, 3이면 NFTPackID
	Quantity      string `json:"quantity"`

	IsAchievedCondition bool `json:"is_achieved_condition"`
}

type SwapSelectTierInfo struct {
	SelectTier              int64   `json:"select_tier"`
	MinimumExchangeQuantity string  `json:"minimum_exchange_quantity"`
	ExchangeRatio           float64 `json:"exchange_ratio"`
}

type ResSwapTierCheck struct {
	SwapTiercheck      []*SwapTiercheck    `json:"swap_tier_check"`
	SwapSelectTierInfo *SwapSelectTierInfo `json:"select_tier"`
}

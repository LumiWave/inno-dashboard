package context

import (
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
)

// /////// Swap Info 전체 포인트, 코인 정보 리스트, swap 가능 정보 조회
type SwapablePoint struct {
	PointID int64 `json:"point_id"`
}

// type Swapable struct {
// 	AppID      int64            `json:"app_id"`
// 	CoinID     int64            `json:"coin_id"`
// 	BaseCoinID int64            `json:"base_coin_id"`
// 	IsEnable   bool             `json:"is_enabled"`
// 	Points     []*SwapablePoint `json:"points"`
// }

type SwapList struct {
	PointList
	AppPoints
	CoinList

	SwapAble `json:"swapable"`
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

// /////// Swap 처리
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
	EventID int64  `json:"tx_type"` // 3: point->coin,  4: coin->point
	OtpCode string `json:"otp_code"`

	SwapFromPoint SwapPoint `json:"from_point"`
	SwapToPoint   SwapPoint `json:"to_point"`

	SwapFromCoin SwapCoin `json:"from_coin"`
	SwapToCoin   SwapCoin `json:"to_coin"`
}

func (o *ReqSwapInfo) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if o.EventID != EventID_toP2C && o.EventID != EventID_toC2P && o.EventID != EventID_toC2C && o.EventID != EventID_toP2P {
		return base.MakeBaseResponse(resultcode.Result_Invalid_EventID_Error)
	}
	// event id에 따라  AdjustPointQuantity AdjustCoinQuantity두 정보 양수 음수 체크
	if o.EventID == EventID_toP2C {
		if !(o.SwapFromPoint.AdjustPointQuantity < 0 && o.SwapToCoin.AdjustCoinQuantity > 0) {
			return base.MakeBaseResponse(resultcode.Result_Invalid_AdjustQuantity_Error)
		}
	} else if o.EventID == EventID_toC2P {
		if !(o.SwapFromCoin.AdjustCoinQuantity < 0 && o.SwapToPoint.AdjustPointQuantity > 0) {
			return base.MakeBaseResponse(resultcode.Result_Invalid_AdjustQuantity_Error)
		}
	}

	return nil
}

////////////////////////////////////////

// swap 상태 변경 요청 : (수료 전송 후 tx정보 저장)
type PutSwapStatus struct {
	TxID              int64  `json:"tx_id"`
	TxStatus          int64  `json:"tx_status"`
	TxHash            string `json:"tx_hash"`
	FromWalletAddress string `json:"from_wallet_address"`
}

func (o *PutSwapStatus) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if o.TxStatus < 2 && o.TxStatus > 4 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_TxStatus)
	}
	if len(o.FromWalletAddress) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Invalid_WalletAddress_Error)
	}

	return nil
}

////////////////////////////////////////

// swap 진행 중인지 체크
type ReqSwapInprogress struct {
	AUID int64 `query:"au_id"`
}

func NewReqSwapIniprogress() *ReqSwapInprogress {
	return new(ReqSwapInprogress)
}

func (o *ReqSwapInprogress) CheckValidate(ctx *InnoDashboardContext) *base.BaseResponse {
	if ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	} else if ctx.GetValue() == nil || o.AUID == 0 {
		return base.MakeBaseResponse(resultcode.Result_Get_Me_AUID_Empty)
	}
	return nil
}

////////////////////////////////////////

// / 스왑 가능 메타 데이터 정보
// coin to coin
type SwapCointoCoin struct {
	// FromBaseCoinID는 전환할 재료 코인의 계열 ID입니다.
	FromBaseCoinID int64 `json:"from_base_coin_id"`

	// FromID는 전환할 재료의 ID입니다.
	FromID int64 `json:"from_id"`

	// ToBaseCoinID는 받을 코인의 계열 ID입니다.
	ToBaseCoinID int64 `json:"to_base_coin_id"`

	// ToID는 받을 코인의 ID입니다.
	ToID int64 `json:"to_id"`

	// IsEnabled는 해당 전환이 활성화 되어있는지 여부를 나타냅니다.
	IsEnabled bool `json:"is_enabled"`

	// 화면에 표출 여부
	IsVisible bool `json:"is_visible"`

	// MinimumExchangeQuantity는 최소 전환량을 나타냅니다.
	MinimumExchangeQuantity string `json:"minimum_exchange_quantity"`

	// ExchangeRatio는 받을 전환 비율을 나타냅니다.
	ExchangeRatio float64 `json:"exchange_ratio"`
}

// point to coin
type SwapPointToCoin struct {
	// FromID는 전환할 재료의 포인트 ID입니다.
	FromID int64 `json:"from_id"`

	// ToBaseCoinID는 받을 코인의 계열 ID입니다.
	ToBaseCoinID int64 `json:"to_base_coin_id"`

	// ToID는 받을 코인의 ID입니다.
	ToID int64 `json:"to_id"`

	// IsEnabled는 해당 전환이 활성화 되어있는지 여부를 나타냅니다.
	IsEnabled bool `json:"is_enabled"`

	// 화면에 표출 여부
	IsVisible bool `json:"is_visible"`

	// MinimumExchangeQuantity는 최소 전환량을 나타냅니다.
	MinimumExchangeQuantity string `json:"minimum_exchange_quantity"`

	// ExchangeRatio는 받을 전환 비율을 나타냅니다.
	ExchangeRatio float64 `json:"exchange_ratio"`
}

// coin to point 정보
type SwapCoinToPoint struct {
	// FromBaseCoinID는 전환할 재료 코인의 계열 ID입니다.
	FromBaseCoinID int64 `json:"from_base_coin_id"`

	// FromID는 전환할 재료의 ID입니다.
	FromID int64 `json:"from_id"`

	// ToID는 받을 포인트의 ID입니다.
	ToID int64 `json:"to_id"`

	// IsEnabled는 해당 전환이 활성화 되어있는지 여부를 나타냅니다.
	IsEnabled bool `json:"is_enabled"`

	// 화면에 표출 여부
	IsVisible bool `json:"is_visible"`

	// MinimumExchangeQuantity는 최소 전환량을 나타냅니다.
	MinimumExchangeQuantity string `json:"minimum_exchange_quantity"`

	// ExchangeRatio는 받을 전환 비율을 나타냅니다.
	ExchangeRatio float64 `json:"exchange_ratio"`
}

// point to point 스왑 정보
type SwapPointToPoint struct {
	// FromID는 전환할 재료의 ID입니다.
	FromID int64 `json:"from_id"`

	// ToID는 받을 포인트의 ID입니다.
	ToID int64 `json:"to_id"`

	// IsEnabled는 해당 전환이 활성화 되어있는지 여부를 나타냅니다.
	IsEnabled bool `json:"is_enabled"`

	// 화면에 표출 여부
	IsVisible bool `json:"is_visible"`

	// MinimumExchangeQuantity는 최소 전환량을 나타냅니다.
	MinimumExchangeQuantity string `json:"minimum_exchange_quantity"`

	// ExchangeRatio는 받을 전환 비율을 나타냅니다.
	ExchangeRatio float64 `json:"exchange_ratio"`
}

////////////////////////////////////////

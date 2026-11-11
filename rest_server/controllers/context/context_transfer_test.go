package context

import (
	"math"
	"testing"
)

func TestIsPositiveQuantity(t *testing.T) {
	cases := []struct {
		quantity float64
		want     bool
	}{
		{1, true},
		{0.000001, true},
		{0, false},
		{-1, false}, // 음수 전송
		{-0.5, false},
		{math.Inf(1), false},
		{math.NaN(), false},
	}
	for _, c := range cases {
		if got := IsPositiveQuantity(c.quantity); got != c.want {
			t.Errorf("수량 %v : %v, 기대 %v", c.quantity, got, c.want)
		}
	}
}

func TestIsMemberIdentity(t *testing.T) {
	if IsMemberIdentity(0) {
		t.Error("계정 없는 토큰(AUID 0)이 회원으로 인정됨")
	}
	if IsMemberIdentity(-1) {
		t.Error("음수 AUID 가 회원으로 인정됨")
	}
	if !IsMemberIdentity(100) {
		t.Error("회원 토큰이 거절됨")
	}
}

func TestReqCoinTransferCheckValidate(t *testing.T) {
	ctx := &InnoDashboardContext{}

	ok := &ReqCoinTransfer{CoinID: 1, ToAddress: "0xto", Quantity: 1.5}
	if resp := ok.CheckValidate(ctx); resp != nil {
		t.Fatalf("정상 전송 요청이 거절됨 : %v", resp.Message)
	}
	for _, q := range []float64{0, -1, -100.5} {
		req := &ReqCoinTransfer{CoinID: 1, ToAddress: "0xto", Quantity: q}
		if resp := req.CheckValidate(ctx); resp == nil {
			t.Errorf("수량 %v 인 전송 요청이 통과함", q)
		}
	}
}

package context

import (
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/baseutil/datetime"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
)

type LogID_type int

const (
	LogID_cp              = 1 // 고객사
	LogID_exchange        = 2 // 전환
	LogID_external_wallet = 3 // 외부지갑
)

type EventID_type int

const (
	EventID_add   = 1  // 재화 증가
	EventID_sub   = 2  // 재화 감소
	EventID_toP2C = 3  // 포인트->코인
	EventID_toC2P = 4  // 코인->포인트
	EventID_toC2C = 26 // 코인->코인
)

type ContextKey struct {
	Idx         int64 `json:"idx" query:"idx"`
	CpMemberIdx int64 `json:"cp_member_idx" query:"cp_member_idx"`
}

// page info
type PageInfo struct {
	PageOffset string `json:"page_offset,omitempty" query:"page_offset" validate:"required"`
	PageSize   string `json:"page_size,omitempty" query:"page_size" validate:"required"`
}

// page response
type PageInfoResponse struct {
	PageInfo
	TotalSize string `json:"total_size"`
}

func MakeDt(data *int64) {
	*data = datetime.GetTS2MilliSec()
}

func (o *PageInfo) CheckValidate() *base.BaseResponse {
	if len(o.PageOffset) == 0 || len(o.PageSize) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Require_PageInfo)
	}
	return nil
}

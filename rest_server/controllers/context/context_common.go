package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/datetime"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
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

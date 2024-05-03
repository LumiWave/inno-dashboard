package context

import (
	"github.com/LumiWave/baseapp/base"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/auth"
)

// InnoDashboardContext API의 Request Context
type InnoDashboardContext struct {
	*base.BaseContext
	VerifyValue *auth.VerifyAuthToken
}

func NewInnoDashboardContext(baseCtx *base.BaseContext) interface{} {
	if baseCtx == nil {
		return nil
	}

	ctx := new(InnoDashboardContext)
	ctx.BaseContext = baseCtx

	return ctx
}

// AppendRequestParameter BaseContext 이미 정의되어 있는 ReqeustParameters 배열에 등록
func AppendRequestParameter() {
}

func (o *InnoDashboardContext) SetVerifyAuthToken(value *auth.VerifyAuthToken) {
	o.VerifyValue = value
}

func (o *InnoDashboardContext) GetValue() *auth.VerifyAuthToken {
	return o.VerifyValue
}

package externalapi

import (
	"github.com/LumiWave/baseapp/base"
	baseconf "github.com/LumiWave/baseapp/config"
	"github.com/LumiWave/baseutil/log"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/auth"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/commonapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/labstack/echo"
)

type ExternalAPI struct {
	base.BaseController

	conf    *config.ServerConfig
	apiConf *baseconf.APIServer
	echo    *echo.Echo
}

func PreCheck(c echo.Context) base.PreCheckResponse {
	conf := config.GetInstance()
	if err := base.SetContext(c, &conf.Config, context.NewInnoDashboardContext); err != nil {
		log.Errorf("%v", err)
		return base.PreCheckResponse{
			IsSucceed: false,
		}
	}

	if model.GetMaintenance() {
		res := base.MakeBaseResponse(resultcode.Result_Error_IsMaintenance)
		return base.PreCheckResponse{
			IsSucceed: false,
			Response:  res,
		}
	}

	// auth token 검증
	if conf.Auth.AuthEnable {
		author, ok := c.Request().Header["Authorization"]
		if !ok {
			// auth token 오류 리턴
			res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)

			return base.PreCheckResponse{
				IsSucceed: false,
				Response:  res,
			}
		}

		// "Bearer " 뒤에 토큰이 없으면 잘라낼 때 panic 이 나므로 길이를 먼저 확인한다.
		if len(author) == 0 || len(author[0]) <= 7 {
			res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)
			return base.PreCheckResponse{
				IsSucceed: false,
				Response:  res,
			}
		}

		if value, returnValue, err := auth.CheckAuthToken(author[0][7:]); err != nil {
			res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)
			return base.PreCheckResponse{
				IsSucceed: false,
				Response:  res,
			}
		} else {
			if value != nil {
				// 대시보드의 토큰 필요 API 는 전부 회원 본인용이다. 계정(au_id)이 없는 토큰
				// (앱 토큰, 인증 서버가 빈 값으로 성공을 돌려준 경우)은 AUID 0 으로 진행되지 않게 거절한다.
				if !context.IsMemberIdentity(value.AUID) {
					res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)
					return base.PreCheckResponse{
						IsSucceed: false,
						Response:  res,
					}
				}
				base.GetContext(c).(*context.InnoDashboardContext).SetVerifyAuthToken(value)
				log.Debugf("from : [companyid:%v][appid:%v][logintype:%v][au_id:%v]", value.CompanyID, value.AppID, value.LoginType, value.AUID)
			} else {
				res := new(base.BaseResponse)
				res.Return = returnValue.Return
				res.Message = returnValue.Message
				return base.PreCheckResponse{
					IsSucceed: false,
					Response:  res,
				}
			}
		}
	}

	return base.PreCheckResponse{
		IsSucceed: true,
	}
}

func (o *ExternalAPI) Init(e *echo.Echo) error {
	o.echo = e
	o.BaseController.PreCheck = PreCheck

	if err := o.MapRoutes(o, e, o.apiConf.Routes); err != nil {
		return err
	}

	return nil
}

func (o *ExternalAPI) GetConfig() *baseconf.APIServer {
	o.conf = config.GetInstance()
	o.apiConf = &o.conf.APIServers[1]
	return o.apiConf
}

func NewAPI() *ExternalAPI {
	return &ExternalAPI{}
}

func (o *ExternalAPI) GetHealthCheck(c echo.Context) error {
	return commonapi.GetHealthCheck(c)
}

func (o *ExternalAPI) GetVersion(c echo.Context) error {
	return commonapi.GetVersion(c, o.BaseController.MaxVersion)
}

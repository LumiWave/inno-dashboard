package app

import (
	"fmt"
	"sync"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/auth"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/externalapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/internalapi"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/point_manager_server"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/upbit"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/schedule"
)

type ServerApp struct {
	base.BaseApp
	conf *config.ServerConfig

	sysMonitor *schedule.SystemMonitor
}

func (o *ServerApp) Init(configFile string) (err error) {
	o.conf = config.GetInstance(configFile)
	base.AppendReturnCodeText(&resultcode.ResultCodeText)
	context.AppendRequestParameter()

	// if err := o.InitScheduler(); err != nil {
	// 	return err
	// }
	auth.InitHttpClient()
	o.InitPointManagerServer(o.conf)

	if err := o.NewDB(o.conf); err != nil {
		return err
	}

	o.InitUpbit()

	return err
}

func (o *ServerApp) CleanUp() {
	fmt.Println("CleanUp")
}

func (o *ServerApp) Run(wg *sync.WaitGroup) error {
	return nil
}

func (o *ServerApp) GetConfig() *baseconf.Config {
	return &o.conf.Config
}

func NewApp() (*ServerApp, error) {
	app := &ServerApp{}

	intAPI := internalapi.NewAPI()
	extAPI := externalapi.NewAPI()

	if err := app.BaseApp.Init(app, intAPI, extAPI); err != nil {
		return nil, err
	}

	return app, nil
}

func (o *ServerApp) InitScheduler() error {
	o.sysMonitor = schedule.GetSystemMonitor()
	return nil
}

func (o *ServerApp) InitUpbit() {
	upbit.InitUpbitInfo()
}

func (o *ServerApp) InitPointManagerServer(conf *config.ServerConfig) {
	pointMgrServer := conf.PointMgrServer
	hostInfo := point_manager_server.HostInfo{
		IntHostUri: pointMgrServer.InternalpiDomain,
		ExtHostUri: pointMgrServer.ExternalpiDomain,
		IntVer:     pointMgrServer.InternalVer,
		ExtVer:     pointMgrServer.ExternalVer,
	}
	point_manager_server.NewPointManagerServerInfo("", hostInfo)
}

func (o *ServerApp) NewDB(conf *config.ServerConfig) error {
	return model.InitDB(conf)
}

package app

import (
	"fmt"
	"sync"

	"github.com/LumiWave/baseapp/base"
	baseconf "github.com/LumiWave/baseapp/config"
	"github.com/LumiWave/inno-dashboard/rest_server/config"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/auth"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/context"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/externalapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/internalapi"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/resultcode"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/inno_market"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/inno_web_server"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/servers/point_manager_server"
	"github.com/LumiWave/inno-dashboard/rest_server/controllers/upbit"
	"github.com/LumiWave/inno-dashboard/rest_server/model"
	"github.com/LumiWave/inno-dashboard/rest_server/schedule"
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
	o.InitMarketServer(o.conf)
	o.InitWebInnoServer(o.conf)

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

func (o *ServerApp) InitMarketServer(conf *config.ServerConfig) error {
	return inno_market.InitMarketServer(conf)
}

func (o *ServerApp) InitWebInnoServer(conf *config.ServerConfig) error {
	return inno_web_server.InitWebInno(conf)
}

func (o *ServerApp) NewDB(conf *config.ServerConfig) error {
	return model.InitDB(conf)
}

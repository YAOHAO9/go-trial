package application

import (
	"math/rand"
	"time"

	"github.com/YAOHAO9/pine/application/config"
	"github.com/YAOHAO9/pine/logger"
	RpcServer "github.com/YAOHAO9/pine/rpc/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Application app
type Application struct {
}

// Start start application
func (app Application) Start() {
	RpcServer.Start()
}

var app *Application

// CreateApp 创建app
func CreateApp() *Application {

	parseConfig()

	logger.SetLogMode(config.GetServerConfig().LogType)

	if app != nil {
		return app
	}
	app = &Application{}

	return app
}

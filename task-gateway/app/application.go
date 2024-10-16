package app

import (
	"sync"
	"task-gateway/config"
	"task-gateway/db"
	"task-gateway/logger"
	"task-gateway/web"
	"task-gateway/web/utils"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	config.LoadConfig()
	conf := config.GetConfig()
	logger.SetupLogger(conf.ServiceName)
	db.InitDB()
	// rabbitmq.InitRabbitMQ()
	// clients.InitGrpcClients()
	utils.InitValidator()
}

func (app *Application) Run() {
	web.StartServer(&app.wg)
}

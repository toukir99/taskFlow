package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"task-gateway/config"
	"task-gateway/web/middlewares"
)

func StartServer(wg *sync.WaitGroup) {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	InitRoutes(mux, manager)
	//swagger.SetupSwagger(mux, manager)
	wg.Add(1)
	go func() {
		defer wg.Done()
		conf := config.GetConfig()
		addr := fmt.Sprintf(":%d", conf.HttpPort)
		slog.Info(fmt.Sprintf("Listening at %s", addr))
		if err := http.ListenAndServe(addr, mux); err != nil {
			slog.Error(err.Error())
		}
	}()

}

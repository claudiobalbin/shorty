package app

import (
	"context"
	"log"
	"net/http"
	"shorty/configs"
	routes "shorty/routes/api"
)

type App struct {
	http.Server
}

func MakeApp() *App {
	port := configs.GetSettings()["PORT"]
	handler := routes.RoutesV1()

	log.Printf("Starting server on :%s", port)

	server := &App{
		Server: http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}

	return server
}

func (s *App) Start() error {
	return s.ListenAndServe()
}

func (s *App) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}

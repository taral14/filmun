package server

import (
	"context"
	"log"
	"net/http"
)

type App struct {
	srv *http.Server
}

func (app App) Start(ctx context.Context) {
	app.srv.ListenAndServe()
}

func (app App) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return app.srv.Shutdown(ctx)
}

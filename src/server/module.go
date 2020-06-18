package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	Module = fx.Options(
		fx.Provide(initRoutes),
		fx.Provide(initApp),
	)
}

func initApp(router *gin.Engine) *App {
	srv := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}
	return &App{
		srv: srv,
	}
}

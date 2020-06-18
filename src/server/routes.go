package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/taral14/filmun/src/auth"
	"github.com/taral14/filmun/src/film"
	"github.com/taral14/filmun/src/person"
	"github.com/taral14/filmun/src/user"
)

func initRoutes(personH *person.Handler, filmH *film.Handler, authH *auth.Handler, userH *user.Handler) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = viper.GetStringSlice("cors.allow_origins")
	config.AllowHeaders = []string{"*"}
	cors := cors.New(config)
	router.Use(cors)

	authH.RegisterHTTPEndpoints(router)
	personH.RegisterHTTPEndpoints(router)
	filmH.RegisterHTTPEndpoints(router)
	userH.RegisterHTTPEndpoints(router)

	return router
}

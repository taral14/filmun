package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/taral14/filmun/src/auth"
	"github.com/taral14/filmun/src/film"
	"github.com/taral14/filmun/src/person"
	"github.com/taral14/filmun/src/user"
	"go.uber.org/fx"
)

func init() {
	viper.SetEnvPrefix("FILMS_")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	app := fx.New(
		person.Module,
		film.Module,
		user.Module,
		auth.Module,

		fx.Provide(initDB),

		fx.Invoke(register),
	)
	app.Run()
}

func register(lifecycle fx.Lifecycle, personH *person.Handler, filmH *film.Handler, authH *auth.Handler) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Println("BEFORE REQUEST")
		//c.Set("example", "12345")
		c.Next()
		fmt.Println("AFTER REQUEST")
	})

	authH.RegisterHTTPEndpoints(router)
	personH.RegisterHTTPEndpoints(router)
	filmH.RegisterHTTPEndpoints(router)

	srv := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Shutting down server...")
				return srv.Shutdown(ctx)
			},
		},
	)
}

func initDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", viper.GetString("mysql.connection"))
	if err != nil {
		log.Panic("init DB: " + err.Error())
	}
	return db
}

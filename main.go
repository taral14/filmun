package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/taral14/filmun/src/auth"
	"github.com/taral14/filmun/src/film"
	"github.com/taral14/filmun/src/person"
	"github.com/taral14/filmun/src/server"
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
		server.Module,
		person.Module,
		film.Module,
		user.Module,
		auth.Module,

		fx.Provide(initDB),
		fx.Invoke(register),
	)
	app.Run()
}

func register(lifecycle fx.Lifecycle, app *server.App) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go app.Start(ctx)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Shutting down server...")
				return app.Shutdown(ctx)
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

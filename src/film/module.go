package film

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	Module = fx.Options(
		fx.Provide(initFilmHandler),
		fx.Provide(initFilmService),
		fx.Provide(initFilmRepository),
	)
}

func initFilmHandler(service *Service) *Handler {
	return NewHandler(service)
}

func initFilmService(db *sqlx.DB) *Service {
	repo := NewMysqlRepository(db)
	service := NewService(repo)
	return service
}

func initFilmRepository(db *sqlx.DB) *MysqlRepository {
	return NewMysqlRepository(db)
}

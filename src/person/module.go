package person

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	Module = fx.Options(
		fx.Provide(initPersonHandler),
		fx.Provide(initPersonRepository),
		fx.Provide(initPersonService),
	)
}

func initPersonHandler(service *Service) *Handler {
	return NewHandler(service)
}

func initPersonRepository(db *sqlx.DB) *MysqlRepository {
	return NewMysqlRepository(db)
}

func initPersonService(repo *MysqlRepository) *Service {
	return NewService(repo)
}

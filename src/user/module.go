package user

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	Module = fx.Options(
		fx.Provide(initUserService),
		fx.Provide(initUserRepository),
		fx.Provide(initHandler),
	)
}

func initHandler(s *Service) *Handler {
	return NewHandler(s)
}

func initUserService(db *sqlx.DB) *Service {
	repo := NewMysqlRepository(db)
	service := NewService(repo)
	return service
}

func initUserRepository(db *sqlx.DB) *MysqlRepository {
	return NewMysqlRepository(db)
}

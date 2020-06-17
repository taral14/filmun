package auth

import (
	"github.com/taral14/filmun/src/user"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	Module = fx.Options(
		fx.Provide(initAuthService),
		fx.Provide(initAuthHandler),
		fx.Provide(initJwtTokenGenerator),
		fx.Provide(initPwdEncoder),
	)
}

func initAuthService(s *user.Service, gen *JwtTokenGenerator, enc *PasswordEncoder) *Service {
	return NewService(s, gen, enc)
}

func initJwtTokenGenerator() *JwtTokenGenerator {
	return NewJwtTokenGenerator([]byte("key"), 10)
}

func initAuthHandler(service *Service) *Handler {
	return NewHandler(service)
}

func initPwdEncoder() *PasswordEncoder {
	return NewPasswordEncoder()
}

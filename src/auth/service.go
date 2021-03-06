package auth

import (
	"fmt"

	"github.com/taral14/filmun/src/entity"
)

type userService interface {
	FindByUsername(username string) (*entity.User, error)
}

type jwtProvider interface {
	CreateToken(user *entity.User) (string, error)
	ParseToken(tokenString string) (*Token, error)
}

type pwdEncoder interface {
	encodePassword(password string) (string, error)
	isPasswordValid(hashPwd, pwd string) bool
}

type Service struct {
	userService userService
	jwtProvider jwtProvider
	pwdEncoder  pwdEncoder
}

func NewService(userServ userService, gen jwtProvider, enc pwdEncoder) *Service {
	return &Service{
		userService: userServ,
		jwtProvider: gen,
		pwdEncoder:  enc,
	}
}

func (s *Service) LogIn(username, password string) (*entity.User, string, error) {
	var token string
	user, err := s.userService.FindByUsername(username)
	if err != nil {
		return user, token, fmt.Errorf("User not found: %w", err)
	}
	if !s.pwdEncoder.isPasswordValid(user.Password, password) {
		return user, token, fmt.Errorf("Incorrect user password by user %v", user.ID)
	}
	token, err = s.jwtProvider.CreateToken(user)
	if err != nil {
		return user, token, fmt.Errorf("Cant create token by user: %w", err)
	}
	return user, token, nil
}

func (s *Service) GetUserIdByToken(tokenString string) (int, error) {
	token, err := s.jwtProvider.ParseToken(tokenString)
	if err != nil {
		return 0, fmt.Errorf("AuthService => %w", err)
	}
	return token.UserId, nil
}

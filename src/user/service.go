package user

import (
	"fmt"

	"github.com/taral14/filmun/src/entity"
)

type repository interface {
	FindByUsername(username string) (*entity.User, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByUsername(username string) (*entity.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return user, fmt.Errorf("UserService.FindByUsername: %v", err)
	}
	return user, nil
}

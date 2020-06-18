package film

import "github.com/taral14/filmun/src/entity"

type repository interface {
	FindByActorId(actorId int, limit, offset int) ([]entity.Film, error)
	FindByDirectorId(directorId int, limit, offset int) ([]entity.Film, error)
	FindAll(limit, offset int) ([]entity.Film, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (uc *Service) ListFilms() ([]entity.Film, error) {
	return uc.repo.FindAll(100, 0)
}

func (uc *Service) ListActorFilms(actorId int) ([]entity.Film, error) {
	return uc.repo.FindByActorId(actorId, 100, 0)
}

func (uc *Service) ListDirectorFilms(directorId int) ([]entity.Film, error) {
	return uc.repo.FindByDirectorId(directorId, 100, 0)
}

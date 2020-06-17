package person

import (
	"github.com/taral14/filmun/src/entity"
)

type repository interface {
	FindOneById(id int) (entity.Person, error)
	FindAll(limit, offset int) ([]entity.Person, error)
	FindByName(name string, limit, offset int) ([]entity.Person, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (uc *Service) GetPerson(id int) (entity.Person, error) {
	return uc.repo.FindOneById(id)
}

func (uc *Service) ListPersons(page int) ([]entity.Person, error) {
	limit := 20
	offset := (page - 1) * limit
	return uc.repo.FindAll(limit, offset)
}

func (uc *Service) SearchByTerm(term string) ([]entity.Person, error) {
	if len(term) == 0 {
		return []entity.Person{}, nil
	}
	limit := 20
	return uc.repo.FindByName(term, limit, 0)
}

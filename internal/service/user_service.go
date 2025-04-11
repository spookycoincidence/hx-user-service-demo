package service

import (
	"github.com/spookycoincidence/hx-user-service-demo/internal/repository"

	"github.com/spookycoincidence/hx-user-service-demo/internal/domain"
)

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id int64) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUser(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

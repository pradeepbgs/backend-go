package service

import (
	"github.com/pradeepbgs/internals/model"
	"github.com/pradeepbgs/internals/repository"
)

type UserServiceInterface interface {
	GetUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
}

type userService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserById (id int) (*model.User, error) {
	return s.repo.FindById(id)
}

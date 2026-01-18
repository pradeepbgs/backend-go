package service

import (
	"github.com/pradeepbgs/internal/model"
	"github.com/pradeepbgs/internal/repository"
)

type UserServiceInterface interface {
	GetUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser (name string, email string) (*model.User, error)
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

func (s *userService) CreateUser(name string, email string) (*model.User, error) {
	createdUser, err := s.repo.CreateUser(name,email)
	if err != nil {
		return nil,err
	}
	return createdUser,nil
}
package service

import (
	"errors"

	"github.com/pradeepbgs/internal/model"
	"github.com/pradeepbgs/internal/repository"
)

type UserServiceInterface interface {
	GetUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser(name string, email string) (*model.User, error)
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

func (s *userService) GetUserById(id int) (*model.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(name string, email string) (*model.User, error) {
	existingUser, err := s.repo.FindByEmail(email)

	if err != nil {
		if err.Error() == "no rows in result set" {
			createdUser, err := s.repo.CreateUser(name, email)
			if err != nil {
				return nil, err
			}
			return createdUser, nil
		}
		return nil, err
	}

	if existingUser != nil && existingUser.ID != 0 {
		return nil, errors.New("user with this email already exists")
	}

	createdUser, err := s.repo.CreateUser(name, email)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

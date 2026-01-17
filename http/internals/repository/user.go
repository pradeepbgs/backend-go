package repository

import "github.com/pradeepbgs/internals/model"

// interface
type UserRepositoryInterface interface {
	FindAll() ([]model.User,error)
	FindById(id int) (*model.User,error)
}

type userRepository struct{}

func FakeUserRespository() UserRepositoryInterface {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	return []model.User{
		{ID: 1, Name: "Pradeep", Email: "p@dev.com"},
		{ID: 2, Name: "Alex", Email: "a@dev.com"},
	}, nil
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	return &model.User{
		ID: id,
		Name: "Pradeep",
		Email: "p@dev.com",
	}, nil
}
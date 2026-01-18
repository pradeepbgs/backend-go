package repository

import (
	"context"

	sqlc "github.com/pradeepbgs/internal/db"
	"github.com/pradeepbgs/internal/model"
)

// interface
type UserRepositoryInterface interface {
	FindAll() ([]model.User, error)
	FindById(id int) (*model.User, error)
	CreateUser(name string, email string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userRepository struct {
	queries *sqlc.Queries
}

func NewUserRespository(queries *sqlc.Queries) UserRepositoryInterface {
	return &userRepository{queries: queries}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	users, err := r.queries.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}

	var result []model.User
	for _, user := range users {
		result = append(result, model.User{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return result, nil
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	user, err := r.queries.GetUserById(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepository) CreateUser(name string, email string) (*model.User, error) {
	user, err := r.queries.CreateUser(context.Background(), sqlc.CreateUserParams{Name: name, Email: email})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	user, err := r.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

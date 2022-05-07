package repository

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"context"
)

type UserRepository struct {
	s *server.Server
	ctx context.Context
	queries *db.Queries
}

func NewUserRepository(s *server.Server, ctx context.Context, queries *db.Queries) *UserRepository {
	return &UserRepository{s, ctx, queries}
}

func (repo *UserRepository) CreateUser(params *db.CreateUserParams) (*db.User, error) {
	user, err := repo.queries.CreateUser(repo.ctx, *params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetAllUsers() ([]db.User, error) {
	users, err := repo.queries.GetUsers(repo.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
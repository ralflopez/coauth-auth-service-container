package repository

import (
	"coauth/pkg/db"
	"context"
)

type UserRepository struct {
	ctx context.Context
	queries *db.Queries
}

func NewUserRepository(ctx context.Context, queries *db.Queries) *UserRepository {
	return &UserRepository{ctx, queries}
}

func (repo *UserRepository) CreateUser(params *db.CreateUserParams) (*db.User, error) {
	user, err := repo.queries.CreateUser(repo.ctx, *params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
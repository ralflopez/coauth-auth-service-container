package repository

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"context"

	"github.com/google/uuid"
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

func (repo *UserRepository) GetUser(id string) (*db.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := repo.queries.GetUser(repo.ctx, uuid)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*db.User, error) {
	user, err := repo.queries.GetUserByEmail(repo.ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) DeleteUser(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = repo.queries.DeleteUser(repo.ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}
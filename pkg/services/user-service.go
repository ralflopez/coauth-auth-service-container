package services

import (
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo}
}

func (u *UserService) CreateUser(dto *userdto.CreateUserDTO) (*db.User, error) {
	var role db.Role
	role.Scan(dto.Role)

	// Persist
	user, err := u.repo.CreateUser(&db.CreateUserParams{
		ID: uuid.New(),
		Name: dto.Name,
		Email: dto.Email,
		Password: dto.Password,
		Role: role,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/repository"

	"github.com/google/uuid"
)

type UserService struct {
	s *server.Server
	repo *repository.UserRepository
}

func NewUserService(s *server.Server, repo *repository.UserRepository) *UserService{
	return &UserService{s, repo}
}

func (service *UserService) CreateUser(dto *userdto.CreateUserDTO) (*db.User, error) {
	var role db.Role
	role.Scan(dto.Role)

	// Persist
	user, err := service.repo.CreateUser(&db.CreateUserParams{
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

func (service *UserService) GetUsers() ([]db.User, error) {
	return service.repo.GetAllUsers()
}
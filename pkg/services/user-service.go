package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/repository"
	"coauth/pkg/utils"

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

	// Validate Role
	err := role.Scan(dto.Role)
	if err != nil {
		return nil, err
	}
	if role == "" {
		role = db.RoleMember
	}
	dto.Role = string(role)

	// Persist
	passwordHash, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	user, err := service.repo.CreateUser(&db.CreateUserParams{
		ID: uuid.New(),
		Name: dto.Name,
		Email: dto.Email,
		Password: passwordHash,
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

func (service *UserService) GetUser(id string) (*db.User, error) {
	return service.repo.GetUser(id)
}

func (service *UserService) GetUserByEmail(email string) (*db.User, error) {
	return service.repo.GetUserByEmail(email)
}

func (service *UserService) DeleteUser(id string) error {
	return service.repo.DeleteUser(id)
}
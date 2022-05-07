package handlers

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
	"coauth/pkg/utils"
	"net/http"
)

type UserHandler struct {
	s *server.Server
	service *services.UserService
}

func NewUserHandler(s *server.Server,service *services.UserService) *UserHandler {
	return &UserHandler{s, service}
}

func (handler *UserHandler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	var createUserDTO userdto.CreateUserDTO
	var role db.Role

	// Unmarshal
	handler.s.Decode(w, r, &createUserDTO)
	handler.s.Logger.Printf("Request Body: %v\n", createUserDTO)

	// Validate: JSON
	err := utils.ValidateStruct(&createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("Validation Error: %v\n", err.Error())
		exceptions.ThrowBadRequestException(w, err.Error())
		return
	}

	// Validate: Role
	err = role.Scan(createUserDTO.Role)
	if err != nil {
		handler.s.Logger.Printf("Validation Error: %v\n", err.Error())
		exceptions.ThrowBadRequestException(w, err.Error())
		return
	}
	if role == "" {
		role = db.RoleMember
	}
	createUserDTO.Role = string(role)

	// Persist
	user, err := handler.service.CreateUser(&createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("Persistence Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, err.Error())
		return
	}
	
	// Return as DTO
	response := &userdto.UserDTO{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	}
	handler.s.Logger.Printf("Response: %v\n", response)
	handler.s.Respond(w, response, http.StatusOK)
}

func (handler *UserHandler) HandleUsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.GetUsers()
	if err != nil {
		handler.s.Logger.Printf("Fetch Error: %v", err.Error())
		exceptions.ThrowInternalServerError(w, err.Error())
		return
	}

	// Convert to DTO
	dtos := []userdto.UserDTO{}
	for _, user := range users {
		userDto := &userdto.UserDTO{
			Id: user.ID.String(),
			Name: user.Name,
			Email: user.Email,
			Role: string(user.Role),
		}
		dtos = append(dtos, *userDto)
	}

	handler.s.Respond(w, dtos, http.StatusOK)
}
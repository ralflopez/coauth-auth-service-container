package handlers

import (
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
	"coauth/pkg/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDTO userdto.CreateUserDTO
	var role db.Role

	// Unmarshal
	utils.JSONToStuct(r.Body, &createUserDTO)

	// Validate: JSON
	err := utils.ValidateStruct(&createUserDTO)
	if err != nil {
		exceptions.ThrowBadRequestException(w, err.Error())
		return
	}

	// Validate: Role
	err = role.Scan(createUserDTO.Role)
	if err != nil {
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
		exceptions.ThrowInternalServerError(w, err.Error())
		return
	}
	
	// Return as DTO
	utils.StructToJSON(w, &userdto.UserDTO{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	})
}

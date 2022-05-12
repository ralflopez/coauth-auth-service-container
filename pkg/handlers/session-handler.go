package handlers

import (
	"coauth/pkg/config/server"
	"coauth/pkg/dtos/sessiondto"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
	"coauth/pkg/utils"
	"fmt"
	"net/http"
)

type SessionHandler struct {
	s *server.Server
	sessionService *services.SessionService
}

func NewSessionHandler(s *server.Server, service *services.SessionService) *SessionHandler {
	return &SessionHandler{s, service}
}

func (handler *SessionHandler) HandleSessionLogin(w http.ResponseWriter, r *http.Request) {
	var loginDTO sessiondto.LoginDTO
	handler.s.Decode(w, r, &loginDTO)

	handler.s.Logger.Printf("Request Body: %v\n", loginDTO)

	// Validation
	err := utils.ValidateStruct(&loginDTO)
	if err != nil {
		handler.s.Logger.Printf("Validation Error: %v\n", err.Error())
		exceptions.ThrowBadRequestException(w, fmt.Sprintf("Validation error: %v\n", err.Error()))
		return
	}

	// Fetch
	user, err := handler.sessionService.Login(w, r, &loginDTO)
	if err != nil {
		handler.s.Logger.Printf("Fetch Error: %v", err.Error())
		exceptions.ThrowBadRequestException(w, fmt.Sprintf("Fetch error: %v\n", err.Error()))
		return
	}

	userDTO := &userdto.UserDTO{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	}

	handler.s.Respond(w, userDTO, http.StatusOK)
}

func (handler *SessionHandler) HandleSessionSignup(w http.ResponseWriter, r *http.Request) {
	var createUserDTO *userdto.CreateUserDTO
	handler.s.Decode(w, r, &createUserDTO)

	handler.s.Logger.Printf("Request Body: %v\n", createUserDTO)

	// Persist
	user, err := handler.sessionService.Signup(w, r, createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("Persistence Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Persistence error: %v\n", err.Error()))
		return
	}

	userdto := &userdto.UserDTO{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	}

	handler.s.Respond(w, userdto, http.StatusOK)
}

func (handler *SessionHandler) HandleSessionLogout(w http.ResponseWriter, r *http.Request) {
	err := handler.sessionService.Logout(w, r)
	if err != nil {
		handler.s.Logger.Printf("Session Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Session error: %v\n", err.Error()))
		return
	}

	handler.s.Respond(w, nil, http.StatusOK)
}

func (handler *SessionHandler) HandleSessionUser(w http.ResponseWriter, r *http.Request) {
	user, err := handler.sessionService.GetLoggedInUser(w, r)
	if err != nil {
		handler.s.Logger.Printf("Session Error: %v\n", err.Error())
		exceptions.ThrowForbiddenException(w, fmt.Sprintf("Session error: %v\n", err.Error()))
		return
	}

	userDTO := &userdto.UserDTO{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	}

	handler.s.Respond(w, userDTO, http.StatusOK)
}

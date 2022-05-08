package handlers

import (
	"coauth/pkg/config/server"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
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
	
}

func (handler *SessionHandler) HandleSessionSignup(w http.ResponseWriter, r *http.Request) {
	var createUserDTO *userdto.CreateUserDTO
	handler.s.Decode(w, r, &createUserDTO)

	handler.s.Logger.Printf("Request Body: %v\n", createUserDTO)

	// Persist
	user, err := handler.sessionService.Signup(w, r, createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("Persistence Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, err.Error())
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
		handler.s.Logger.Panicf("Session Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, err.Error())
		return
	}

	handler.s.Respond(w, nil, http.StatusOK)
}
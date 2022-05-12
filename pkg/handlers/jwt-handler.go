package handlers

import (
	"coauth/pkg/config/server"
	"coauth/pkg/dtos/sessiondto"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
	"fmt"
	"net/http"
)

type JwtHandler struct {
	s          *server.Server
	jwtService *services.JwtService
}

func NewJwtHandler(s *server.Server, jwtService *services.JwtService) *JwtHandler {
	return &JwtHandler{s, jwtService}
}

func (handler *JwtHandler) HandleJwtSignup(w http.ResponseWriter, r *http.Request) {
	var createUserDTO *userdto.CreateUserDTO
	handler.s.Decode(w, r, &createUserDTO)

	if createUserDTO == nil {
		handler.s.Logger.Printf("request body invalid")
		exceptions.ThrowInternalServerError(w, "Request body invalid")
		return
	}

	handler.s.Logger.Printf("Request Body: %v\n", createUserDTO)

	user, err := handler.jwtService.Signup(createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("User Creation Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Signup error: %v\n", err.Error()))
		return
	}

	jwtResponse, err := handler.jwtService.GenerateTokens(user.ID.String())
	if err != nil {
		handler.s.Logger.Printf("Token Generator Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Token generation error: %v\n", err.Error()))
		return
	}

	handler.s.Respond(w, jwtResponse, http.StatusOK)
}

func (handler *JwtHandler) HandleJwtLogin(w http.ResponseWriter, r *http.Request) {
	var loginDTO *sessiondto.LoginDTO
	handler.s.Decode(w, r, &loginDTO)

	if loginDTO == nil {
		handler.s.Logger.Printf("request body invalid")
		exceptions.ThrowInternalServerError(w, "Request body invalid")
		return
	}

	handler.s.Logger.Printf("Request Body: %v\n", loginDTO)

	user, err := handler.jwtService.Login(loginDTO)
	if err != nil {
		handler.s.Logger.Printf("Login Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Login error: %v\n", err.Error()))
		return
	}

	jwtResponse, err := handler.jwtService.GenerateTokens(user.ID.String())
	if err != nil {
		handler.s.Logger.Printf("Token Generator Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, fmt.Sprintf("Token generation error: %v\n", err.Error()))
		return
	}

	handler.s.Respond(w, jwtResponse, http.StatusOK)
}

func (handler *JwtHandler) HandleJwtUser(w http.ResponseWriter, r *http.Request) {

	user, err := handler.jwtService.GetLoggedInUser(r)
	if err != nil {
		handler.s.Logger.Printf("Jwt error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, err.Error())
		return
	}

	userDTO := &userdto.UserDTO{
		Id:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	handler.s.Respond(w, userDTO, http.StatusOK)
}

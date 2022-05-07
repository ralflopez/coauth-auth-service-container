package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/userdto"
	"fmt"
	"net/http"
)

type SessionService struct {
	s *server.Server
	userService *UserService
}

func NewSessionService(s *server.Server, userService *UserService) *SessionService {
	return &SessionService{s, userService}
}

func (service *SessionService) Signup(w http.ResponseWriter, r *http.Request, dto *userdto.CreateUserDTO) (*db.User, error) {
	// TODO: recheck if still needed because email column should be set to UNIQUE
	// Validate Email
	_, err := service.userService.GetUserByEmail(dto.Email)
	if err == nil {
		return nil, fmt.Errorf("email already taken")
	}
	
	// Save to DB
	user, err := service.userService.CreateUser(dto)
	if err != nil {
		return nil, err
	}

	// Save in Session
	session, err := service.s.SessionStore.Get(r, "user-session")
	if err != nil {
		return nil, err
	}
	session.Values["userId"] = user.ID.String()
	err = session.Save(r, w)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

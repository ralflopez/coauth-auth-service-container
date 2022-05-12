package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/sessiondto"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/utils"
	"fmt"
	"net/http"
)

type SessionService struct {
	s           *server.Server
	userService *UserService
}

func NewSessionService(s *server.Server, userService *UserService) *SessionService {
	return &SessionService{s, userService}
}

func (service *SessionService) Login(w http.ResponseWriter, r *http.Request, dto *sessiondto.LoginDTO) (*db.User, error) {
	// Find in DB
	user, err := service.userService.GetUserByEmail(dto.Email)
	if err != nil {
		return nil, fmt.Errorf("incorrect email / password")
	}

	// Check hash
	match := utils.CheckPasswordHash(dto.Password, user.Password)
	if !match {
		return nil, fmt.Errorf("incorrect email / password")
	}

	// Save in Session
	session, err := service.s.SessionStore.Get(r, "user-session")
	if err != nil {
		return nil, fmt.Errorf("unable to log in user")
	}
	session.Values["userId"] = user.ID.String()
	err = session.Save(r, w)
	if err != nil {
		return nil, fmt.Errorf("unable to log in user")
	}

	return user, nil
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

	return user, nil
}

func (service *SessionService) Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := service.s.SessionStore.Get(r, "user-session")
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	delete(session.Values, "userId")
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (service *SessionService) GetLoggedInUser(w http.ResponseWriter, r *http.Request) (*db.User, error) {
	session, err := service.s.SessionStore.Get(r, "user-session")
	if err != nil {
		return nil, err
	}

	val, ok := session.Values["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	user, err := service.userService.GetUser(val)
	if err != nil {
		return nil, err
	}

	return user, nil
}

package routes

import (
	"net/http"

	"coauth/pkg/config/di"
	"coauth/pkg/config/server"
)

func RegisterRoutes(s *server.Server, di *di.DIContainer) {

	// Session Authentication
	sessionAuth := s.Router.PathPrefix("/session").Subrouter().StrictSlash(false)
	sessionAuth.HandleFunc("/login", di.SessionHandler.HandleSessionLogin).Methods(http.MethodPost)
	sessionAuth.HandleFunc("/signup", di.SessionHandler.HandleSessionSignup).Methods(http.MethodPost)
	sessionAuth.HandleFunc("/logout", di.SessionHandler.HandleSessionLogout).Methods(http.MethodPost)
	sessionAuth.HandleFunc("/user", di.SessionHandler.HandleSessionUser).Methods(http.MethodGet)

	// Users
	user := s.Router.PathPrefix("/users").Subrouter().StrictSlash(false)
	user.HandleFunc("/{id}", di.UserHandler.HandleUserGet).Methods(http.MethodGet)
	user.HandleFunc("", di.UserHandler.HandleUsersGet).Methods(http.MethodGet)
	user.HandleFunc("", di.UserHandler.HandleUserCreate).Methods(http.MethodPost)
	// user.HandleFunc("/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	user.HandleFunc("/{id}", di.UserHandler.HandleUserDelete).Methods(http.MethodDelete)
}

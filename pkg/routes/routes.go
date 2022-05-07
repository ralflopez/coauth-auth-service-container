package routes

import (
	"net/http"

	"coauth/pkg/config/di"
	"coauth/pkg/config/server"
)

func RegisterRoutes(s *server.Server, di *di.DIContainer) {

	// Users
	user := s.Router.PathPrefix("/users").Subrouter().StrictSlash(false)
	// user.HandleFunc("/{id}", ).Methods(http.MethodGet)
	// user.HandleFunc("", di.UserHandler.GetUsers).Methods(http.MethodGet)
	user.HandleFunc("", di.UserHandler.CreateUser).Methods(http.MethodPost)
	// user.HandleFunc("/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	// user.HandleFunc("/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
}
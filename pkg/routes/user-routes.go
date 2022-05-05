package routes

import (
	"net/http"

	"coauth/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	user := r.PathPrefix("/users").Subrouter().StrictSlash(false)
 
	user.HandleFunc("/{id}", controllers.GetUser).Methods(http.MethodGet)
	user.HandleFunc("", controllers.GetUsers).Methods(http.MethodGet)
	user.HandleFunc("", controllers.CreateUser).Methods(http.MethodPost)
	user.HandleFunc("/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	user.HandleFunc("/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
}
/* Dependency Injection Container */
package di

import (
	"coauth/pkg/config/server"
	"coauth/pkg/handlers"
	"coauth/pkg/repository"
	"coauth/pkg/services"
	"context"
)

type DIContainer struct {
	UserHandler *handlers.UserHandler
	userService *services.UserService
	userRepository *repository.UserRepository
}

func NewDIContainer(s *server.Server) *DIContainer {
	// user
	userRepo := repository.NewUserRepository(s, context.Background(), s.DB)
	userService := services.NewUserService(s, userRepo)
	userHandler := handlers.NewUserHandler(s, userService)
	
	return &DIContainer{
		userHandler,
		userService,
		userRepo,
	}
}
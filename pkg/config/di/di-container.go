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

	SessionHandler *handlers.SessionHandler
	SessionService *services.SessionService

	JwtHandler *handlers.JwtHandler
	jwtService *services.JwtService
}

func NewDIContainer(s *server.Server) *DIContainer {
	// user
	userRepo := repository.NewUserRepository(s, context.Background(), s.DB)
	userService := services.NewUserService(s, userRepo)
	userHandler := handlers.NewUserHandler(s, userService)

	// session
	sessionService := services.NewSessionService(s, userService)
	sessionHandler := handlers.NewSessionHandler(s, sessionService)

	// jwt
	jwtService := services.NewJwtService(s, userService)
	jwtHandler := handlers.NewJwtHandler(s, jwtService)
	
	return &DIContainer{
		userHandler,
		userService,
		userRepo,
		sessionHandler,
		sessionService,
		jwtHandler,
		jwtService,
	}
}
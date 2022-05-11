package handlers

import (
	"coauth/pkg/config/server"
	"coauth/pkg/dtos/userdto"
	"coauth/pkg/exceptions"
	"coauth/pkg/services"
	"net/http"
)

type JwtHandler struct {
	s *server.Server
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
		exceptions.ThrowInternalServerError(w, "Token generation error")
		return
	}

	handler.s.Logger.Printf("Request Body: %v\n", createUserDTO)
	
	user, err := handler.jwtService.Signup(createUserDTO)
	if err != nil {
		handler.s.Logger.Printf("User Creation Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, "Token generation error")
		return
	}

	jwtResponse, err := handler.jwtService.GenerateTokens(user.ID.String())
	if err != nil {
		handler.s.Logger.Printf("Token Generator Error: %v\n", err.Error())
		exceptions.ThrowInternalServerError(w, "Token generation error")
		return
	}

	handler.s.Respond(w, jwtResponse, http.StatusOK)
}

// func (handler *JwtHandler) HandleJwtUser(w http.ResponseWriter, r *http.Request) {
// 	// Get from jwt token
// 	reqToken := r.Header.Get("Authorization")
// 	splitToken := strings.Split(reqToken, "Bearer ")
// 	jwtFromHeader := splitToken[1]

// 	handler.s.Logger.Printf("Received jwt: %v\n", jwtFromHeader)

// 	// Parse
// 	token, err := jwt.ParseWithClaims(jwtFromHeader, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		handler.s.Logger.Printf("Parsing error: %v\n", err.Error())
// 		exceptions.ThrowBadRequestException(w, "jwt tampered")
// 		return
// 	}

// 	// Extract claims
// 	claims, ok := token.Claims.(*CustomClaims)
// 	if !ok {
// 		handler.s.Logger.Printf("Error extracting claims")
// 		exceptions.ThrowInternalServerError(w, "error extracting claims")
// 		return
// 	}

// 	// Check
// 	if claims.ExpiresAt < time.Now().UTC().Unix() {
// 		handler.s.Logger.Printf("Jwt expired")
// 		exceptions.ThrowInternalServerError(w, "jwt expired")
// 		return
// 	}

// 	// Extract value
// 	userId := claims.UserId
// 	handler.s.Logger.Printf("UserId: %v\n", userId)
	
// 	handler.s.Respond(w, userId, http.StatusOK)
// }
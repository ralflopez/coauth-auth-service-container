package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/jwtdto"
	"coauth/pkg/dtos/userdto"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	s *server.Server
	userService *UserService
}

func NewJwtService(s *server.Server, userService *UserService) *JwtService {
	return &JwtService{s, userService}
}

func (service *JwtService) GenerateTokens(userId string) (*jwtdto.JwtResponse, error){
	accessTokenClaims := jwtdto.CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer: "website_name",
		},
	}

	refreshTokenClaims := jwtdto.CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
			Issuer: "website_name",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	signedAccessToken, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	jwtResponse := &jwtdto.JwtResponse{
		RefreshToken: signedRefreshToken,
		AccessToken: signedAccessToken,
	}

	return jwtResponse, nil
}

func (service *JwtService) Signup(dto *userdto.CreateUserDTO) (*db.User, error) {
	user, err := service.userService.CreateUser(dto)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *JwtService) GetLoggedInUser(r *http.Request) (*db.User, error) {
	// Get from jwt token
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	jwtFromHeader := splitToken[1]

	service.s.Logger.Printf("Received jwt: %v\n", jwtFromHeader)

	// Parse
	token, err := jwt.ParseWithClaims(jwtFromHeader, &jwtdto.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(*jwtdto.CustomClaims)
	if !ok {
		return nil, err
	}

	// Check
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, err
	}

	// Extract value
	userId := claims.UserId
	service.s.Logger.Printf("UserId: %v\n", userId)

	// Fetch from database
	user, err := service.userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
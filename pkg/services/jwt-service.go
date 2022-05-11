package services

import (
	"coauth/pkg/config/server"
	"coauth/pkg/db"
	"coauth/pkg/dtos/jwtdto"
	"coauth/pkg/dtos/userdto"
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
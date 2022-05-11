package jwtdto

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserId string `json:"username"`
	jwt.StandardClaims
}

type JwtResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
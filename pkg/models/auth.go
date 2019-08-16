package models

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type Login struct {
	Username string
	Password string
}

type Token struct {
	JWTToken string
}

// Create a struct that will be encoded to a JWT.
type Claims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

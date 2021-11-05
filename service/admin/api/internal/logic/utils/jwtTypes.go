package utils

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtInfo struct {
	Issuer string
	Secret string
}

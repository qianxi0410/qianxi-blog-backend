package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(info JwtInfo, username string, expireDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    info.Issuer,
		},
	})

	return token.SignedString([]byte(info.Secret))
}

func ParseToken(info JwtInfo, tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return info.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	// check
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

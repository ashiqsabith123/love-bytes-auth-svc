package utils

import (
	"time"

	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 300).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

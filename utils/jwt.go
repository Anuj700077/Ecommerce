package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("my_secret_key")

func GenerateToken(email string, role string) (string, error) {

	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SecretKey)
}

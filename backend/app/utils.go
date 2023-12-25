package app

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type key struct {
	Key []byte
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["userId"] = userId

	key := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

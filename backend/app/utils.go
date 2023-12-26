package app

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type key struct {
	Key []byte
}

func GenerateJWT(userId uint, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 10).Format(time.RFC3339Nano)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["sub"] = username
	claims["iat"] = time.Now().Format(time.RFC3339Nano)

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

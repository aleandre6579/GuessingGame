package app

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["userId"] = userId

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

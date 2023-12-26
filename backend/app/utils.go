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
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	claims["sub"] = userId
	claims["iat"] = time.Now().Unix()
	claims["alg"] = jwt.SigningMethodHS256.Name

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

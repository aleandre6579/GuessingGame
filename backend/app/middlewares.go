package app

import (
	"context"
	"fmt"
	"guessing-game/db/models/user"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type BaseContext struct {
	app *App
}

func (baseCtx *BaseContext) SetContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "db_client", baseCtx.app.DBClient)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Header["Token"] != nil {
			fmt.Println("No token in request headers")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, _ := jwt.Parse(r.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Check expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				w.WriteHeader(http.StatusUnauthorized)
			}

			// Check username exists
			dbClient := ctx.Value("db_client").(*gorm.DB)
			userId := claims["userId"].(uint)
			_, err := user.GetUser(dbClient, userId)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

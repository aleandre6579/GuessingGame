package app

import (
	"context"
	"fmt"
	"guessing-game/db/models/user"
	"net/http"
	"os"
	"strings"
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
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Header["Authorization"] == nil {
			fmt.Println("No token in request headers")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokenString := strings.Split(r.Header["Authorization"][0], "Bearer ")[1]
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Check expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				fmt.Println("Token unauthorized")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Check username exists
			dbClient := ctx.Value("db_client").(*gorm.DB)
			userId := uint(claims["sub"].(float64))

			_, err := user.GetUser(dbClient, userId)
			if err != nil {
				fmt.Println("Token userId does not exist: ", userId)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

		} else {
			fmt.Println("Token not valid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

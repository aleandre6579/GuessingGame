package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func loadRoutes(app *App) *chi.Mux {
	router := chi.NewRouter()
	context := BaseContext{app: app}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		Debug:          true,
	})
	router.Use(c.Handler)

	router.Use(middleware.Logger)
	router.Use(context.SetContextMiddleware)

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		Login(app.DBClient, w, r)
		w.WriteHeader(http.StatusOK)
	})

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		Register(app.DBClient, w, r)
		w.WriteHeader(http.StatusOK)
	})

	router.Group(func(router chi.Router) {
		router.Use(AuthMiddleware)

		router.Post("/guess", func(w http.ResponseWriter, r *http.Request) {
			Guess(app.DBClient, w, r)
			w.WriteHeader(http.StatusOK)
		})
	})

	return router
}

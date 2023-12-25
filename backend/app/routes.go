package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func loadRoutes(app *App) *chi.Mux {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		Debug:          false,
	})
	router.Use(c.Handler)

	router.Use(middleware.Logger)

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		Login(w, r)
		w.WriteHeader(http.StatusOK)
	})

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		Register(w, r)
		w.WriteHeader(http.StatusOK)
	})

	router.Post("/guess", func(w http.ResponseWriter, r *http.Request) {
		Guess(w, r)
		w.WriteHeader(http.StatusOK)
	})

	return router
}

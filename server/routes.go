package server

import (
	"github.com/ODawah/url-shortener/handlers"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routes() http.Handler {
	r := GetRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handlers.Healthcheck())

	r.Post("/users", handlers.CreateUser())
	r.Get("/users/{id}", handlers.FindUser())

	r.Post("/urls", handlers.EncodeURL())
	return r
}

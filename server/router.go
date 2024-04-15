package server

import "github.com/go-chi/chi/v5"

func GetRouter() *chi.Mux {
	return chi.NewRouter()
}

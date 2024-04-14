package main

import (
	"github.com/ODawah/url-shortener/persistence"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	err := persistence.InitializeSQL()
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}

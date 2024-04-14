package main

import (
	"github.com/ODawah/url-shortener/persistence"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	err := persistence.IntializeSQL()
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

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
	err := persistence.InitializeSQL()
	if err != nil {
		log.Println(err)
	}
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				err = persistence.InitializeSQL()
				if err != nil {
					log.Println(err)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, World!"))
	})
	log.Println("up and running")
	log.Fatal(http.ListenAndServe(":8080", r))

}

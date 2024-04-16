package main

import (
	"github.com/ODawah/url-shortener/persistence"
	"github.com/ODawah/url-shortener/server"
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
	r := server.Routes()
	http.ListenAndServe(":8080", r)
}

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
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				err := persistence.InitializeSQL(&quit)
				if err != nil {
					log.Println(err)
				}

			case <-quit:
				log.Println("RUNNING")
				ticker.Stop()
				return
			}
		}
	}()
	r := server.Routes()
	http.ListenAndServe(":8080", r)
}

package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Construct MySQL connection string
	dbURI := os.Getenv("DB_URI")
	// Open a connection to the database
	time.Sleep(30 * time.Second)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.Run()
}

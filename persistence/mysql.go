package persistence

import (
	"database/sql"
	"os"
)

var (
	DB *sql.DB
)

func IntializeSQL() error {
	dbURI := os.Getenv("DB_URI")
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}
	defer db.Close()
	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return err
	}
	return err
}

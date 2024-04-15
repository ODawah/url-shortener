package persistence

import (
	"database/sql"
	"os"
)

var (
	DB *sql.DB
)

func InitializeSQL() error {
	dbURI := os.Getenv("DB_URI")
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}
	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = db
	return err
}

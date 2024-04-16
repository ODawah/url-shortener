package persistence

import (
	"database/sql"
	"fmt"
	"github.com/ODawah/url-shortener/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func InitializeSQL() error {
	dbURI := fmt.Sprintf("%s?parseTime=true", os.Getenv("DB_URI"))
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}
	// Ping the database to verify the connection
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		DSN:  dbURI,
	}), &gorm.Config{})
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = gormDB
	err = DB.AutoMigrate(&models.User{}, &models.URL{})
	if err != nil {
		return err
	}
	return err
}

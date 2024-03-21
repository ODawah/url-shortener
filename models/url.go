package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	Original string
	UserID   int
	Short    string
}

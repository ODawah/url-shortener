package models

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) CreateUser(db *gorm.DB) error {
	err := db.Create(&u).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func FindUserByID(id int, db *gorm.DB) (User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

package models

import (
	"fmt"
	"github.com/ODawah/url-shortener/encoders"
	"gorm.io/gorm"
	"time"
)

type URL struct {
	gorm.Model
	UserID   int  `json:"userID"`
	ID       uint `gorm:"primaryKey"`
	Original string
	Short    string `gorm:"UNIQUE"`
	User     User   `gorm:"foreignkey:UserID" json:"-"`
}

func (url *URL) CreateURL(db *gorm.DB) error {
	err := url.EncodeURL()
	if err != nil {
		return err
	}
	err = db.Create(&url).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalURL(shortURL string, db *gorm.DB) (URL, error) {
	var url URL
	if err := db.Where("short = ?", shortURL).First(&url).Error; err != nil {
		return url, err
	}
	return url, nil
}

func (url *URL) EncodeURL() error {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	md5encode, err := encoders.MD5Encode(fmt.Sprintf("%s%s", url.Original, currentTimeString))
	if err != nil {
		return err
	}
	base62encode, err := encoders.Base62Encode(md5encode)
	if err != nil {
		return err
	}
	result := base62encode[:7]
	url.Short = result
	return nil
}

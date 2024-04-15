package models

import (
	"github.com/ODawah/url-shortener/encoders"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Original string
	UserID   int
	Short    string
}

func (url *URL) CreateURL(db *gorm.DB) error {
	err := db.Create(&url).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalURL(shortURL string, db *gorm.DB) (URL, error) {
	var url URL
	if err := db.First(&url, shortURL).Error; err != nil {
		return url, err
	}
	return url, nil
}

func (url *URL) EncodeURL() (string, error) {
	md5encode, err := encoders.MD5Encode(url.Original)
	if err != nil {
		return "", err
	}
	base62encode, err := encoders.Base62Encode(md5encode)
	if err != nil {
		return "", err
	}
	result := base62encode[:8]
	return result, nil
}

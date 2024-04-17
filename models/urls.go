package models

import (
	"github.com/ODawah/url-shortener/encoders"
	"gorm.io/gorm"
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
	if err := db.First(&url, shortURL).Error; err != nil {
		return url, err
	}
	return url, nil
}

func (url *URL) EncodeURL() error {
	md5encode, err := encoders.MD5Encode(url.Original)
	if err != nil {
		return err
	}
	base62encode, err := encoders.Base62Encode(md5encode)
	if err != nil {
		return err
	}
	result := base62encode[:8]
	url.Short = result
	return nil
}

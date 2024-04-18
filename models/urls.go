package models

import (
	"errors"
	"fmt"
	"github.com/ODawah/url-shortener/encoders"
	"gorm.io/gorm"
	"net/url"
	"strings"
	"time"
)

type URL struct {
	gorm.Model
	UserID   int  `json:"userID"`
	ID       uint `gorm:"primaryKey"`
	Original string
	Short    string `gorm:"UNIQUE"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
}

func (u *URL) CreateURL(db *gorm.DB) error {
	err := u.ValidateOriginalURL()
	if err != nil {
		return err
	}
	for {
		err = u.EncodeURL()
		if err != nil {
			return err
		}
		err = db.Create(&u).Error
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				continue
			}
			return err
		}
		break
	}
	return nil
}

func GetOriginalURL(shortURL string, db *gorm.DB) (URL, error) {
	var u URL
	if err := db.Where("short = ?", shortURL).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (u *URL) EncodeURL() error {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	md5encode, err := encoders.MD5Encode(fmt.Sprintf("%s%s", u.Original, currentTimeString))
	if err != nil {
		return err
	}
	base62encode, err := encoders.Base62Encode(md5encode)
	if err != nil {
		return err
	}
	result := base62encode[:7]
	u.Short = result
	return nil
}

func (u *URL) ValidateOriginalURL() error {
	_, err := url.Parse(u.Original)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(u.Original, "https://") && !strings.HasPrefix(u.Original, "http://") {
		return errors.New(`the url must start with "http://" or "https://" `)
	}

	return err
}

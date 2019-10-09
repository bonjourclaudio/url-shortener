package models

import (
	"errors"
	"github.com/claudioontheweb/url-shortener/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
	"github.com/lithammer/shortuuid"
)

type UrlShorten struct {
	gorm.Model

	OriginalUrl string `json:"originalUrl"`
	UrlCode string `json:"urlCode" gorm:"unique"`
	ShortUrl string `json:"shortUrl"`
}

func GetOriginalUrl(db *gorm.DB, code string) (string, error) {
	var urlShorten UrlShorten

	if db.Where("url_code = ?", code).Find(&urlShorten).RecordNotFound() {
		return "", errors.New("Record not found")
	} else {
		return urlShorten.OriginalUrl, nil
	}

}

func CreateShortUrl(db *gorm.DB, urlShorten UrlShorten) (string, error) {

	config.GetConfig()

	_, err := url.ParseRequestURI(urlShorten.OriginalUrl)
	if err != nil {
		return "", errors.New("Invalid URL")
	}

	urlShorten.UrlCode = shortuuid.New()

	urlShorten.ShortUrl = viper.GetString("BASE_URL") + "/" + urlShorten.UrlCode

	db.Create(&urlShorten)

	return urlShorten.ShortUrl,nil
}
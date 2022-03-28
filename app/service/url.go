package service

import (
	"URL-Shortener/app/model"
	"URL-Shortener/app/persistence"
	"time"
)

var UrlFields = []string{"id", "origin_URL", "expired_Date", "url_ID"}

func AddUrl(origin_URL string, expired_Date time.Time, url_ID string) (*model.Url, error){
	url := &model.Url{
		Origin_URL:   origin_URL,
		Expired_Date: expired_Date,
		Url_ID: url_ID,
	}
	err := persistence.SqlSession.Model(&model.Url{}).Create(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}

func QueryUrl(url_id string) (*model.Url, error) {
	url := &model.Url{}
	err := persistence.SqlSession.Select(UrlFields).Where("url_ID=?", url_id).First(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}



package service

import (
	"URL-Shortener/app/model"
	"URL-Shortener/app/persistence"
	"time"
)

var UrlFields = []string{"id", "origin_URL", "expired_Date"}

func AddUrl(origin_URL string, expired_Date time.Time) (*model.Url, error){
	url := &model.Url{
		Origin_URL:   origin_URL,
		Expired_Date: expired_Date,
	}
	err := persistence.SqlSession.Model(&model.Url{}).Create(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}

func QueryUrl(id int64) (*model.Url, error) {
	url := &model.Url{}
	err := persistence.SqlSession.Select(UrlFields).Where("id=?", id).First(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}



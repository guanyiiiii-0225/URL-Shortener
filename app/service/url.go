package service

import (
	"example.com/url/app/model"
	"example.com/url/app/persistence"
	"time"
)

var UrlFields = []string{"id", "originURL", "expiredDate", "shortenURL"}

func AddUrl(originURL string) (*model.Url, error){
	url := &model.Url{
		OriginURL:   originURL,
		ExpiredDate: time.Now(),
		ShortenURL:  "12345",
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



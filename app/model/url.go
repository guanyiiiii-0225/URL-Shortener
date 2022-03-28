package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ID      int64     `gorm:"primary_key;auto_increment" json:"id"`
	Origin_URL    string    `gorm:"not null" json:"origin_URL"`
	Expired_Date time.Time `gorm:"not null" json:"expired_Date"`
	Url_ID string `gorm:"not null;unique" json:"url_ID"`
}

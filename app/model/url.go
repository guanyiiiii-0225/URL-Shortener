package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ID      int64     `gorm:"primary_key;auto_increment" json:"id"`
	Origin_URL    string    `gorm:"not null" json:"origin_URL"`
	Expired_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"expired_Date"`
}

// CREATE TABLE URLS	
// (
// 	ID serial PRIMARY KEY,
// 	OriginalURL varchar not null,
// 	ExpiredDate TIMESTAMP
// )

package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ID      int64     `gorm:"primary_key;auto_increment" json:"id"`
	OriginURL    string    `gorm:"not null" json:"originURL"`
	ExpiredDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"expiredDate"`
	ShortenURL string    `gorm:"not null" json:"shortenURL"`
}

// CREATE TABLE URL
// (
// 	ID serial PRIMARY KEY,
// 	OriginalURL varchar not null,
// 	ShortenURL varchar not null unique,
// 	ExpiredDate TIMESTAMP
// )

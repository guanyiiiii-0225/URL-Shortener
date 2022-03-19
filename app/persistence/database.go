package persistence

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	SqlSession *gorm.DB
)

func Initialize(host string, port string, username string, password string, dbName string) (*gorm.DB, error) {
	dbConfig := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName
	var err error
	SqlSession, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return SqlSession, err
}

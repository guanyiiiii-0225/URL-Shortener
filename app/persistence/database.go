package persistence

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var SqlSession *gorm.DB

func Initialize(host string, port string, username string, password string, dbName string) (*gorm.DB, error) {
	dbConfig := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName
	var err error
	SqlSession, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	sqlDB, err := SqlSession.DB() // Get generic database object sql.DB to use its functions	
	sqlDB.SetMaxIdleConns(10) // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.	
	sqlDB.SetMaxOpenConns(100) // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Hour)// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	return SqlSession, err
}

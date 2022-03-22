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

	// Get generic database object sql.DB to use its functions
	sqlDB, err := SqlSession.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return SqlSession, err
}

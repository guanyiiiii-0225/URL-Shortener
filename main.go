package main

import (
	"github.com/joho/godotenv"
	"fmt"
	_"log"
    "os"
	"URL-Shortener/app/model"
	"URL-Shortener/app/persistence"
	"URL-Shortener/app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	envErr := godotenv.Load()
	handleErr(envErr)

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")
	db, ormErr := persistence.Initialize(host, port, username, password, dbName)
	handleErr(ormErr)
	// dbConfig := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName
	// db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	// handleErr(err)

	// db.Migrator().HasTable(&User{})
	fmt.Printf("exists or not: %v\n", db.Migrator().HasTable(&model.Url{}))
	migrateErr := db.AutoMigrate(&model.Url{})
	handleErr(migrateErr)

	router := gin.Default()
	router.GET("/test", test)
	config.RouteUrls(router)
	router.Run("localhost:8080")
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func test(c *gin.Context) {
	var message = "Hello World"
	c.IndentedJSON(http.StatusOK, message)
}
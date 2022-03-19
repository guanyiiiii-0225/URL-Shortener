package main

import (
	"github.com/joho/godotenv"
	_"fmt"
	_"log"
    "os"
	"example.com/url/app/model"
	"example.com/url/app/persistence"
	"example.com/url/app/config"
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
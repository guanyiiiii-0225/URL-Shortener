package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"URL-Shortener/app/service"
	"fmt"
	"time"
	_"log"
	"os"
	"math/rand"
)

type UrlController struct{}

func Url_Controller() UrlController {
	return UrlController{}
}

type AddUrlInput struct {
	Origin_URL    string `json:"url" binding:"required" example:"origin_URL"`
	Expired_Date  time.Time `json:"expireAt" binding:"required" example:"2021-02-08T09:20:41Z"`
}

// screate random string for url_id
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
func randStr(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


// AddUrl @Summary
// @Tags Url
// @version 1.0
// @produce application/json
// @param data body AddUrlInput true "AddUrlInput"
// @Success 200 string successful return data
// @Router /api/v1/urls [post] 
func (u UrlController) AddUrl(c *gin.Context) {
	var form AddUrlInput
	bindErr := c.BindJSON(&form)
	
	if bindErr == nil { // catch bind json error
		rand.Seed(time.Now().UnixNano())
		var random_urlID = randStr(10)
		url, err := service.AddUrl(form.Origin_URL, form.Expired_Date, random_urlID)
		if err == nil { // catch AddUrl error
			shortUrl := fmt.Sprintf("%s%s%s%s", "http://", os.Getenv("DOMAIN"), "/", url.Url_ID)
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"url_id": url.Url_ID,
				"shortUrl": shortUrl,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":   err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  bindErr.Error(),
			"msg": "Input field incorrect.",
		})
	}
}

// QueryUrl @Summary
// @Tags Url
// @version 1.0
// @produce application/json
// @param url_id path string true "url_id"
// @Success 301 string successful redirect to original URL
// @Router /{url_id} [get]
func (u UrlController) QueryUrl(c *gin.Context) {
	url_id := c.Params.ByName("url_id")

	url, err := service.QueryUrl(url_id) 
	if err != nil { // check shorten URL is exist or not
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  "Non-existent shorten URL.",  //err.Error(),
		})
	} else {
		today := time.Now()
		if today.Before(url.Expired_Date) == true { // compare today and the expired date
			c.Redirect(http.StatusMovedPermanently, url.Origin_URL)
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "failed",
				"msg": "URL is expired.",
			})
		}
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"URL-Shortener/app/service"
	"strconv"
	"fmt"
	_"github.com/pkg/browser"
	"time"
	_"log"
)

type UrlController struct{}

func Url_Controller() UrlController {
	return UrlController{}
}

type AddUrlInput struct {
	Origin_URL    string `json:"url" binding:"required" example:"origin_URL"`
	Expired_Date  time.Time `json:"expireAt" binding:"required" example:"2021-02-08T09:20:41Z"`
}


// AddUrl @Summary
// @Tags Url
// @version 1.0
// @produce application/json
// @param data body AddUrlInput true "AddUrlInput"
// @Success 200 string successful return data
// @Router /url [post] 
func (u UrlController) AddUrl(c *gin.Context) {
	var form AddUrlInput
	bindErr := c.BindJSON(&form)
	// debug
	fmt.Printf("form: %v\n", form)
	fmt.Printf("c: %v\n", &c)
	
	// catch bind json error
	if bindErr == nil {
		// call AddUrl function from service
		url, err := service.AddUrl(form.Origin_URL, form.Expired_Date)
		// catch AddUrl error
		if err == nil {
			shortUrl := fmt.Sprintf("%s%d", "http://localhost:8080/url/", url.ID)
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"id":    url.ID,
				"shortUrl": shortUrl,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"data":    nil,
				"error":   err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"data":    nil,
			"error":   bindErr.Error(),
		})
	}
}

// QueryUrl @Summary
// @Tags Url
// @version 1.0
// @produce application/json
// @param url_id path int true "url_id"
// @Success 200 string successful return data
// @Router /url/{url_id} [get]
func (u UrlController) QueryUrl(c *gin.Context) {
	id := c.Params.ByName("url_id")
	// debug
	// fmt.Printf("id: %v\n", id)
	// fmt.Printf("%v\n", c.Params)
	
	urlId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":   err.Error(),
		})
	}
	url, err := service.QueryUrl(urlId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"error":   err.Error(),
		})
	} else {
		today := time.Now()
		if today.Before(url.Expired_Date) == true {
			// http.Handle("/", http.RedirectHandler(url.Origin_URL, 200))
			// browser.OpenURL(url.Origin_URL)
			c.Redirect(http.StatusMovedPermanently, url.Origin_URL)
			// c.JSON(http.StatusOK, gin.H{
			// 	"status": "success",
			// 	"id": url.ID,
			// 	"origin_URL": url.Origin_URL,
			// 	"expired_Date": url.Expired_Date,
			// })
			// c.JSON(http.StatusOK, gin.H{})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"msg": "URL is expired.",
			})
		}
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"example.com/url/app/service"
	"strconv"
	"fmt"
	"github.com/pkg/browser"
)

type UrlController struct{}

func AddUrlController() UrlController {
	return UrlController{}
}

func QueryUrlController() UrlController {
	return UrlController{}
}

type AddUrlInput struct {
	Origin_URL    string `json:"origin_URL" binding:"required" example:"origin_URL"`
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
	// fmt.Printf("form: %v\n", form)
	// fmt.Printf("c: %v\n", &c)
	
	// catch bind json error
	if bindErr == nil {
		// call AddUrl function from service
		url, err := service.AddUrl(form.Origin_URL)
		// catch AddUrl error
		if err == nil {
			shortUrl := fmt.Sprintf("%s%d", "http://localhost:8080/url/", url.ID)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"id":    url.ID,
				"shortUrl": shortUrl,
				"error":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"data":    nil,
				"error":   err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success": false,
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
	fmt.Printf("id: %v\n", id)
	// fmt.Printf("c: %v\n", &c)
	fmt.Printf("%v\n", c.Params)
	urlId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	url, err := service.QueryUrl(urlId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	} else {
		browser.OpenURL(url.Origin_URL)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"id": url.ID,
			"origin_URL": url.Origin_URL,
			"expired_Date": url.Expired_Date,
			"error":   nil,
		})
	}
}

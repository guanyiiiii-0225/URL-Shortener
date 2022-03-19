package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"example.com/url/app/service"
	"strconv"
	"fmt"
)

type UrlController struct{}

func AddUrlController() UrlController {
	return UrlController{}
}

func QueryUrlController() UrlController {
	return UrlController{}
}

type AddUrlInput struct {
	OriginURL    string `json:"originURL" binding:"required" example:"originURL"`
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
		url, err := service.AddUrl(form.OriginURL)
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
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    url,
			"error":   nil,
		})
	}
}

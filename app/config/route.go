package config

import (
	"github.com/gin-gonic/gin"
	"URL-Shortener/app/controller"
)

func RouteUrls(r *gin.Engine) {
	urls := r.Group("")
	{
		urls.POST("/api/v1/urls", controller.Url_Controller().AddUrl)
		urls.GET("/:url_id", controller.Url_Controller().QueryUrl)
	}
}



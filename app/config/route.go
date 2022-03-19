package config

import (
	"github.com/gin-gonic/gin"
	"example.com/url/app/controller"
)

func RouteUrls(r *gin.Engine) {
	urls := r.Group("/url")
	{
		urls.POST("", controller.AddUrlController().AddUrl)
		urls.GET("/:url_id", controller.QueryUrlController().QueryUrl)
	}
}



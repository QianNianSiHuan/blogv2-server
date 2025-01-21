package router

import (
	"blogv2/api"
	"github.com/gin-gonic/gin"
)

func SearchRouter(r *gin.RouterGroup) {
	app := api.App.SearchApi
	r.GET("/search/article", app.ArticleSearchView)
	r.GET("/search/text", app.TextSearchView)
}

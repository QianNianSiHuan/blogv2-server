package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func SearchRouter(r *gin.RouterGroup) {
	app := api.App.SearchApi
	r.GET("/search/article", app.ArticleSearchView)
	r.POST("/search/index", middleware.AdminMiddleware, app.ArticleSearchIndexView)
	r.GET("/search/text", app.TextSearchView)
	r.GET("/search/tags", app.TagAggView)
}

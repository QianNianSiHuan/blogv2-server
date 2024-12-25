package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddleware, app.ArticleCreateView)
}

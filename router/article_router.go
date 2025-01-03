package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddleware, app.ArticleCreateView)
	r.PUT("article", middleware.AuthMiddleware, app.ArticleUpdateView)
	r.GET("article", app.ArticleListView)
	r.GET("article/:id", app.ArticleDetailView)
	r.POST("article/examine", middleware.AdminMiddleware, app.ArticleExamineView)
	r.GET("article/digg/:id", middleware.AuthMiddleware, app.ArticleDiggView)
	r.POST("article/collect", middleware.AuthMiddleware, app.ArticleCollectView)
	r.POST("article/look", app.ArticleLookView)
}

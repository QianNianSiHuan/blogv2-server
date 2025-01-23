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

	r.DELETE("article/:id", middleware.AuthMiddleware, app.ArticleRemoveUserView)
	r.DELETE("article", middleware.AdminMiddleware, app.ArticleRemoveView)

	r.POST("article/examine", middleware.AdminMiddleware, app.ArticleExamineView)

	r.GET("article/digg/:id", middleware.AuthMiddleware, app.ArticleDiggView)
	r.POST("article/collect", middleware.AuthMiddleware, app.ArticleCollectView)
	r.DELETE("article/collect", middleware.AuthMiddleware, app.ArticleCollectPatchRemoveView)

	r.POST("article/history", app.ArticleLookView)
	r.GET("article/history", middleware.AuthMiddleware, app.ArticleLookListView)
	r.DELETE("article/history", middleware.AuthMiddleware, app.ArticleLookRemoveView)

	r.POST("category", middleware.AuthMiddleware, app.CategoryCreateView)
	r.GET("category", app.CategoryListView)
	r.DELETE("category", middleware.AuthMiddleware, app.CategoryRemoveView)

	r.POST("collect", middleware.AuthMiddleware, app.CollectCreateView)
	r.GET("collect", app.CollectListView)
	r.DELETE("collect", middleware.AuthMiddleware, app.CollectRemoveView)

	r.GET("category/options", middleware.AuthMiddleware, app.CategoryOptionsView)
	r.GET("article/tag/options", middleware.AuthMiddleware, app.ArticleTagOptions)
	r.GET("article/article_recommend", app.ArticleRecommendView)

}

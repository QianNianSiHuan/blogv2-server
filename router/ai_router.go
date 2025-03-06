package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func AiRouter(r *gin.RouterGroup) {
	app := api.App.AiApi
	r.POST("ai/analysis", app.AiArticleAnalysis)
	r.GET("ai/article", middleware.AuthMiddleware, app.ArticleAiView)
}

package router

import (
	"blogv2/api"
	"github.com/gin-gonic/gin"
)

func AiRouter(r *gin.RouterGroup) {
	app := api.App.AiApi
	r.POST("ai/analysis", app.AiArticleAnalysis)
}

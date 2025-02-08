package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func FeedBackRouter(r *gin.RouterGroup) {
	app := api.App.FeedBackApi
	r.POST("/feedback", middleware.AuthMiddleware, app.UserFeedBackView)
}

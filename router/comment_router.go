package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("comment", middleware.AuthMiddleware, app.CommentCreateView)
	r.GET("comment/tree/:id", app.CommentTreeView)
}

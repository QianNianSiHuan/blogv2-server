package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("comment", middleware.AuthMiddleware, app.CommentCreateView)
	r.POST("comment/examine", middleware.AdminMiddleware, app.CommentExamineView)
	r.GET("comment/tree/:id", app.CommentTreeView)
	r.GET("comment", middleware.AuthMiddleware, app.CommentListView)
	r.DELETE("comment/:id", middleware.AuthMiddleware, app.CommentRemoveView)
	r.GET("comment/digg/:id", middleware.AuthMiddleware, app.CommentDiggView)
}

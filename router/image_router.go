package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddleware, app.ImageUploadView)
	r.POST("images/transfer_deposit", middleware.AuthMiddleware, app.TransferDepositView)
	r.POST("images/qiniu", middleware.AuthMiddleware, app.QiNiuGenToken)
	r.GET("images", middleware.AuthMiddleware, app.ImageListView)
	r.DELETE("images", middleware.AuthMiddleware, app.ImageRemoveView)
}

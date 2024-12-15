package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi
	r.GET("banner", app.BannerListView)
	r.POST("banner", middleware.AuthMiddleware, app.BannerCreatView)
	r.PUT("banner/:id", middleware.AuthMiddleware, app.BannerUpdateView)
	r.DELETE("banner", middleware.AuthMiddleware, app.BannerRemoveView)
}

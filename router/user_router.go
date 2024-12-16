package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", middleware.CaptchaMiddleware, app.RegisterEmailView)
	r.POST("user/qq", app.QQLoginView)
	r.POST("user/login", middleware.CaptchaMiddleware, app.PwdLoginView)
	r.GET("user/detail", middleware.AuthMiddleware, app.UserDetailView)
	r.GET("user/base", app.UserBaseInfoView)
}

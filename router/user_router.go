package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", middleware.EmailVerifyMiddleware, app.RegisterEmailView)
	r.POST("user/admin", middleware.AdminMiddleware, app.RegisterAdminView)
	r.POST("user/qq", app.QQLoginView)
	r.POST("user/login", middleware.LoginCountByIPMiddleware, middleware.CaptchaMiddleware, app.PwdLoginView)
	r.POST("user/logout", middleware.AuthMiddleware, app.UserLogoutView)
	r.GET("user/detail", middleware.AuthMiddleware, app.UserDetailView)
	r.GET("user/login", middleware.AuthMiddleware, app.UserLoginListView)
	r.PUT("user/password", middleware.AuthMiddleware, app.UpdatePasswordView)
	r.PUT("user/password/reset", middleware.EmailVerifyMiddleware, app.ResetPasswordView)
	r.PUT("user/email/bind", middleware.EmailVerifyMiddleware, middleware.AuthMiddleware, app.BindEmailView)
	r.PUT("user", middleware.AuthMiddleware, app.UserInfoUpdateView)
	r.PUT("user/admin", middleware.AdminMiddleware, app.AdminUserInfoUpdateView)
	r.GET("user/base", app.UserBaseInfoView)
	r.GET("user/user_list", middleware.AuthMiddleware, middleware.AdminMiddleware, app.UserListView)
	r.DELETE("user/user_list", middleware.AdminMiddleware, app.UserRemoveView)
	r.POST("user/article/top", middleware.AuthMiddleware, app.UserArticleTopView)
}

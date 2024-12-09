package router

import (
	"blogv2/api"
	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	app := api.App.LoginApi
	r.POST("login", app.Login)
	r.POST("registered", app.Register)
}

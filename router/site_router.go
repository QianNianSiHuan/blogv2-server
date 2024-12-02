package router

import (
	"blogv2.0/api"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteAPi
	r.GET("site", app.SiteInfoView)
}

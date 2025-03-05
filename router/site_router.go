package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteAPi
	r.GET("site/qq_url", app.SiteInfoQQView)
	r.GET("site/:name", app.SiteInfoView)
	r.PUT("site/:name", middleware.AdminMiddleware, app.SiteUpdateView)
	r.GET("site/ai_info", app.AISiteInfoView)
}

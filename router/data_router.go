package router

import (
	"blogv2/api"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddleware, app.SumView)
	r.GET("data/article/year", middleware.AdminMiddleware, app.ArticleYearDataView)
	r.GET("data/computer", middleware.AdminMiddleware, app.ComputerDataView)
	r.GET("data/growth", middleware.AdminMiddleware, app.GrowthDataView)
}

package router

import (
	"blogv2/flags"
	"blogv2/global"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Run() {
	if !flags.FlagOptions.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/bootstrap5", "./static/bootstrap5")
	r.LoadHTMLGlob("static/view/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})
	nr := r.Group("/api")
	LoginRouter(nr)

	nr.Use(middleware.LogMiddleware)
	SiteRouter(nr)
	LogRouter(nr)
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		logrus.Fatalf("路由错误: %s ", err)
	}
	logrus.Info("路由配置成功")
	return
}

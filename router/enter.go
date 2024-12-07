package router

import (
	"blogv2/flags"
	"blogv2/global"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	if !flags.FlagOptions.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	nr := r.Group("/api")
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

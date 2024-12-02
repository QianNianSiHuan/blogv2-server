package router

import (
	"blogv2.0/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	r := gin.Default()
	nr := r.Group("/api")
	SiteRouter(nr)
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		logrus.Fatalf("路由错误: %s ", err)
	}
	logrus.Info("路由配置成功")
	return
}

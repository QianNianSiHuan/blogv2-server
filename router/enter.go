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
	r.Static("/uploads", "uploads")
	nr := r.Group("/api")
	nr.Use(middleware.LogMiddleware)
	AiRouter(nr)
	SiteRouter(nr)
	LogRouter(nr)
	UserRouter(nr)
	CaptchaRouter(nr)
	BannerRouter(nr)
	ImageRouter(nr)
	ArticleRouter(nr)
	CommentRouter(nr)
	SearchRouter(nr)
	FeedBackRouter(nr)
	DataRouter(nr)
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		logrus.Fatalf("路由错误: %s ", err)
	}
	logrus.Info("路由配置成功")
	return
}

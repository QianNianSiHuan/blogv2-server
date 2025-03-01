package core

import (
	"blogv2/global"
	//"blogv2/service/river_service/river"
	"github.com/sirupsen/logrus"
)

func InitMysqlEs() {
	//ProgressbarMsg <- "es-数据库同步初始化..."
	logrus.Println(!global.Config.River.Enable)
	if !global.Config.River.Enable {
		logrus.Info("es-Mysql同步未启用")
		return
	}
	//r, err := river.NewRiver()
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//go r.Run()
}

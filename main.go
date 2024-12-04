package main

import (
	"blogv2/core"
	"blogv2/flags"
	"blogv2/global"
	"blogv2/router"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	logrus.Infof("当前配置文件为: %s", flags.FlagOptions.File)
	global.DB = core.InitDB()
	core.InitIPDb()
	flags.Run()
	//启动web程序
	router.Run()
}

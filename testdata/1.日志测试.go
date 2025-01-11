package main

import (
	"blogv2/core"
	"blogv2/flags"
	"blogv2/global"
	"blogv2/service/log_service"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	logrus.Infof("当前配置文件为: %s", flags.FlagOptions.File)
	global.DB = core.InitDB()
	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDateHour)
	log.SetItem("文章", 11)
	log.SetTitle("同步")
	log.Save()
}

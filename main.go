package main

import (
	"blogv2/core"
	"blogv2/core/core_observer"
	"blogv2/flags"
	"blogv2/global"
	"blogv2/router"
	"blogv2/service/cron_service"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	logrus.Infof("当前配置文件为: %s", flags.FlagOptions.File)
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	global.ESClient = core.EsConnect()
	global.IP = core.InitIPDb()
	global.SensitiveWords = core.InitSensitiveWords()
	global.AhoCorasick = core.InitAhoCorasick()
	core_observer.InitObserver()
	core.InitGse() //分词
	flags.Run()
	core.InitMysqlEs()
	//定时任务
	cron_service.Cron()
	//启动web程序
	router.Run()
}

package flags

import (
	"blogv2/global"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	File    string
	DB      bool
	Version bool
	Debug   bool
}

var FlagOptions = new(Options)

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	flag.BoolVar(&FlagOptions.Debug, "debug", false, "数据库debug模式")
	flag.Parse()
}
func Run() {
	if FlagOptions.DB {
		FlagDB()
		os.Exit(0)
	}
	if FlagOptions.Debug {
		global.DB = global.DB.Debug()
		logrus.Warnf("gorm debug模式 已开启")
	}
}

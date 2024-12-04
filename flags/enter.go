package flags

import (
	"flag"
	"os"
)

type Options struct {
	File    string
	DB      bool
	Version bool
}

var FlagOptions = new(Options)

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	flag.Parse()
}
func Run() {
	if FlagOptions.DB {
		FlagDB()
		os.Exit(0)
	}
}

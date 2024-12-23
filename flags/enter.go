package flags

import (
	"blogv2/artFontFiles"
	"blogv2/flags/flag_user"
	"blogv2/global"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	File    string
	DB      bool
	Version bool
	ES      bool
	Debug   bool
	Type    string
	Sub     string
}

var FlagOptions = new(Options)

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings-dev.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	flag.BoolVar(&FlagOptions.ES, "es", false, "es索引创建")
	flag.StringVar(&FlagOptions.Type, "t", "", "类型")
	flag.StringVar(&FlagOptions.Sub, "s", "", "子类")
	flag.BoolVar(&FlagOptions.Debug, "debug", false, "数据库debug模式")
	flag.Parse()
}
func Run() {
	if FlagOptions.DB {
		FlagDB()
		artFontFiles.OutPutArtisticFont(artFontFiles.SUCCESS)
		os.Exit(0)
	}
	if FlagOptions.ES {
		EsIndex()
		artFontFiles.OutPutArtisticFont(artFontFiles.SUCCESS)
		os.Exit(0)
	}
	if FlagOptions.Debug {
		global.DB = global.DB.Debug()
		logrus.Warnf("gorm debug模式 已开启")
		artFontFiles.OutPutArtisticFont(artFontFiles.GIN_DEBUG)
		artFontFiles.OutPutArtisticFont(artFontFiles.GORM_DEBUG)
	}
	switch FlagOptions.Type {
	case "user":
		u := flag_user.FlagUser{}
		switch FlagOptions.Sub {
		case "create":
			u.Creat()
			os.Exit(0)
		}
	}

}

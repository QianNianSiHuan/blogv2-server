package core

import (
	"blogv2/artFontFiles"
	"blogv2/conf"
	"blogv2/flags"
	"blogv2/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConf() (c *conf.Config) {
	byteDate, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		artFontFiles.OutPutArtisticFont(artFontFiles.FAIL)
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteDate, c)
	if err != nil {
		panic(fmt.Sprintf("yaml格式错误%s", err))
	}
	return
}
func SetConf() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Errorf("conf读取失败 %s ", err)
		return
	}
	err = os.WriteFile(flags.FlagOptions.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("设置配置文件失败 %s", err)
		return
	}
}

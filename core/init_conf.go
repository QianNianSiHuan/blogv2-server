package core

import (
	"blogv2/artFontFiles"
	"blogv2/conf"
	"blogv2/flags"
	"fmt"
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

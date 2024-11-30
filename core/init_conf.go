package core

import (
	"blogv2.0/conf"
	"blogv2.0/flags"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConf() (c *conf.Config) {
	byteDate, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteDate, c)
	if err != nil {
		panic(fmt.Sprintf("yaml格式错误%s", err))
	}
	fmt.Printf("读取配置文件 %s 成功\n", flags.FlagOptions.File)
	return
}

package main

import (
	"blogv2.0/core"
	"blogv2.0/flags"
	"blogv2.0/global"
	"fmt"
)

func main() {
	flags.Parse()
	fmt.Println(flags.FlagOptions)
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
}

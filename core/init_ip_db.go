package core

import (
	"embed"
	_ "embed"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

//go:embed files/ip2region.xdb
var ipdb embed.FS

func InitIPDb() *xdb.Searcher {
	//ProgressbarMsg <- "IP数据库初始化..."
	ip, err := ipdb.ReadFile("files/ip2region.xdb")
	if err != nil {
		logrus.Warnf("ip地址数据库文件加载失败 %s", err)
		return nil
	}
	//searcher, err := xdb.NewWithFileOnly(dbPath)
	searcher, err := xdb.NewWithBuffer(ip)
	if err != nil {
		logrus.Warnf("ip地址数据库加载失败 %s", err)
		return nil
	}
	logrus.Info("ip地址库加载成功")
	return searcher
}

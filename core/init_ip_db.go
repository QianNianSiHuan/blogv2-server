package core

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

func InitIPDb() *xdb.Searcher {
	//ProgressbarMsg <- "IP数据库初始化..."
	var dbPath = "init/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Warnf("ip地址数据库加载失败 %s", err)
		return nil
	}
	logrus.Info("ip地址库加载成功")
	return searcher
}

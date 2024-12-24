package core

import (
	ipUnitls "blogv2/unitls/ip"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

func InitIPDb() {
	ProgressbarMsg <- "IP数据库初始化..."
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Warnf("ip地址数据库加载失败 %s", err)
		return
	}
	searcher = _searcher
	logrus.Info("ip地址库加载成功")
}
func GetIpAddr(ip string) (addr string) {
	if ipUnitls.HasLocalIPAddr(ip) {
		logrus.Infof("ip: %s 地址: 内网", ip)
		return "内网"
	}
	addr, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("ip: %s 地址: ip错误", ip)
		return "ip错误"
	}
	addrList := strings.Split(addr, "|")
	if len(addrList) != 5 {
		logrus.Warnf("ip: %s 地址: ip异常", ip)
		return "ip异常"
	}
	//国家 0 省份 市区 运营商
	country := addrList[0]
	province := addrList[2]
	city := addrList[3]
	if province != "0" && city != "0" {
		logrus.Infof("ip地址解析成功 ip: %s 地址: %s", ip, addr)
		return fmt.Sprintf("%s·%s", province, city)
	}
	if province != "0" && country != "0" {
		logrus.Infof("ip地址解析成功 ip: %s 地址: %s", ip, addr)
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		logrus.Infof("ip地址解析成功 ip: %s 地址: %s", ip, addr)
		return fmt.Sprintf("%s", country)
	}
	return
}

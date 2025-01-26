package ip

import (
	"blogv2/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
)

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
func GetIpAddr(ip string) (addr string) {
	if HasLocalIPAddr(ip) {
		logrus.Infof("ip: %s 地址: 内网", ip)
		return "内网"
	}
	addr, err := global.IP.SearchByStr(ip)
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

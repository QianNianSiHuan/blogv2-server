package redis_login

import (
	"blogv2/global"
	"fmt"
	"time"
)

type LoginCacheType string

const (
	LoginCacheIP LoginCacheType = "login_ip_key"
	LoginCacheID LoginCacheType = "login_id_key"
)

func SetLoginCountByIP(ip string) {
	count := 0
	count = GetLoginCountByIP(ip) + 1
	global.Redis.Set(string(LoginCacheIP)+fmt.Sprintf("_%s", ip), count, 600*time.Second)
}
func SetLoginCountByID(id string) {
	count := 0
	count = GetLoginCountByID(id) + 1
	global.Redis.Set(string(LoginCacheID)+fmt.Sprintf("_%s", id), count, 600*time.Second)
}
func GetLoginCount(t LoginCacheType, val string) (count int) {
	count, _ = global.Redis.Get(string(t) + fmt.Sprintf("_%s", val)).Int()
	return
}
func GetLoginCountByIP(ip string) (count int) {
	count = GetLoginCount(LoginCacheIP, ip)
	return
}
func GetLoginCountByID(id string) (count int) {
	count = GetLoginCount(LoginCacheID, id)
	return
}
func ClearLoginCount(t LoginCacheType, val string) {
	global.Redis.Del(string(t) + fmt.Sprintf("_%s", val))
}
func ClearLoginCountByIP(ip string) {
	ClearLoginCount(LoginCacheIP, ip)
}
func ClearLoginCountByID(id string) {
	ClearLoginCount(LoginCacheID, id)
}

func ClearLoginCountAll(id string, ip string) {
	ClearLoginCount(LoginCacheID, id)
	ClearLoginCount(LoginCacheIP, ip)
}

package redis_jwt

import (
	"blogv2/global"
	jwts "blogv2/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type BlackType int8

const (
	UserBlackType   BlackType = 1 //用户注销登录
	AdminBlackType  BlackType = 2 //管理员让你手动下线
	DeviceBlackType BlackType = 3 //被其他设备挤下来
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}
func (b BlackType) Msg() string {
	switch b {
	case UserBlackType:
		return "已注销"
	case AdminBlackType:
		return "禁止登录"
	case DeviceBlackType:
		return "其他设备登录"
	}
	return "已注销"
}
func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}
func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}
	//计算token剩余过期时间
	second := claims.ExpiresAt.Unix() - time.Now().Unix()
	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf("redis黑名单添加失败 %s", err)
	}
}
func HasTokenBlack(token string) (blk BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blk = ParseBlackType(value)
	return blk, true
}
func HasTokenBlackByGin(c *gin.Context) (blk BlackType, ok bool) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
		if token == "" {
			token, _ = c.Cookie("token")
		}
	}
	return HasTokenBlack(token)
}

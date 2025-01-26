package log_service

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	ip2 "blogv2/utils/ip"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := ip2.GetIpAddr(ip)
	userID := uint(0)
	userName := ""
	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		userID = claims.UserID
		userName = claims.Username
	}
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    userName,
		Pwd:         "",
		LoginType:   loginType,
	})
}
func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, pwd string) {
	ip := c.ClientIP()
	addr := ip2.GetIpAddr(ip)
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}

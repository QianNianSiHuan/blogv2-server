package log_service

import (
	"blogv2.0/core"
	"blogv2.0/global"
	"blogv2.0/models"
	"blogv2.0/models/enum"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)
	token := c.GetHeader("token")
	//TODO:通过jwt获取userID
	fmt.Println(token)
	userID := 1
	userName := ""
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      uint(userID),
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
	addr := core.GetIpAddr(ip)
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}

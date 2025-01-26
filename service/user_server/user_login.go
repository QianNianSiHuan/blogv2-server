package user_server

import (
	"blogv2/global"
	"blogv2/models"
	ip2 "blogv2/utils/ip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (u UserService) UserLogin(c *gin.Context) {
	ip := c.ClientIP()
	addr := ip2.GetIpAddr(ip)
	ua := c.GetHeader("User-Agent")
	c.GetHeader("user")
	err := global.DB.Create(&models.UserLoginModel{
		UserID: u.userModel.ID,
		IP:     ip,
		Addr:   addr,
		UA:     ua,
	}).Error
	if err != nil {
		logrus.Errorf("用户登录日志写入失败")
		return
	}
}

package site_api

import (
	"blogv2.0/models/enum"
	"blogv2.0/service/log_service"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "sakura", "1234")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
}
func (SiteApi) SiteUpdateView(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
}

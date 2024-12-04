package site_api

import (
	"blogv2/models/enum"
	"blogv2/service/log_service"
	"github.com/gin-gonic/gin"
	"time"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "sakura", "1234")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
}

type SiteUpdateRequest struct {
	Name string `json:"name"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)
	log.ShowResponse()
	log.ShowRequest()
	log.ShowResponseHeader()
	log.ShowResquestHeader()
	log.SetTitle("更新站点")
	log.SetItemInfo("请求时间", time.Now())
	log.SetImage("http://qiniuyun.starletter.cn/picture/202411251703921.jpg")
	log.SetLink("学习地址", "http://qiniuyun.starletter.cn/picture/202411251703921.jpg")
	c.Header("666", "sakura")
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.SetError("参数绑定失败", err)
	}
	log.SetItemInfo("结构体", cr)
	log.SetItemInfo("切片", []string{"a", "b"})
	log.SetItemInfo("字符串", "你好")
	log.SetItemInfo("数字", 123)
	id := log.Save()
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息更改", "data": id})
}

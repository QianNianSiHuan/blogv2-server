package site_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/middleware"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}
type SiteInfoRequest struct {
	Name string `uri:"name"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
	}
	if cr.Name == "site" {
		res.SuccessWithData(c, global.Config.Site)
		return
	}
	//判断角色是否管理员
	middleware.AdminMiddleware(c)
	_, ok := c.Get("claims")
	if !ok {
		return
	}
	var data any
	switch cr.Name {
	case "email":
		data = global.Config.Email
	case "qq":
		data = global.Config.QQ
	case "ai":
		data = global.Config.Ai
	case "qiNiu":
		data = global.Config.QiNiu
	default:
		res.FailWithMsg(c, "不存在的配置")
		return
	}
	res.SuccessWithData(c, data)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required" label:"姓名"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	//log := log_service.GetLog(c)
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "更新成功")
	return
}

package site_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"github.com/gin-gonic/gin"
)

type AISiteInfoResponse struct {
	Enable   bool   `json:"enable"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Abstract string `json:"abstract"`
}

func (SiteApi) AISiteInfoView(c *gin.Context) {
	if !global.Config.Ai.Enable {
		res.FailWithMsg(c, "站点ai未启用")
		return
	}
	res.SuccessWithData(c, AISiteInfoResponse{
		Enable:   global.Config.Ai.Enable,
		Nickname: global.Config.Ai.Nickname,
		Avatar:   global.Config.Ai.Avatar,
		Abstract: global.Config.Ai.Abstract,
	})
}

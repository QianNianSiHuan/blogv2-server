package site_api

import (
	"blogv2/common/res"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.SuccessWithData(c, "xxx")
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

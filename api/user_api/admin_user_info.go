package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/unitls/maps"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AdminUserInfoUpdateRequest struct {
	UserID   *uint          `json:"userID" binding:"required"`
	Username *string        `json:"username" s-u:"username"`
	Nickname *string        `json:"nickname" s-u:"nickname"`
	Avatar   *string        `json:"avatar" s-u:"avatar"`
	Abstract *string        `json:"abstract" s-u:"abstract"`
	Role     *enum.RoleType `json:"role" s-u:"role"`
}

func (UserApi) AdminUserInfoUpdateView(c *gin.Context) {
	var cr AdminUserInfoUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	userMap, err := maps.StructToMap(cr, "s-u")
	if err != nil {
		res.FailWithMsg(c, "用户map转化失败")
		return
	}
	logrus.Info(userMap)
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		return
	}
	err = global.DB.Model(&user).Updates(userMap).Error
	if err != nil {
		res.FailWithMsg(c, "用户信息修改失败")
		return
	}

	res.SuccessWithMsg(c, "用户信息更新成功")
}

package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/maps"
	"github.com/gin-gonic/gin"
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
	claims := jwts.GetClaims(c)
	if cr.Role != nil && *cr.UserID != claims.UserID {
		res.FailWithMsg(c, "不能修改自己的角色")
		return
	}
	userMap, err := maps.StructToMap(cr, "s-u")
	if err != nil {
		res.FailWithMsg(c, "用户map转化失败")
		return
	}
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

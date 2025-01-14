package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/pwd"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RegisterAdminRequest struct {
	Username string `json:"username" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterAdminView(c *gin.Context) {
	var cr RegisterAdminRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	if claims.Role != enum.AdminRole {
		res.FailWithMsg(c, "权限不足")
		return
	}

	hashPwd, err := pwd.GenerateFromPassword(cr.Pwd)
	if err != nil {
		res.FailWithMsg(c, "密码转换失败")
		return
	}
	//创建用户
	var user = models.UserModel{
		Username:       cr.Username,
		Nickname:       cr.Username,
		RegisterSource: enum.RegisterAdminSourceType,
		Password:       hashPwd,
		Role:           enum.UserRole,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		logrus.Errorf("管理员注册账号失败")
		res.FailWithMsg(c, "管理员注册账号失败")
		return
	}
	res.SuccessWithMsg(c, "用户创建成功")
}

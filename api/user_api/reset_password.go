package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/utils/pwd"
	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Pwd string `json:"pwd" binding:"required"`
}

func (UserApi) ResetPasswordView(c *gin.Context) {
	var cr ResetPasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg(c, "站点未启用邮箱注册")
		return
	}
	_email, _ := c.Get("email")
	email := _email.(string)
	var user models.UserModel
	err = global.DB.Take(&user, "email = ?", email).Error
	if err != nil {
		res.FailWithMsg(c, "不存在的用户")
		return
	}
	//todo:该修改密码逻辑可后续优化
	if user.RegisterSource != enum.RegisterEmailSourceType {
		res.FailWithMsg(c, "非邮箱注册用户,不能更改密码")
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(user).Update("password", hashPwd)
	res.SuccessWithMsg(c, "密码重置成功")
}

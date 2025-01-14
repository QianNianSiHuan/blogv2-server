package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	Pwd    string `json:"pwd" binding:"required"`
}

func (UserApi) UpdatePasswordView(c *gin.Context) {
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		return
	}
	//邮箱注册的，绑定邮箱的
	if !(user.RegisterSource == enum.RegisterEmailSourceType || user.Email != "") {
		res.FailWithMsg(c, "不支持修改密码")
		return
	}
	//校验密码
	fmt.Println(user.Password)
	if !pwd.CompareHashAndPassword(user.Password, cr.OldPwd) {
		res.FailWithMsg(c, "旧密码错误")
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.SuccessWithMsg(c, "密码修改成功")
}

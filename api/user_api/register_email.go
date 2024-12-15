package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/unitls/email_store"
	jwts "blogv2/unitls/jwt"
	"blogv2/unitls/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type RegisterEmailRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterEmail(c *gin.Context) {
	var cr RegisterEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	value, ok := global.EmailVerifyStore.Load(cr.EmailID)
	if !ok {
		res.FailWithMsg(c, "邮箱验证失败")
		return
	}
	info, ok := value.(email_store.EmailStoreInfo)
	if !ok {
		res.FailWithMsg(c, "邮箱验证失败")
		return
	}
	if info.Code != cr.EmailCode {
		global.EmailVerifyStore.Delete(cr.EmailID)
		res.FailWithMsg(c, "邮箱验证失败")
		return
	}
	global.EmailVerifyStore.Delete(cr.EmailID)
	hashPwd, err := pwd.GenerateFromPassword(cr.Pwd)
	if err != nil {
		res.FailWithMsg(c, "密码转换失败")
		return
	}
	//创建用户
	uname := base64Captcha.RandText(5, "0123456789")
	var user = models.UserModel{
		Username:       fmt.Sprintf("email_%s", uname),
		Nickname:       "邮箱用户",
		RegisterSource: enum.RegisterEmailSourceType,
		Password:       hashPwd,
		Email:          info.Email,
		Role:           enum.UserRole,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		logrus.Errorf("邮箱注册失败")
		res.FailWithMsg(c, "邮箱注册失败")
		return
	}
	//颁发token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		res.FailWithMsg(c, "邮箱登录失败")
		return
	}
	res.SuccessWithData(c, token)
}

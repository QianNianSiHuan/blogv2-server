package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/user_server"
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

func (UserApi) RegisterEmailView(c *gin.Context) {
	var cr RegisterEmailRequest
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
		Email:          email,
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
	user_server.NewUserServiceApp(user).UserLogin(c)
	res.SuccessWithData(c, token)
}

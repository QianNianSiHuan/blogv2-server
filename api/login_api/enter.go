package login_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
	"regexp"
)

type LoginApi struct {
}

type UserLoginMsg struct {
	Username string `bind:"require" json:"username"`
	Password string `bind:"require" json:"password"`
	Remember bool   `bind:"require" json:"remember"`
}

func (LoginApi) Register(c *gin.Context) {
	var cr UserLoginMsg
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
	}
	matched, _ := regexp.MatchString(`^\d{11}$`, cr.Username)
	if !matched {
		res.FailWithMsg(c, "账号应为11位数字")
		return
	}
	if len(cr.Password) > 16 || len(cr.Password) < 8 {
		res.FailWithMsg(c, "密码长度应为8-16位")
		return
	}
	var count int64
	global.DB.Find(&models.UserModel{}, "username = ? ", cr.Username).Count(&count)
	if count != 0 {
		res.FailWithMsg(c, "账号已存在")
		return
	}
	err = global.DB.Create(&models.UserModel{
		Username:       cr.Username,
		RegisterSource: 1,
		Password:       cr.Password,
		Role:           2,
	}).Error
	if err != nil {
		res.FailWithMsg(c, "注册失败,服务出现错误")
	}
	res.SuccessWithMsg(c, "注册成功")
	return
}

func (LoginApi) Login(c *gin.Context) {
	var cr UserLoginMsg
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
	}
	matched, _ := regexp.MatchString(`^\d{11}$`, cr.Username)
	if !matched {
		res.FailWithMsg(c, "账号应为11位数字")
		return
	}
	if len(cr.Password) > 16 || len(cr.Password) < 8 {
		res.FailWithMsg(c, "密码长度应为8-16位")
		return
	}
	var tokenMsg = struct {
		ID       uint
		Username string
		Password string
		Role     enum.RoleType
	}{}
	var count int64
	global.DB.Select("password").Find(&models.UserModel{}, "username = ? ", cr.Username).Take(&tokenMsg).Count(&count)
	if count < 1 {
		res.FailWithMsg(c, "用户不存在")
		return
	}
	if tokenMsg.Password != cr.Password {
		res.FailWithMsg(c, "密码错误")
		return
	}
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   tokenMsg.ID,
		Username: tokenMsg.Username,
		Role:     tokenMsg.Role,
	})
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	c.SetCookie("token", token, 10800, "", "", false, true)
	res.SuccessWithMsg(c, "登陆成功")
}

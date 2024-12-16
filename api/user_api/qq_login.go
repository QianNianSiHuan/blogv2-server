package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/qq_service"
	"blogv2/service/user_server"
	jwts "blogv2/unitls/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type QQLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

func (UserApi) QQLoginView(c *gin.Context) {
	var cr QQLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if !global.Config.Site.Login.QQLogin {
		res.FailWithMsg(c, "站点未启用qq登录")
		return
	}
	info, err := qq_service.GetUserInfo(cr.Code)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		//创建用户
		uname := base64Captcha.RandText(8, "0123456789")
		user = models.UserModel{
			Username:       fmt.Sprintf("qq_%s", uname),
			Nickname:       info.Nickname,
			Avatar:         info.Avatar,
			RegisterSource: enum.RegisterQQSourceType,
			OpenID:         info.OpenID,
			Role:           enum.UserRole,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			logrus.Errorf(err.Error())
			res.FailWithMsg(c, "qq登录失败")
			return
		}
	}
	//颁发token
	token, _ := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	user_server.NewUserServiceApp(user).UserLogin(c)
	res.SuccessWithData(c, token)
}

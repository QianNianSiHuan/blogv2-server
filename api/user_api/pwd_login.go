package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/user_server"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/pwd"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginView(c *gin.Context) {
	var cr PwdLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg(c, "密码登录未启用")
	}
	var user models.UserModel
	err = global.DB.Take(&user, "(username=? or email = ?)and password <> ''",
		cr.Val, cr.Val,
	).Error
	if err != nil {
		res.FailWithMsg(c, "用户名或密码错误")
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg(c, "用户名或密码错误")
		return
	}

	token, _ := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	logrus.Info(user.Role)
	user_server.NewUserServiceApp(user).UserLogin(c)
	res.SuccessWithData(c, token)
}

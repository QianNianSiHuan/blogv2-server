package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

func (UserApi) BindEmailView(c *gin.Context) {
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg(c, "站点未启用邮箱注册")
		return
	}
	_email, _ := c.Get("email")
	email := _email.(string)

	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg(c, "密码转换失败")
		return
	}
	//更改绑定
	global.DB.Model(&user).Update("email", email)
	res.SuccessWithMsg(c, "邮箱绑定成功")
}

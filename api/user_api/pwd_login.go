package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/log_service"
	"blogv2/service/redis_service/redis_login"
	"blogv2/service/user_server"
	ip2 "blogv2/utils/ip"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/pwd"
	"github.com/gin-gonic/gin"
	"strconv"
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
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "(username=? or email = ?)and password <> ''",
		cr.Val, cr.Val,
	).Error
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		redis_login.SetLoginCountByIP(c.ClientIP())
		log_service.NewLoginFail(c, enum.UserPwdLoginType, "登录失败", cr.Val, cr.Password)
	}
	count := redis_login.GetLoginCountByID(strconv.Itoa(int(user.ID)))
	if count > 5 {
		res.FailWithMsg(c, "用户登录达到限制,请过段时间再次尝试")
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg(c, "用户密码错误")
		redis_login.SetLoginCountByID(strconv.Itoa(int(user.ID)))
		redis_login.SetLoginCountByIP(c.ClientIP())
		log_service.NewLoginFail(c, enum.UserPwdLoginType, "登录失败", cr.Val, cr.Password)
		return
	}
	token, _ := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	redis_login.ClearLoginCountAll(strconv.Itoa(int(user.ID)), c.ClientIP())
	ip := c.ClientIP()
	addr := ip2.GetIpAddr(ip)

	global.DB.Model(models.UserModel{}).Where("id = ?", user.ID).Updates(models.UserModel{
		IP:   ip,
		Addr: addr,
	})
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	user_server.NewUserServiceApp(user).UserLogin(c)
	res.SuccessWithData(c, token)
}

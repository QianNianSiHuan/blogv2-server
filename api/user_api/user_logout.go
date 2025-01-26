package user_api

import (
	"blogv2/common/res"
	"blogv2/service/redis_service/redis_jwt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "" {
		res.SuccessWithMsg(c, "注销成功")
		return
	}
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	res.SuccessWithMsg(c, "注销成功")
	return
}

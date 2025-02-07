package middleware

import (
	"blogv2/common/res"
	"blogv2/service/redis_service/redis_login"
	"github.com/gin-gonic/gin"
)

func LoginCountByIPMiddleware(c *gin.Context) {
	count := redis_login.GetLoginCountByIP(c.ClientIP())
	if count > 30 {
		res.FailWithMsg(c, "ip登录次数到达限制,请过段时间再次尝试")
		c.Abort()
		return
	}
}

package middleware

import (
	"blogv2/common/res"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_jwt"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
)

// 用户验证中间件
func AuthMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(c, err)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(c, blcType.Msg())
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}

// 管理员验证中间件
func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(c, err)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(c, blcType.Msg())
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg(c, "权限错误")
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}

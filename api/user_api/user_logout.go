package user_api

import (
	"blogv2/common/res"
	"blogv2/service/redis_service/redis_jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserLogoutVRequest struct {
	TokeBlackType redis_jwt.BlackType `json:"tokenBlackType" binding:"required,oneof=1 2 3"`
}

func (UserApi) UserLogoutView(c *gin.Context) {
	var cr UserLogoutVRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	token := c.GetHeader("token")
	fmt.Println(token)
	if token == "" {
		res.SuccessWithMsg(c, "注销成功")
		return
	}
	redis_jwt.TokenBlack(token, cr.TokeBlackType)
	res.SuccessWithMsg(c, "注销成功")
	return
}

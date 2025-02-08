package feed_back_api

import (
	"blogv2/common/res"
	"blogv2/service/email_service"
	"blogv2/service/redis_service/redis_email"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserFeedBackRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

func (FeedBackApi) UserFeedBackView(c *gin.Context) {
	var cr UserFeedBackRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	if redis_email.IsExistExEmailFeedBackID(strconv.Itoa(int(claims.UserID))) {
		res.FailWithMsg(c, "过段时间再次尝试吧")
		return
	}

	redis_email.SetEmailFeedBackID(strconv.Itoa(int(claims.UserID)))
	err = email_service.SendFeedBack(claims.Username, cr.Email, cr.Content)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "反馈信息发送成功")
}

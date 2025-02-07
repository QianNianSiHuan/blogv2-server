package middleware

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/service/redis_service/redis_login"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

type CaptchaMiddlewareRequest struct {
	CaptchaID   string `json:"captchaID" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func CaptchaMiddleware(c *gin.Context) {
	if !global.Config.Site.Login.Captcha {
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg(c, "请求体获取错误")
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr CaptchaMiddlewareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(c, "图形验证失败")
		c.Abort()
		return
	}
	if !global.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		redis_login.SetLoginCountByIP(c.ClientIP())
		res.FailWithMsg(c, "验证码错误")
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	c.Next()
}

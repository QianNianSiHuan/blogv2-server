package middleware

import (
	"blogv2/common/res"
	"blogv2/global"
	"bytes"
	"fmt"
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
		fmt.Println(cr.CaptchaID)
		fmt.Println(cr.CaptchaCode)
		res.SuccessWithMsg(c, "验证码错误")
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	c.Next()
}

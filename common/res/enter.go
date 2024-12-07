package res

import (
	"blogv2/unitls/vaildate"
	"github.com/gin-gonic/gin"
)

type Code int

const (
	SuccessCode     Code = 0
	FailValidCode   Code = 1001 //校验错误
	FailServiceCode Code = 1002 //服务错误
)

func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailValidCode:
		return "校验失败"
	case FailServiceCode:
		return "服务异常"
	default:
		return ""
	}
}

var empty = map[string]any{}

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json(c *gin.Context) {
	c.JSON(200, r)
}
func Success(c *gin.Context, msg string, data any) {
	Response{SuccessCode, data, msg}.Json(c)
}
func SuccessWithData(c *gin.Context, data any) {
	Response{SuccessCode, data, "success"}.Json(c)
}
func SuccessWithMsg(c *gin.Context, msg string) {
	Response{SuccessCode, empty, msg}.Json(c)
}
func SuccessWithList(c *gin.Context, list any, count int64) {
	Response{SuccessCode, map[string]any{
		"list":  list,
		"count": count,
	}, "success"}.Json(c)
}
func FailWithMsg(c *gin.Context, msg string) {
	Response{FailValidCode, empty, msg}.Json(c)
}
func FailWithData(c *gin.Context, msg string, data any) {
	Response{FailServiceCode, data, msg}.Json(c)
}
func FailWithCode(code Code, c *gin.Context) {
	Response{code, empty, code.String()}.Json(c)
}
func FailWithError(c *gin.Context, err error) {
	data, msg := vaildate.ValidateError(err)
	FailWithData(c, msg, data)
}

package middleware

import (
	"blogv2/service/log_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}
func (w *ResponseWriter) Header() http.Header {
	return w.Head
}
func LogMiddleware(c *gin.Context) {
	//接收日志中间件
	log := log_service.NewActionLogByGin(c)
	log.SetRequest(c)
	c.Set("log", log)
	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Head:           make(http.Header),
	}
	log.SetTitle("未命名日志")
	c.Writer = res
	c.Next()
	//响应日志中间件
	log = log_service.GetLog(c)
	log.ShowResquestHeader()
	log.ShowRequest()
	log.ShowResponseHeader()
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.MiddlewareSave()
}

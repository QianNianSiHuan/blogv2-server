package res

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SSESuccess(c *gin.Context, data any) {
	byteData, _ := json.Marshal(Response{SuccessCode, data, "success"})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}
func SSEFail(c *gin.Context, msg string) {
	byteData, _ := json.Marshal(Response{SuccessCode, struct{}{}, msg})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}

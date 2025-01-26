package web_header

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GetUserAgentInfo(c *gin.Context) *user_agent.UserAgent {
	uaHeader := c.GetHeader("User-Agent")
	return user_agent.New(uaHeader)
}

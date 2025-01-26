package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginRestrictionMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.Abort()
	} else {

	}
}

package log_service

import "github.com/gin-gonic/gin"

func LogRecordsAll(c *gin.Context, title string) *ActionLog {
	log := GetLog(c)
	log.SetTitle(title)
	return log
}

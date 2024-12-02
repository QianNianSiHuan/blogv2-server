package log_service

import (
	"blogv2.0/core"
	"blogv2.0/global"
	"blogv2.0/models"
	"blogv2.0/models/enum"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ActionLog struct {
	c     *gin.Context
	level enum.LogLevelType
	title string
}

func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

func (ac *ActionLog) SetLevle(level enum.LogLevelType) {
	ac.level = level
}
func (ac ActionLog) Save() {
	ip := ac.c.ClientIP()
	addr := core.GetIpAddr(ip)
	UserID := uint(1)
	err := global.DB.Create(models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.title,
		Content: "",
		Level:   ac.level,
		UserID:  UserID,
		IP:      ip,
		Addr:    addr,
	}).Error
	if err != nil {
		logrus.Errorf("日志创建失败 %s", err)
		return
	}
}
func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}

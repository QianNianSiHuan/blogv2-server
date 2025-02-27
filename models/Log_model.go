package models

import (
	"blogv2/models/enum"
	"gorm.io/gorm"
)

type LogModel struct {
	gorm.Model
	LogType     enum.LogType      `json:"logType"`
	Title       string            `gorm:"size:64" json:"title"`
	Content     string            `json:"content"`
	Level       enum.LogLevelType `json:"level"`
	UserID      uint              `json:"userID"`
	UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-"`
	IP          string            `gorm:"size:32" json:"ip"`
	Addr        string            `gorm:"size:64" json:"addr" `
	Method      string            `gorm:"size10" json:"method"`
	IsRead      bool              `json:"isRead"`
	LoginStatus bool              `json:"loginStatus"`
	Username    string            `gorm:"size:32" json:"username" `
	Pwd         string            `gorm:"size:32" json:"pwd" `
	LoginType   enum.LoginType    `json:"loginType" `
	ServiceName string            `gorm:"size:32" json:"serviceName"`
}

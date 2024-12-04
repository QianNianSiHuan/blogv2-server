package models

import "blogv2/models/enum"

type LogModel struct {
	Model       `json:"model"`
	LogType     enum.LogType      `json:"logType" json:"logType"`
	Title       string            `gorm:"size:64" json:"title" json:"title"`
	Content     string            `json:"content" json:"content"`
	Level       enum.LogLevelType `json:"level" json:"level"`
	UserID      uint              `json:"userID" json:"userID"`
	UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-" json:"userModel"`
	IP          string            `gorm:"size:32" json:"ip" json:"IP"`
	Addr        string            `gorm:"size:64" json:"addr" json:"addr"`
	IsRead      bool              `json:"isRead" json:"isRead"`
	LoginStatus bool              `json:"loginStatus" json:"loginStatus"`
	Username    string            `gorm:"size:32" json:"username" json:"username"`
	Pwd         string            `gorm:"size:32" json:"pwd" json:"pwd"`
	LoginType   enum.LoginType    `json:"loginType" json:"loginType"`
	ServiceName string            `gorm:"size:32" json:"serviceName"`
}

package models

import "blogv2/models/enum"

type LogModel struct {
	Model       `json:"model"`
	LogType     enum.LogType      `json:"logType" json:"logType,omitempty"`
	Title       string            `gorm:"size:64" json:"title" json:"title,omitempty"`
	Content     string            `json:"content" json:"content,omitempty"`
	Level       enum.LogLevelType `json:"level" json:"level,omitempty"`
	UserID      uint              `json:"userID" json:"userID,omitempty"`
	UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-" json:"userModel"`
	IP          string            `gorm:"size:32" json:"ip" json:"IP,omitempty"`
	Addr        string            `gorm:"size:64" json:"addr" json:"addr,omitempty"`
	IsRead      bool              `json:"isRead" json:"isRead,omitempty"`
	LoginStatus bool              `json:"loginStatus" json:"loginStatus,omitempty"`
	Username    string            `gorm:"size:32" json:"username" json:"username,omitempty"`
	Pwd         string            `gorm:"size:32" json:"pwd" json:"pwd,omitempty"`
	LoginType   enum.LoginType    `json:"loginType" json:"loginType,omitempty"`
	ServiceName string            `gorm:"size:32" json:"serviceName"`
}

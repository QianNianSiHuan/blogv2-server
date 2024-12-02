package models

import "time"

type UserModel struct {
	Model
	Username       string `gorm:"size:32" json:"username"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Avatar         string `gorm:"size:256" json:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract"`
	RegisterSource int8   `json:"registerSource"` //注册源
	CodeAge        int    `json:"codeAge"`
	Password       string `gorm:"size:64" json:"-"`
	Email          string `gorm:"size:256" json:"email"`
	OpenID         string `gorm:"size:64" json:"openID"`
	Role           int8   `json:"role"`
}
type UserConfModel struct {
	UserID             uint       `gorm:"unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameTime *time.Time `json:"updateUsernameTime"` //上次更新用户名的时间
	OpenCollect        bool       `json:"openCollect"`        //收藏
	OpenFollow         bool       `json:"openFollow"`         //关注
	OpenFans           bool       `json:"openFans"`           //粉丝
	HomeStyleID        uint       `json:"homeStyleID"`        //主页
}

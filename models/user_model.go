package models

import (
	"blogv2/models/enum"
	"gorm.io/gorm"
	"math"
	"time"
)

type UserModel struct {
	Model
	Username       string                  `gorm:"size:32;unique" json:"username"`
	Nickname       string                  `gorm:"size:32" json:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"avatar"`
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"` //注册源
	Password       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"`
	OpenID         string                  `gorm:"size:64" json:"openID"`
	Role           enum.RoleType           `json:"role"`
	UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID" json:"-"`
	IP             string                  `json:"ip"`
	Addr           string                  `json:"addr"`
}

func (u *UserModel) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Create(&UserConfModel{UserID: u.ID}).Error
	err = tx.Create(&CollectModel{UserID: u.ID, IsDefault: true, Title: "默认收藏夹"}).Error
	return
}
func (u *UserModel) CodeAge() int {
	sub := time.Now().Sub(u.CreatedAt)
	codeAge := int(math.Ceil(sub.Hours() / 24 / 365))
	return codeAge
}

type UserConfModel struct {
	UserID             uint       `gorm:"unique;primaryKey" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameTime *time.Time `json:"updateUsernameTime"` //上次更新用户名的时间
	OpenCollect        bool       `json:"openCollect"`        //收藏
	OpenFollow         bool       `json:"openFollow"`         //关注
	OpenFans           bool       `json:"openFans"`           //粉丝
	HomeStyleID        uint       `json:"homeStyleID"`        //主页
}

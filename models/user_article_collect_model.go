package models

import "time"

type UserArticleCollectModel struct {
	UserID       uint          `gorm:"uniqueIndex:idx_name" json:"userID"`
	UserModel    UserModel     `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    uint          `gorm:"uniqueIndex:idx_name" json:"articleID"`
	ArticleModel ArticleModel  `gorm:"foreignKey:ArticleID" json:"-"`
	CollectID    uint          `gorm:"uniqueIndex:idx_name" json:"collectID"`
	CollectModel CollectModels `gorm:"foreignKey:CollectID" json:"-"`
	CreatedAt    time.Time     `json:"createdAt"`
}

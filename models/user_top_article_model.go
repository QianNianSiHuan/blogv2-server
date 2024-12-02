package models

import "time"

type UserTopArticleModel struct {
	UserID       string       `gorm:"uniqueIndex:idx_name" json:"userID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    string       `gorm:"uniqueIndex:idx_name" json:"articleID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	CreatedAt    time.Time    `json:"createdAt"`
}

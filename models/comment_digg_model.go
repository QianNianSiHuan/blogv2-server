package models

import "time"

type CommentDiggModel struct {
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	CommentID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"`
	CommentModel CommentModel `gorm:"foreignKey:CommentID" json:"-"`
	CreatedAt    time.Time    `json:"createdAt"`
}

package models

import "blogv2/models/enum"

type CommentModel struct {
	Model
	Content        string             `gorm:"size:256" json:"content"`
	UserID         uint               `json:"userID"`
	UserModel      UserModel          `gorm:"foreignKey:UserID" json:"-"`
	ArticleID      uint               `json:"articleID"`
	ArticleModel   ArticleModel       `gorm:"foreignKey:ArticleID" json:"-"`
	ParentID       *uint              `json:"parentID"`
	ParentModel    *CommentModel      `gorm:"foreignKey:ParentID" json:"-"`
	SubCommentList []*CommentModel    `gorm:"foreignKey:ParentID" json:"-"` //子评论列表
	RootParentID   *uint              `json:"rootParentID"`
	DiggCount      int                `json:"diggCount"`
	ApplyCount     int                `json:"applyCount"` //评论回复数
	Status         enum.CommentStatus `json:"status"`     //评论状态 1 审核中 2已发布 3审核未通过
}

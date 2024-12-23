package models

import _ "embed"

type ArticleModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   uint      `json:"categoryID"` //分类
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList" `
	Cover        string    `gorm:"size:256" json:"cover"`
	UserID       uint      `json:"userID"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int       `json:"lookCount"`
	DiggCount    int       `json:"diggCount"`
	CommentCount int       `json:"commentCount"`
	CollectCount int       `json:"collectCount"`
	OpenComment  bool      `json:"openComment"` //评论开关
	Status       int8      `json:"status"`      //状态
}

//go:embed mappings/article_mapping.json
var articleMappings string

func (a ArticleModel) Mapping() string {
	return articleMappings
}
func (a ArticleModel) Index() string {
	return "article_index"
}

package models

type CollectModels struct {
	Model
	Title        string    `json:"title"`
	Abstract     string    `json:"abstract"`
	Cover        string    `json:"cover"`
	ArticleCount int       `json:"articleCount"`
	UserID       uint      `json:"userID"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"`
}

package models

type CategoryModel struct {
	Model
	Title     string    `gorm:"size:32" json:"title"`
	UserID    uint      `json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
}

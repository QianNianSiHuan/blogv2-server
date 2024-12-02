package models

type CommentModel struct {
	Model
	Content        string          `gorm:"size:256" json:"content"`
	UserID         uint            `json:"userID"`
	UserModel      UserModel       `gorm:"foreignKey:UserID" json:"-"`
	ArticleID      uint            `json:"articleID"`
	ArticleModel   ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`
	ParentID       *uint           `json:"parentID"`
	ParentModel    *CommentModel   `gorm:"foreignKey:ParentID" json:"-"`
	SubCommentList []*CommentModel `gorm:"foreignKey:ParentID" json:"-"` //子评论列表
	RootParentID   *uint           `json:"rootParentID"`
	DiggCount      int             `json:"diggCount"`
}

package models

import (
	"blogv2/global"
	"blogv2/models/ctype"
	"blogv2/models/enum"
	_ "embed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleModel struct {
	Model
	Title        string             `gorm:"size:32" json:"title"`
	Abstract     string             `gorm:"size:256" json:"abstract"`
	Content      string             `json:"content"`
	CategoryID   *uint              `json:"categoryID"` //分类
	TagList      ctype.List         `gorm:"type:longtext" json:"tagList" `
	Cover        string             `gorm:"size:256" json:"cover"`
	UserID       uint               `json:"userID"`
	UserModel    UserModel          `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int                `json:"lookCount"`
	DiggCount    int                `json:"diggCount"`
	CommentCount int                `json:"commentCount"`
	CollectCount int                `json:"collectCount"`
	OpenComment  bool               `json:"openComment"` //评论开关
	Status       enum.ArticleStatus `json:"status"`      //状态
}

//go:embed mappings/article_mapping.json
var articleMappings string

func (a ArticleModel) Mapping() string {
	return articleMappings
}
func (a ArticleModel) Index() string {
	return "article_index"
}

func (a *ArticleModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 评论
	var commentList []CommentModel
	global.DB.Find(&commentList, "article_id = ?", a.ID).Delete(&commentList)
	// 点赞
	var diggList []ArticleDiggModel
	global.DB.Find(&diggList, "article_id = ?", a.ID).Delete(&diggList)
	// 收藏
	var collectList []UserArticleCollectModel
	global.DB.Find(&collectList, "article_id = ?", a.ID).Delete(&collectList)
	// 置顶
	var topList []UserTopArticleModel
	global.DB.Find(&topList, "article_id = ?", a.ID).Delete(&topList)
	// 浏览
	var lookList []UserArticleLookHistoryModel
	global.DB.Find(&lookList, "article_id = ?", a.ID).Delete(&lookList)

	logrus.Infof("删除关联评论 %d 条", len(commentList))
	logrus.Infof("删除关联点赞 %d 条", len(diggList))
	logrus.Infof("删除关联收藏 %d 条", len(collectList))
	logrus.Infof("删除关联置顶 %d 条", len(topList))
	logrus.Infof("删除关联浏览 %d 条", len(lookList))
	return
}

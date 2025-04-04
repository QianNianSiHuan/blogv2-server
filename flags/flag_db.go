package flags

import (
	"blogv2/global"
	"blogv2/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
		&models.ArticleModel{},
		&models.ArticleDiggModel{},
		&models.CategoryModel{},
		&models.CollectModel{},
		&models.UserArticleCollectModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.GlobalNotificationModel{},
		&models.ImageModel{},
		&models.UserLoginModel{},
		&models.UserTopArticleModel{},
		&models.CommentDiggModel{},
		&models.TextModel{}, //全文搜索表
		&models.SiteFlowModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s ", err)
		return
	}
	logrus.Info("数据库迁移成功")
}

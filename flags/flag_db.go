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
		&models.CollectModels{},
		&models.UserArticleCollectModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.GlobalNotificationModel{},
		&models.ImageModel{},
		models.UserLoginModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s ", err)
		return
	}
	logrus.Info("数据库迁移成功")
}

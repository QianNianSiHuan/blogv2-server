package cron_service

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/redis_service/redis_comment"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SyncComment() {
	applyMap := redis_comment.GetAllCacheApply()
	diggMap := redis_comment.GetAllCacheDigg()

	var list []models.CommentModel
	global.DB.Find(&list)

	for _, model := range list {
		apply := applyMap[model.ID]
		digg := diggMap[model.ID]

		if apply == 0 || digg == 0 {
			continue
		}

		err := global.DB.Model(&model).Updates(map[string]any{
			"apply_count": gorm.Expr("apply_count + ?", apply),
			"digg_count":  gorm.Expr("digg_count + ?", digg),
		}).Error
		if err != nil {
			logrus.Errorf("更新失败 %s", err)
			continue
		}
		logrus.Infof("评论%d 更新成功", model.ID)
	}

	// 走完之后清空掉
	redis_comment.Clear()

}

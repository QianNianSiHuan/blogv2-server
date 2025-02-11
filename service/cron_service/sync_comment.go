package cron_service

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/log_service"
	"blogv2/service/redis_service/redis_comment"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SyncComment() {

	runtimeLog := log_service.NewRuntimeLog("数据库同步", 1)
	runtimeLog.SetTitle("评论数定时同步")
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
			runtimeLog.SetItemError("评论数同步失败", err)
			runtimeLog.Save()
			logrus.Errorf("更新失败 %s", err)
			continue
		}
		runtimeLog.SetItemError("评论数同步成功", err)
		runtimeLog.Save()
		logrus.Infof("评论%d 更新成功", model.ID)
	}
	// 走完之后清空掉
	redis_comment.Clear()
	runtimeLog.SetItemInfo("评论数缓存清空成功", "成功")
	runtimeLog.Save()
}

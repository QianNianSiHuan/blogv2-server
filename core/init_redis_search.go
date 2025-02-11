package core

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/text_service"
	"github.com/sirupsen/logrus"
)

func InitRedisSearch() {
	var list []text_service.ParticipleTextModel
	global.DB.Model(&models.TextModel{}).Find(&list)
	text_service.Participle(list)
	logrus.Info("redis全文查询索引构建完成")
}

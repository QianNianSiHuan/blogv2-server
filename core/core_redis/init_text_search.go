package core_redis

import (
	"blogv2/service/text_service"
	"github.com/sirupsen/logrus"
)

func initRedisTextSearch(list []text_service.ParticipleTextModel) {
	text_service.TextParticiple(list)
	logrus.Info("redis全文查询索引构建完成")
}

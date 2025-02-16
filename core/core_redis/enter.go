package core_redis

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/text_service"
	"github.com/sirupsen/logrus"
)

func InitRedisService() {
	var list []text_service.ParticipleArticleModel
	global.DB.Model(models.ArticleModel{}).Find(&list)
	logrus.Info("标签聚合加载中...")
	initRedisTagAgg(list)
	logrus.Info("标签聚合加载完成")
	logrus.Info("文章搜索分词加载中...")
	initRedisArticleSearch(list)
	logrus.Info("文章搜索分词加载完成")
	logrus.Info("文章搜索排序加载中...")
	initRedisArticleSort(list)
	logrus.Info("文章搜索排序加载完成")
	logrus.Info("全文搜索分词加载中...")
	initRedisTextSearch()
	logrus.Info("全文搜索分词加载完成")
}

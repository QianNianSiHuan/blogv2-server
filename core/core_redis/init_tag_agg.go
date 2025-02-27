package core_redis

import (
	"blogv2/service/redis_service/redis_article"
	"blogv2/service/text_service"
)

func initRedisTagAgg(list []text_service.ParticipleArticleModel) {
	for _, model := range list {
		for _, tag := range model.TagList {
			redis_article.RemoveTagAgg(model.ID, tag)
			redis_article.SetTagAgg(tag, model.ID)
			redis_article.SetTagAggAdd(tag)
		}
	}
}

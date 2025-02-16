package core_redis

import (
	"blogv2/service/text_service"
)

func initRedisArticleSearch(list []text_service.ParticipleArticleModel) {
	text_service.ArticleParticiple(list)
}

package core_redis

import (
	"blogv2/service/redis_service/redis_article"
	"blogv2/service/text_service"
)

func initRedisArticleSort(list ...text_service.ParticipleArticleModel) {
	cacheCollect := redis_article.GetAllCacheCollect(1)
	cacheComment := redis_article.GetAllCacheComment(1)
	cacheDigg := redis_article.GetAllCacheDigg(1)
	cacheLook := redis_article.GetAllCacheLook(1)
	for _, model := range list {
		text_service.DeleteTextParticiple(model.ID)
		redis_article.ClearArticleSortByID(model.ID)
		collectCount := cacheCollect[model.ID] + model.CollectCount
		commentCount := cacheComment[model.ID] + model.CommentCount
		diggCount := cacheDigg[model.ID] + model.DiggCount
		lookCount := cacheLook[model.ID] + model.LookCount
		redis_article.SetCacheCollectSortByCount(model.ID, collectCount)
		redis_article.SetCacheCommentSortByCount(model.ID, commentCount)
		redis_article.SetCacheDiggSortByCount(model.ID, diggCount)
		redis_article.SetCacheLookSortByCount(model.ID, lookCount)
		redis_article.SetCacheAllSort(model.ID)
	}
}

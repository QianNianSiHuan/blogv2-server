package observer_article

import "blogv2/service/redis_service/redis_article"

type ArticleCollect struct {
}

func NewArticleCollect() *ArticleCollect {
	return &ArticleCollect{}
}
func (a ArticleCollect) AfterArticleCollectIncr(articleID uint, status int8) {
	redis_article.SetCacheCollect(articleID, 1)
	redis_article.SetCacheCollectSort(articleID, 1)
}

func (a ArticleCollect) AfterArticleCollectDec(articleID uint, status int8) {
	redis_article.SetCacheCollect(articleID, -1)
	redis_article.SetCacheCollectSort(articleID, -1)
}

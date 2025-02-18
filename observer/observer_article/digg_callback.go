package observer_article

import "blogv2/service/redis_service/redis_article"

type ArticleDigg struct {
}

func NewArticleDigg() *ArticleDigg {
	return &ArticleDigg{}
}

func (a ArticleDigg) AfterArticleDiggIncr(articleID uint) {
	redis_article.SetCacheDigg(articleID, 1)
	redis_article.SetCacheAllSortIncr(articleID, 2*1)
}

func (a ArticleDigg) AfterArticleDiggDec(articleID uint) {
	redis_article.SetCacheDigg(articleID, -1)
	redis_article.SetCacheAllSortIncr(articleID, -2*1)
}

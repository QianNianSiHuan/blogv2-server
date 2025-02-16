package observer_article

import (
	"blogv2/service/redis_service/redis_article"
)

type ArticleComment struct {
}

func NewArticleComment() *ArticleComment {
	return &ArticleComment{}
}
func (a ArticleComment) AfterArticleCommentIncr(articleID uint) {
	redis_article.SetCacheComment(articleID, 1)
}

func (a ArticleComment) AfterArticleCommentDec(articleID uint, n int) {
	redis_article.SetCacheComment(articleID, n)
}

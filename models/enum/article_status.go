package enum

type ArticleStatus int8

const (
	ArticleDraftStatus     ArticleStatus = 1 //草稿
	ArticleStatusExamine   ArticleStatus = 2 //审核中
	ArticleStatusPublished ArticleStatus = 3 //已发布
	ArticleStatusFail      ArticleStatus = 4 //审核失败
)

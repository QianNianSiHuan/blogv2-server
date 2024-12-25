package enum

type ArticleStatus int8

const (
	DraftArticleStatus     ArticleStatus = 1 //草稿
	ExamineArticleStatus   ArticleStatus = 2 //审核中
	PublishedArticleStatus ArticleStatus = 3 //已发布
)

package observer_article

type ArticleExamine struct {
}

func NewArticleExamine() *ArticleExamine {
	return &ArticleExamine{}
}

func (a ArticleExamine) AfterArticleExamine(articleID uint, status int8) {

}

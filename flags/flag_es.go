package flags

import (
	"blogv2/models"
	"blogv2/service/es_service"
)

func EsIndex() {
	article := models.ArticleModel{}
	es_service.CreatIndexV2(article.Index(), article.Mapping())
}

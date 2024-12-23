package flags

import (
	"blogv2/artFontFiles"
	"blogv2/models"
	"blogv2/service/es_service"
	"github.com/sirupsen/logrus"
)

func EsIndex() {
	article := models.ArticleModel{}
	err := es_service.CreatIndexV2(article.Index(), article.Mapping())
	if err != nil {
		logrus.Errorf(err.Error())
		artFontFiles.OutPutArtisticFont(artFontFiles.FAIL)
		return
	}
}

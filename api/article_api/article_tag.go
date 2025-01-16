package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/ctype"
	"blogv2/models/enum"
	"blogv2/utils"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleTagOptions(c *gin.Context) {
	claims := jwts.GetClaims(c)
	var articleList []models.ArticleModel
	global.DB.Find(&articleList, "user_id = ? and status = ?", claims.UserID, enum.ArticleStatusPublished)
	var tagList ctype.List
	for _, model := range articleList {
		tagList = append(tagList, model.TagList...)
	}
	tagList = utils.Unique(tagList)
	var list = make([]models.OptionsResponse[string], 0)
	for _, s := range tagList {
		list = append(list, models.OptionsResponse[string]{
			Label: s,
			Value: s,
		})
	}
	res.SuccessWithData(c, list)
}

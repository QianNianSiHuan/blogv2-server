package search_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (SearchApi) ArticleSearchIndexView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "后台开始重建索引...")
	var articleList []models.ArticleModel
	global.DB.Model(models.ArticleModel{}).Where("id in ?", cr.IDList).Find(&articleList)
	for _, article := range articleList {
		err = article.AfterUpdate(global.DB)
		if err != nil {
			logrus.Error(err)
		}
	}
}

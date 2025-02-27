package search_api

import (
	"blogv2/common/res"
	"blogv2/core/core_redis"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/text_service"
	"github.com/gin-gonic/gin"
)

func (SearchApi) ArticleSearchIndexView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "后台开始重建索引...")
	var articleList []text_service.ParticipleArticleModel
	global.DB.Model(models.ArticleModel{}).Where("id in ?", cr.IDList).Find(&articleList)
	var textList []text_service.ParticipleTextModel
	global.DB.Model(&models.TextModel{}).Where("id in ?", cr.IDList).Find(&textList)
	core_redis.InitRedisIndex(articleList, textList)
	res.SuccessWithMsg(c, "索引重建成功")
}

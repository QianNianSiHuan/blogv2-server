package article_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"github.com/gin-gonic/gin"
)

type ArticleRecommendResponse struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	LookCount int    `json:"lookCount" gorm:"column:lookCount"`
}

func (ArticleApi) ArticleRecommendView(c *gin.Context) {
	var cr common.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var list = make([]ArticleRecommendResponse, 0)
	global.DB.Model(models.ArticleModel{}).
		Order("look_count desc").
		Limit(cr.Limit).Select("id", "title", "look_count").Scan(&list)

	res.SuccessWithList(c, list, int64(len(list)))
}

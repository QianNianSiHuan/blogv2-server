package search_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"github.com/gin-gonic/gin"
	"sync"
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
	var wg sync.WaitGroup
	for _, article := range articleList {
		go func() {
			wg.Add(1)
			article.AfterUpdate(global.DB)
			wg.Done()
		}()
	}
	wg.Wait()
}

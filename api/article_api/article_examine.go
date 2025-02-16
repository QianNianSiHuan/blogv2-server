package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/global/global_observer"
	"blogv2/models"
	"blogv2/models/enum"
	"github.com/gin-gonic/gin"
)

type ArticleExamineRequest struct {
	ArticleID uint               `json:"articleID" binding:"required"`
	Status    enum.ArticleStatus `json:"status" binding:"required,oneof=3 4"`
	Msg       string             `json:"msg"` // 为4的时候，传递进来
}

func (ArticleApi) ArticleExamineView(c *gin.Context) {
	var cr ArticleExamineRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}

	global.DB.Model(&article).Update("status", cr.Status)
	global_observer.ArticleNotifier.AfterArticleExamineNotify(article.ID, int8(cr.Status))
	// TODO: 给文章的发布人发一个系统通知

	res.SuccessWithMsg(c, "审核成功")
}

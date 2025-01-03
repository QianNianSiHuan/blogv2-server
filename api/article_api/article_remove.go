package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindUri(&cr)

	var list []models.ArticleModel
	global.DB.Take(&list, " id in ?", cr.IDList)

	err = global.DB.Delete(&list).Error
	if len(list) > 0 {
		if err != nil {
			logrus.Error(err.Error())
			res.FailWithError(c, err)
			return
		}
	}
	res.SuccessWithMsg(c, fmt.Sprintf("删除成功 %d 篇文章", len(list)))
}

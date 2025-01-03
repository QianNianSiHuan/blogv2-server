package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ArticleApi) ArticleRemoveUserView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindJSON(&cr)

	claims := jwts.GetClaims(c)
	var model models.ArticleModel
	err = global.DB.Take(&model, "user_id = ? and id =?", claims.UserID, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}

	err = global.DB.Delete(&model).Error
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMsg(c, "删除文章失败")
		return
	}

	res.SuccessWithMsg(c, "删除成功")
}

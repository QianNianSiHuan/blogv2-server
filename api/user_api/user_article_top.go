package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserArticleTopRequest struct {
	ArticleID uint `json:"articleID" binding:"required"`
	Type      int8 `json:"type" binding:"required,oneof=1 2" ` //1用户 2管理员置顶文章
}

func (UserApi) UserArticleTopView(c *gin.Context) {
	var cr UserArticleTopRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, "id = ?", cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}
	if article.Status != enum.ArticleStatusPublished {
		res.FailWithMsg(c, "文章未发布")
		return
	}
	claims := jwts.GetClaims(c)
	switch cr.Type {
	case 1:
		if claims.UserID != article.UserID {
			res.FailWithMsg(c, "权限不足")
			return
		}
		var UserTopArticleList []models.UserTopArticleModel
		err = global.DB.Find(&UserTopArticleList, "user_id = ?",
			claims.UserID).Error
		if len(UserTopArticleList) == 1 {
			uta := UserTopArticleList[0]
			if uta.ArticleID != cr.ArticleID {
				res.FailWithMsg(c, "置顶文章只能有一篇")
				return
			}
		}
		if len(UserTopArticleList) == 0 {
			global.DB.Create(&models.UserTopArticleModel{
				UserID:    claims.UserID,
				ArticleID: cr.ArticleID,
			})
			res.SuccessWithMsg(c, "文章置顶成功")
			return
		}
		global.DB.Delete(&UserTopArticleList[0])
		res.SuccessWithMsg(c, "取消文章置顶成功")
		return
	case 2:
		if claims.Role != enum.AdminRole {
			res.FailWithMsg(c, "权限不足")
			return
		}
		var UserTopArticle models.UserTopArticleModel
		err = global.DB.Take(&UserTopArticle, "user_id =? and article_id = ? ",
			claims.UserID, cr.ArticleID).Error
		if err != nil {
			global.DB.Create(&models.UserTopArticleModel{
				UserID:    claims.UserID,
				ArticleID: cr.ArticleID,
			})
			res.SuccessWithMsg(c, "文章置顶成功")
			return
		}
		global.DB.Delete(&UserTopArticle)
		res.SuccessWithMsg(c, "文章取消置顶成功")
		return
	}
}

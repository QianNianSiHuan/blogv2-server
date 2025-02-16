package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/global/global_observer"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

// 用户点赞
func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}
	claims := jwts.GetClaims(c)
	//查一下之前有没有点过赞
	var userDiggArticle models.ArticleDiggModel
	err = global.DB.Take(&userDiggArticle, "user_id = ? and article_id = ?", claims.UserID, article.ID).Error
	if err != nil {
		// 点赞
		err = global.DB.Create(&models.ArticleDiggModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		}).Error
		if err != nil {
			res.FailWithMsg(c, "点赞失败")
			return
		}
		// TODO: 更新点赞数到缓存里面
		//redis_article.SetCacheDigg(cr.ID, true)
		global_observer.ArticleNotifier.AfterArticleDiggIncrNotify(cr.ID)
		res.SuccessWithMsg(c, "点赞成功")
		return
	}
	// 取消点赞
	//redis_article.SetCacheDigg(cr.ID, false)
	global_observer.ArticleNotifier.AfterArticleDiggDecNotify(cr.ID)
	global.DB.Delete(&userDiggArticle)
	res.SuccessWithMsg(c, "取消点赞成功")
}

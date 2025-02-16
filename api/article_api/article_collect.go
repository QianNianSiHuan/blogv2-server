package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/global/global_observer"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleCollectRequest struct {
	ArticleID uint `json:"articleID" binding:"required"`
	CollectID uint `json:"collectID"`
}
type ArticleCollectResponse struct {
	Collect   bool `json:"collect"`
	IsCollect bool `json:"isCollect"`
}

func (ArticleApi) ArticleCollectView(c *gin.Context) {
	var cr = ArticleCollectRequest{}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}
	var collectModel models.CollectModel
	claims := jwts.GetClaims(c)
	if cr.CollectID == 0 {
		// 是默认收藏夹
		err = global.DB.Take(&collectModel, "user_id = ? and is_default = ?", claims.UserID, 1).Error
		if err != nil {
			// 创建一个默认收藏夹
			collectModel.Title = "默认收藏夹"
			collectModel.UserID = claims.UserID
			collectModel.IsDefault = true
			global.DB.Create(&collectModel)
		}
		cr.CollectID = collectModel.ID
	} else {
		// 判断收藏夹是否存在，并且是否是自己创建的
		err = global.DB.Take(&collectModel, "user_id = ? and  id = ?", claims.UserID, cr.CollectID).Error
		if err != nil {
			res.FailWithMsg(c, "收藏夹不存在")
			return
		}
	}

	// 判断是否收藏
	var articleCollect models.UserArticleCollectModel
	err = global.DB.Where(models.UserArticleCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		CollectID: cr.CollectID,
	}).Take(&articleCollect).Error
	//查不到记录就收藏
	if err != nil {
		// 收藏
		err = global.DB.Create(&models.UserArticleCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ArticleID,
			CollectID: cr.CollectID,
		}).Error
		if err != nil {
			res.FailWithMsg(c, "收藏失败")
			return
		}

		res.Success(c, "收藏成功", ArticleCollectResponse{
			Collect:   true,
			IsCollect: true})

		// 对收藏夹进行加1
		//redis_article.SetCacheCollect(cr.ArticleID, true)
		global_observer.ArticleNotifier.AfterArticleCollectIncrNotify(article.ID)
		//global_gse.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count + 1"))
		return
	}
	// 取消收藏
	//redis_article.SetCacheCollect(cr.ArticleID, false)
	err = global.DB.Delete(&articleCollect).Error
	if err != nil {
		res.FailWithMsg(c, "取消收藏失败")
		return
	}
	var articleCollectList []models.UserArticleCollectModel
	global.DB.Find(&articleCollectList, "user_id = ? and article_id = ? ", claims.UserID, cr.ArticleID)
	var isCollect bool
	if len(articleCollectList) > 0 {
		isCollect = true
	}
	res.Success(c, "取消收藏成功", ArticleCollectResponse{Collect: false,
		IsCollect: isCollect})
	//TODO:收藏数同步缓存
	global_observer.ArticleNotifier.AfterArticleCollectDecNotify(article.ID)
	global.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count - 1"))
	return
}

func (ArticleApi) ArticleCollectPatchRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	var userConfList []models.UserArticleCollectModel

	global.DB.Preload("CollectModel").Find(&userConfList, "id in ? and user_id = ?", cr.IDList, claims.UserID)
	var collectTitle string
	if len(userConfList) > 0 {
		collectTitle = userConfList[0].CollectModel.Title
		global.DB.Delete(&userConfList)
	}
	res.SuccessWithMsg(c, fmt.Sprintf("成功将 %d 篇文章, 移出 %s 收藏夹", len(userConfList), collectTitle))
}

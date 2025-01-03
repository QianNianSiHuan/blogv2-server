package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type ArticleLookRequest struct {
	ArticleID  uint `json:"articleID" binding:"required"`
	TimeSecond int  `json:"timeSecond"` // 读文章一共用了多久
}

func (ArticleApi) ArticleLookView(c *gin.Context) {
	var cr ArticleLookRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	// TODO: 未登录用户，浏览量如何加
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.SuccessWithMsg(c, "未登录")
		return
	}

	// 引入缓存
	if redis_article.GetUserArticleHistoryCache(cr.ArticleID, claims.UserID) {
		logrus.Infof("UserID: %d ArticleID: %d 在足迹缓存里面", cr.ArticleID, claims.UserID)
		res.SuccessWithMsg(c, "成功")
		return
	}
	// 当天这个用户请求这个文章之后，将用户id和文章id作为key存入缓存，在这里进行判断，如果存在就直接返回

	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}

	// 查这个文章今天有没有在足迹里面
	var history models.UserArticleLookHistoryModel
	err = global.DB.Take(&history,
		"user_id = ? and article_id = ? and created_at < ? and created_at > ?",
		claims.UserID, cr.ArticleID,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02")+" 00:00:00",
	).Error
	if err == nil {
		res.SuccessWithMsg(c, "成功")
		return
	}

	err = global.DB.Create(&models.UserArticleLookHistoryModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
	}).Error
	if err != nil {
		res.FailWithMsg(c, "失败")
		return
	}

	redis_article.SetCacheLook(cr.ArticleID, true)
	redis_article.SetUserArticleHistoryCache(cr.ArticleID, claims.UserID)
	res.SuccessWithMsg(c, "成功")
}

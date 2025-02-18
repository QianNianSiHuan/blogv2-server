package observer_article

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type ArticleLookHistory struct {
}

func NewArticleLookHistory() *ArticleLookHistory {
	return &ArticleLookHistory{}
}

func (a *ArticleLookHistory) AfterArticleLook(c *gin.Context, articleID uint) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		//res.SuccessWithMsg(c, "未登录")
		return
	}
	// 引入缓存
	if redis_article.GetUserArticleHistoryCache(articleID, claims.UserID) {
		logrus.Infof("UserID: %d ArticleID: %d 在足迹缓存里面", articleID, claims.UserID)
		//res.SuccessWithMsg(c, "成功")
		return
	}
	// 当天这个用户请求这个文章之后，将用户id和文章id作为key存入缓存，在这里进行判断，如果存在就直接返回
	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, articleID).Error
	if err != nil {
		//res.FailWithMsg(c, "文章不存在")
		return
	}

	// 查这个文章今天有没有在足迹里面
	var history models.UserArticleLookHistoryModel
	err = global.DB.Take(&history,
		"user_id = ? and article_id = ? and created_at < ? and created_at > ?",
		claims.UserID, articleID,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02")+" 00:00:00",
	).Error
	if err == nil {
		//res.SuccessWithMsg(c, "成功")
		return
	}

	err = global.DB.Create(&models.UserArticleLookHistoryModel{
		UserID:    claims.UserID,
		ArticleID: articleID,
	}).Error
	if err != nil {
		//res.FailWithMsg(c, "失败")
		return
	}
	redis_article.SetCacheLook(articleID, 1)
	redis_article.SetUserArticleHistoryCache(articleID, claims.UserID)
	redis_article.SetCacheAllSortIncr(articleID, 1)
	res.SuccessWithMsg(c, "成功")
}

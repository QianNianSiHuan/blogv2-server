package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	UserAvatar string `json:"userAvatar"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.IDRequest
	err := c.BindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	// 未登录的用户，只能看到发布成功的文章

	// 登录用户，能看到自己的所有文章

	// 管理员，能看到全部的文章
	var article models.ArticleModel
	err = global.DB.Preload("UserModel").Take(&article, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		// 没登录的
		if article.Status != enum.ArticleStatusPublished {
			res.FailWithMsg(c, "文章不存在")
			return
		}
	}
	switch claims.Role {
	case enum.UserRole:
		if claims.UserID != article.UserID {
			// 登录的人看到不是自己的
			if article.Status != enum.ArticleStatusPublished {
				res.FailWithMsg(c, "文章不存在")
				return
			}
		}
	}
	// TODO: 从缓存里面获取浏览量和点赞数
	collectCount := redis_article.GetCacheCollect(article.ID)
	diggCount := redis_article.GetCacheDigg(article.ID)
	lookCount := redis_article.GetCacheLook(article.ID)
	CommentCount := redis_article.GetCacheComment(article.ID)

	article.DiggCount = article.DiggCount + diggCount
	article.CollectCount = article.CollectCount + collectCount
	article.LookCount = article.LookCount + lookCount
	article.CommentCount = article.CommentCount + CommentCount
	res.SuccessWithData(c, ArticleDetailResponse{
		ArticleModel: article,
		Username:     article.UserModel.Username,
		Nickname:     article.UserModel.Nickname,
		UserAvatar:   article.UserModel.Avatar,
	})
}

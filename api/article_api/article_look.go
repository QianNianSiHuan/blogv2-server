package article_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/utils/jwt"
	"fmt"
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

	redis_article.SetCacheLook(cr.ArticleID, 1)
	redis_article.SetUserArticleHistoryCache(cr.ArticleID, claims.UserID)
	res.SuccessWithMsg(c, "成功")
}

type ArticleLookListRequest struct {
	common.PageInfo
	UserID uint `form:"userID"`
	Type   int8 `form:"type" binding:"required,oneof=1 2"`
}

type ArticleLookListResponse struct {
	ID        uint      `json:"id"`       // 浏览记录的id
	LookDate  time.Time `json:"lookDate"` // 浏览的时间
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	UserID    uint      `json:"userID"`
	ArticleID uint      `json:"articleID"`
}

func (ArticleApi) ArticleLookListView(c *gin.Context) {
	var cr ArticleLookListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	claims := jwts.GetClaims(c)

	switch cr.Type {
	case 1:
		cr.UserID = claims.UserID
	}

	_list, count, _ := common.ListQuery(models.UserArticleLookHistoryModel{
		UserID: cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"UserModel", "ArticleModel"},
	})

	var list = make([]ArticleLookListResponse, 0)
	for _, model := range _list {
		list = append(list, ArticleLookListResponse{
			ID:        model.ID,
			LookDate:  model.CreatedAt,
			Title:     model.ArticleModel.Title,
			Cover:     model.ArticleModel.Cover,
			Nickname:  model.UserModel.Nickname,
			Avatar:    model.UserModel.Avatar,
			UserID:    model.UserID,
			ArticleID: model.ArticleID,
		})
	}

	res.SuccessWithList(c, list, count)
}

func (ArticleApi) ArticleLookRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	var list []models.UserArticleLookHistoryModel
	global.DB.Find(&list, "user_id = ? and id in ?", claims.UserID, cr.IDList)

	if len(list) > 0 {
		err := global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg(c, "足迹删除失败")
			return
		}
	}
	res.SuccessWithMsg(c, fmt.Sprintf("删除足迹成功 共删除%d条", len(list)))
}

package article_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleListRequest struct {
	common.PageInfo
	Type       int8               `form:"type" binding:"required,oneof=1 2 3"` // 1 用户查别人的  2 查自己的  3 管理员查
	UserID     uint               `form:"userID"`
	CategoryID *uint              `form:"categoryID"`
	Status     enum.ArticleStatus `form:"status"`
	CollectID  uint               `form:"collectID"`
}

type ArticleListResponse struct {
	models.ArticleModel
	UserTop       bool    `json:"userTop"`       //用户置顶
	AdminTop      bool    `json:"adminTop"`      //管理员置顶
	CategoryTitle *string `json:"categoryTitle"` //分类标签
	UserNickname  string  `json:"userNickname"`
	UserAvatar    string  `json:"userAvatar"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var orderColumnOrder = map[string]bool{
		"look_count desc":    true,
		"digg_count desc":    true,
		"comment_count desc": true,
		"collect_count desc": true,
		"look_count asc":     true,
		"digg_count asc":     true,
		"comment_count asc":  true,
		"collect_count asc":  true,
	}
	var topArticleIDList []uint
	switch cr.Type {
	case 1:
		// 查别人。用户id就是必填的
		if cr.UserID == 0 {
			res.FailWithMsg(c, "用户id必填")
			return
		}
		if cr.Page > 2 || cr.Limit > 10 {
			res.FailWithMsg(c, "查询更多，请登录")
			return
		}
		cr.Status = 0
		cr.Order = ""
		if cr.CollectID != 0 {
			//如果传了收藏夹的id,需要判断
			if cr.UserID == 0 {
				res.FailWithMsg(c, "清传入用户ID")
				return
			}
			var userConf models.UserConfModel
			err = global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
			if err != nil {
				res.FailWithMsg(c, "用户不存在")
				return
			}
			if !userConf.OpenCollect {
				res.FailWithMsg(c, "用户未开启我的收藏")
				return
			}
		}
	case 2:
		// 查自己的
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithMsg(c, "请登录")
			return
		}
		cr.UserID = claims.UserID
	case 3:
		// 管理员
		claims, err := jwts.ParseTokenByGin(c)
		if !(err == nil && claims.Role == enum.AdminRole) {
			res.FailWithMsg(c, "角色错误")
			return
		}
	}
	if cr.Order != "" {
		_, ok := orderColumnOrder[cr.Order]
		if !ok {
			res.FailWithMsg(c, "非法的排序字段")
			return
		}
	}

	query := global.DB.Where("")

	if cr.CollectID != 0 {
		var articleIDList []uint
		global.DB.Model(models.UserArticleCollectModel{}).Where("collect_id = ?", cr.CollectID).Select("article_id").Scan(&articleIDList)

		query.Where("id in ?", articleIDList)
	}

	var userTopMap = map[uint]bool{}
	var adminTopMap = map[uint]bool{}
	var userTopQuery = global.DB.Where("")
	if cr.UserID != 0 {
		userTopQuery.Where("user_id = ?", cr.UserID)
	}
	var userTopArticleList []models.UserTopArticleModel
	global.DB.Preload("UserModel").Order("created_at desc").Where(userTopQuery).Find(&userTopArticleList)

	for _, i2 := range userTopArticleList {
		topArticleIDList = append(topArticleIDList, i2.ArticleID)
		if i2.UserModel.Role == enum.AdminRole {
			adminTopMap[i2.ArticleID] = true
		}
		userTopMap[i2.ArticleID] = true
	}
	var options = common.Options{
		Likes:        []string{"title"},
		PageInfo:     cr.PageInfo,
		Where:        query,
		DefaultOrder: "created_at desc",
		Preloads:     []string{"CategoryModel", "UserModel"},
	}

	if len(topArticleIDList) > 0 {
		options.DefaultOrder = fmt.Sprintf("%s, created_at desc", sql.ConvertSliceOrderSql(topArticleIDList))
	}

	_list, count, _ := common.ListQuery(models.ArticleModel{
		UserID:     cr.UserID,
		CategoryID: cr.CategoryID,
		Status:     cr.Status,
	}, options)

	var list = make([]ArticleListResponse, 0)
	collectMap := redis_article.GetAllCacheCollect(0)
	diggMap := redis_article.GetAllCacheDigg(0)
	lookMap := redis_article.GetAllCacheLook(0)
	commentMap := redis_article.GetAllCacheComment(0)
	for _, model := range _list {
		model.Content = ""
		model.DiggCount = model.DiggCount + diggMap[model.ID]
		model.CollectCount = model.CollectCount + collectMap[model.ID]
		model.LookCount = model.LookCount + lookMap[model.ID]
		model.CommentCount = model.CommentCount + commentMap[model.ID]
		data := ArticleListResponse{
			ArticleModel: model,
			UserTop:      userTopMap[model.ID],
			AdminTop:     adminTopMap[model.ID],
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
		}
		if model.CategoryModel != nil {
			data.CategoryTitle = &model.CategoryModel.Title
		}
		list = append(list, data)
	}
	res.SuccessWithList(c, list, count)
}

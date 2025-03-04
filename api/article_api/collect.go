package article_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CollectCreateRequest struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" binding:"required,max=32"`
	Abstract string `json:"abstract"`
	Cover    string `json:"cover"`
}

func (ArticleApi) CollectCreateView(c *gin.Context) {
	var cr CollectCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	var model models.CollectModel
	if cr.ID == 0 {
		// 创建
		err := global.DB.Take(&model, "user_id = ? and title = ?", claims.UserID, cr.Title).Error
		if err == nil {
			res.FailWithMsg(c, "收藏夹名称重复")
			return
		}

		err = global.DB.Create(&models.CollectModel{
			Title:    cr.Title,
			UserID:   claims.UserID,
			Abstract: cr.Abstract,
			Cover:    cr.Cover,
		}).Error
		if err != nil {
			res.FailWithMsg(c, "创建收藏夹失败")
			return
		}

		res.SuccessWithMsg(c, "创建收藏夹成功")
		return
	}

	err = global.DB.Take(&model, "user_id = ? and id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "收藏夹不存在")
		return
	}

	err = global.DB.Model(&model).Updates(map[string]any{
		"title":    cr.Title,
		"abstract": cr.Abstract,
		"cover":    cr.Cover,
	}).Error
	if err != nil {
		res.FailWithMsg(c, "更新收藏夹错误")
		return
	}

	res.SuccessWithMsg(c, "更新收藏夹成功")
	return
}

type CollectListRequest struct {
	common.PageInfo
	UserID    uint `form:"userID"`
	Type      int8 `form:"type" binding:"required,oneof=1 2 3"` // 1 查自己 2 查别人 3 后台
	ArticleID uint `form:"articleID"`
}

type CollectListResponse struct {
	models.CollectModel
	ArticleCount int    `json:"articleCount"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	ArticleUse   bool   `json:"articleUse,omitempty"`
}

func (ArticleApi) CollectListView(c *gin.Context) {
	var cr CollectListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var preload = []string{"ArticleList"}
	switch cr.Type {
	case 1:
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithError(c, err)
			return
		}
		cr.UserID = claims.UserID
	case 2:
		var userConf models.UserConfModel
		err = global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
		if err != nil {
			res.FailWithMsg(c, "用户不存在")
			return
		}
		if !userConf.OpenCollect {
			res.FailWithMsg(c, "用户未开启收藏")
			return
		}
	case 3:
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithError(c, err)
			return
		}
		if claims.Role != enum.AdminRole {
			res.FailWithMsg(c, "权限错误")
			return
		}
		preload = append(preload, "UserModel")
	}

	_list, count, _ := common.ListQuery(models.CollectModel{
		UserID: cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Likes:    []string{"title"},
		Preloads: preload,
	})

	var list = make([]CollectListResponse, 0)
	for _, i2 := range _list {
		item := CollectListResponse{
			CollectModel: i2,
			ArticleCount: len(i2.ArticleList),
			Nickname:     i2.UserModel.Nickname,
			Avatar:       i2.UserModel.Avatar,
		}
		for _, model := range i2.ArticleList {
			if model.ArticleID == cr.ArticleID {
				item.ArticleUse = true
				break
			}
		}
		list = append(list, item)
	}

	res.SuccessWithList(c, list, count)
}

func (ArticleApi) CollectRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var list []models.CollectModel
	query := global.DB.Where("id in ?", cr.IDList)
	claims := jwts.GetClaims(c)
	if claims.Role != enum.AdminRole {
		query.Where("user_id = ?", claims.UserID)
	}

	global.DB.Where(query).Find(&list)

	if len(list) > 0 {
		err := global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg(c, "删除分类失败")
			return
		}
	}

	msg := fmt.Sprintf("删除收藏夹成功 共删除%d条", len(list))

	res.SuccessWithMsg(c, msg)
}

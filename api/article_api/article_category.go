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

type CategoryCreateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" binding:"required,max=32"`
}

func (ArticleApi) CategoryCreateView(c *gin.Context) {
	var cr CategoryCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	claims := jwts.GetClaims(c)
	var model models.CategoryModel
	if cr.ID == 0 {
		// 创建
		err := global.DB.Take(&model, "user_id = ? and title = ?", claims.UserID, cr.Title).Error
		if err == nil {
			res.FailWithMsg(c, "分类名称重复")
			return
		}

		err = global.DB.Create(&models.CategoryModel{
			Title:  cr.Title,
			UserID: claims.UserID,
		}).Error
		if err != nil {
			res.FailWithMsg(c, "创建分类错误")
			return
		}

		res.SuccessWithMsg(c, "创建分类成功")
		return
	}

	err = global.DB.Take(&model, "user_id = ? and id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "分类不存在")
		return
	}

	err = global.DB.Model(&model).Update("title", cr.Title).Error

	if err != nil {
		res.FailWithMsg(c, "更新分类错误")
		return
	}

	res.SuccessWithMsg(c, "更新分类成功")
	return
}

type CategoryListRequest struct {
	common.PageInfo
	UserID uint `form:"userID"`
	Type   int8 `form:"type" binding:"required,oneof=1 2 3"` // 1 查自己 2 查别人 3 后台
}

type CategoryListResponse struct {
	models.CategoryModel
	ArticleCount int    `json:"articleCount"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
}

func (ArticleApi) CategoryListView(c *gin.Context) {
	var cr CategoryListRequest
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

	_list, count, err := common.ListQuery(models.CategoryModel{
		UserID: cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Likes:    []string{"title"},
		Preloads: preload,
	})
	if err != nil {
		res.FailWithMsg(c, "列表查询出错")
		return
	}
	var list = make([]CategoryListResponse, 0)

	for _, i2 := range _list {
		fmt.Println(i2.UserModel)
		list = append(list, CategoryListResponse{
			CategoryModel: i2,
			ArticleCount:  len(i2.ArticleList),
			Nickname:      i2.UserModel.Nickname,
			Avatar:        i2.UserModel.Avatar,
		})
	}

	res.SuccessWithList(c, list, count)
}
func (ArticleApi) CategoryRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	var list []models.CategoryModel
	query := global.DB.Where("id in ?", cr.IDList)
	claims := jwts.GetClaims(c)
	if claims.Role != enum.AdminRole {
		query.Where("user_id = ?", claims.UserID)
	}

	global.DB.Where(query).Find(&list)

	if len(list) > 0 {
		err = global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg(c, "删除分类失败")
			return
		}
	}

	msg := fmt.Sprintf("删除分类成功 共删除%d条", len(list))

	res.SuccessWithMsg(c, msg)
}

func (ArticleApi) CategoryOptionsView(c *gin.Context) {
	claims := jwts.GetClaims(c)

	var list []models.OptionsResponse[uint]
	global.DB.Model(models.CategoryModel{}).Where("user_id = ?", claims.UserID).
		Select("id as value", "title as label").Scan(&list)
	res.SuccessWithData(c, list)
}

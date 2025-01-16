package article_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/ctype"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/markdown"
	"blogv2/utils/xss"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleUpdateRequest struct {
	ID          uint               `json:"ID" binding:"required"`
	Title       string             `json:"title" binding:"required"`
	Abstract    string             `json:"abstract"`
	Content     string             `json:"content" binding:"required"`
	CategoryID  *uint              `json:"categoryID"`
	TagList     ctype.List         `json:"tagList"`
	Cover       string             `json:"cover"`
	OpenComment bool               `json:"openComment"`
	Status      enum.ArticleStatus `json:"status" binding:"required,oneof=1 2"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println("err:", err)
		res.FailWithError(c, err)
		return
	}
	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		return
	}
	//找文章
	var article models.ArticleModel
	err = global.DB.Take(&article, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}
	if article.UserID != user.ID {
		res.FailWithMsg(c, "文章归属出错")
		return
	}
	//判断分类ID是否自己创建的
	var category models.CategoryModel
	if cr.CategoryID != nil {
		err = global.DB.Take(&category, "id = ? and user_id =?", *cr.CategoryID, user.ID).Error
		if err != nil {
			res.FailWithMsg(c, "文章分类不存在")
			return
		}
	}
	//文章正文放xss注入
	cr.Content = xss.XSSFilter(cr.Content)
	//如果不传简介，从正文获取内容
	if cr.Abstract == "" {
		//把markdown转成html
		doc, err := markdown.ExtractContent(cr.Content, 100)
		if err != nil {
			res.FailWithMsg(c, "正文解析错误")
			return
		}
		cr.Abstract = doc
	}

	mps := map[string]any{
		"title":        cr.Title,
		"abstract":     cr.Abstract,
		"content":      cr.Content,
		"category_id":  cr.CategoryID,
		"tag_list":     cr.TagList,
		"cover":        cr.Cover,
		"open_comment": cr.OpenComment,
	}
	if article.Status == enum.ArticleStatusPublished && !global.Config.Site.Article.NoExamine {
		// 如果是已发布的文章，进行编辑，那么就要改成待审核
		mps["status"] = enum.ArticleStatusExamine
	}
	if article.Status == enum.ArticleStatusFail && !global.Config.Site.Article.NoExamine {
		// 如果是已拒绝的文章，进行编辑，那么就要改成待审核
		mps["status"] = enum.ArticleStatusExamine
	}

	err = global.DB.Model(&article).Updates(mps).Error
	if err != nil {
		res.FailWithMsg(c, "更新失败")
		return
	}
	//正文内容图片转存
	res.SuccessWithMsg(c, "文章更新成功")
}

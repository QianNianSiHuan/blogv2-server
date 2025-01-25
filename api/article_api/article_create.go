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

type ArticleCreateRequest struct {
	Title       string             `json:"title" binding:"required"`
	Abstract    string             `json:"abstract"`
	Content     string             `json:"content" binding:"required"`
	CategoryID  *uint              `json:"categoryID"`
	TagList     ctype.List         `json:"tagList"`
	Cover       string             `json:"cover"`
	OpenComment bool               `json:"openComment"`
	Status      enum.ArticleStatus `json:"status" binding:"required,oneof=1 2"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleCreateRequest
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
	if global.Config.Site.SiteInfo.Mode == 2 && user.Role != enum.AdminRole {
		res.FailWithMsg(c, "博客模式下,普通用户不能发文章")
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

	//正文内容图片转存
	var article = models.ArticleModel{
		Title:       cr.Title,
		Abstract:    cr.Abstract,
		Content:     cr.Content,
		UserID:      user.ID,
		CategoryID:  cr.CategoryID,
		TagList:     cr.TagList,
		Cover:       cr.Cover,
		OpenComment: cr.OpenComment,
		Status:      cr.Status,
	}
	if global.Config.Site.Article.NoExamine && cr.Status == 2 {
		article.Status = enum.ArticleStatusPublished
	}
	err = global.DB.Create(&article).Error
	if err != nil {
		res.FailWithMsg(c, "文章创建失败")
		return
	}
	res.SuccessWithMsg(c, "文章创建成功")
}

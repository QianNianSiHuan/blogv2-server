package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/comment_service"
	"blogv2/service/redis_service/redis_article"
	"blogv2/service/redis_service/redis_comment"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
)

type CommentCreateRequest struct {
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"articleID" binding:"required"`
	ParentID  *uint  `json:"parentID"` //父级
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	var article models.ArticleModel
	err = global.DB.Take(&article, "id = ? and status =?",
		cr.ArticleID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
	}

	claims := jwts.GetClaims(c)

	model := models.CommentModel{
		Content:   cr.Content,
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		ParentID:  cr.ParentID,
	}
	//去找这个评论的根评论
	if cr.ParentID != nil {
		//有父评论
		parentList := comment_service.GetParents(*cr.ParentID)
		if len(parentList)+1 > global.Config.Site.Article.CommentLine {
			res.FailWithMsg(c, "评论层级达到限制")
			return
		}
		if len(parentList) > 0 {
			model.RootParentID = &parentList[len(parentList)-1].ID
			for _, commentModel := range parentList {
				redis_comment.SetCacheApply(commentModel.ID, 1)
			}
		}
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		res.SuccessWithMsg(c, "评论发布失败")
		return
	}
	redis_article.SetCacheComment(cr.ArticleID, 1)
	res.SuccessWithMsg(c, "评论发布成功")
}

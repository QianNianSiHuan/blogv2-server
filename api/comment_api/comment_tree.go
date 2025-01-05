package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/comment_service"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentTreeView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id =?", enum.ArticleStatusPublished, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "文章不存在")
		return
	}
	//把根评论查出来
	var commentList []models.CommentModel
	global.DB.Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]comment_service.CommentResponse, 0)
	for _, model := range commentList {
		response := comment_service.GetCommentTreeV4(model.ID)
		list = append(list, *response)
	}
	res.SuccessWithList(c, list, int64(len(list)))
}

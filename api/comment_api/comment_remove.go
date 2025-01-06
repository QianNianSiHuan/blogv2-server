package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/comment_service"
	"blogv2/service/redis_service/redis_comment"
	jwts "blogv2/unitls/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	claims := jwts.GetClaims(c)

	var model models.CommentModel
	err = global.DB.Preload("ArticleModel").Take(&model, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "评论不存在")
		return
	}

	if claims.Role != enum.AdminRole {
		// 普通用户只能删自己发的评论，或者自己发的文章的评论
		if !(model.UserID == claims.UserID || model.ArticleModel.UserID == claims.UserID) {
			res.FailWithMsg(c, "权限错误")
			return
		}
	}

	// 找所有的子评论，还要找所有的父评论
	subList := comment_service.GetCommentOneDimensional(model.ID)

	if model.ParentID != nil {
		// 有父评论
		parentList := comment_service.GetParents(*model.ParentID)
		for _, commentModel := range parentList {
			redis_comment.SetCacheApply(commentModel.ID, -len(subList))
		}
	}
	// 删评论
	global.DB.Delete(&subList)

	msg := fmt.Sprintf("删除成功，共删除评论%d条", len(subList))
	res.SuccessWithMsg(c, msg)

}

package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/redis_service/redis_comment"
	jwts "blogv2/unitls/jwt"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentDiggView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var comment models.CommentModel
	err = global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "评论不存在")
		return
	}

	claims := jwts.GetClaims(c)

	// 查一下之前有没有点过
	var userDiggComment models.CommentDiggModel
	err = global.DB.Take(&userDiggComment, "user_id = ? and comment_id = ?", claims.UserID, comment.ID).Error
	if err != nil {
		// 点赞
		err = global.DB.Create(&models.CommentDiggModel{
			UserID:    claims.UserID,
			CommentID: cr.ID,
		}).Error
		if err != nil {
			res.FailWithMsg(c, "点赞失败")
			return
		}
		redis_comment.SetCacheDigg(cr.ID, 1)
		res.SuccessWithMsg(c, "点赞成功")
		return
	}
	// 取消点赞
	global.DB.Model(models.CommentDiggModel{}).Delete("user_id = ? and comment_id = ?", claims.UserID, comment.ID)
	res.SuccessWithMsg(c, "取消点赞成功")
	redis_comment.SetCacheDigg(cr.ID, -1)
	return
}

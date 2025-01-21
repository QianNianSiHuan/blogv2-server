package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/comment_service"
	"blogv2/utils"
	jwts "blogv2/utils/jwt"
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

	var userDiggCommentMap = map[uint]bool{}
	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		// 登录了
		var commentList []models.CommentModel // 文章的评论id列表
		global.DB.Find(&commentList, "article_id = ?", cr.ID)

		if len(commentList) > 0 {
			// 查我点赞的评论id列表
			var commentIDList []uint
			var userIDList []uint
			for _, model := range commentList {
				commentIDList = append(commentIDList, model.ID)
				userIDList = append(userIDList, model.UserID)
			}
			userIDList = utils.Unique(userIDList) // 对用户id去重
			var commentDiggList []models.CommentDiggModel
			global.DB.Find(&commentDiggList, "user_id = ? and comment_id in ?", claims.UserID, commentIDList)
			for _, model := range commentDiggList {
				userDiggCommentMap[model.CommentID] = true
			}
		}
	}
	//把根评论查出来
	var commentList []models.CommentModel
	global.DB.Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]comment_service.CommentResponse, 0)
	for _, model := range commentList {
		response := comment_service.GetCommentTreeV4(model.ID, userDiggCommentMap)
		list = append(list, *response)
	}
	res.SuccessWithList(c, list, int64(len(list)))
}

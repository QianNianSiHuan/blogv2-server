package comment_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"github.com/gin-gonic/gin"
)

type CommentExamineRequest struct {
	CommentID uint               `json:"commentID"`
	Status    enum.CommentStatus `json:"status"` //"0"全部"1"待审核"2"已发布"3"未通过
}

func (CommentApi) CommentExamineView(c *gin.Context) {
	var cr CommentExamineRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	err = global.DB.Where("id = ? ", cr.CommentID).Model(&models.CommentModel{}).Update("status", cr.Status).Error
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "评论审核成功")
}

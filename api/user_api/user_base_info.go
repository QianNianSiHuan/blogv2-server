package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserBaseInfoResponse struct {
	UserID       uint   `json:"userID"`
	CodeAge      int    `json:"codeAge"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	LookCount    int    `json:"lookCount"`
	ArticleCount int    `json:"articleCount"`
	FansCount    int    `json:"fansCount"`
	FollowCount  int    `json:"followCount"`
	Place        string `json:"place"` //ip归属地
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var user models.UserModel
	logrus.Info(cr.ID)
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		return
	}
	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.CodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    1,
		ArticleCount: 1,
		FansCount:    1,
		FollowCount:  1,
		Place:        user.Addr,
	}
	res.SuccessWithData(c, data)
}

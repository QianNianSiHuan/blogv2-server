package data_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_site"
	"github.com/gin-gonic/gin"
)

type SumResponse struct {
	FlowCount     int   `json:"flowCount"`
	UserCount     int64 `json:"userCount"`
	ArticleCount  int64 `json:"articleCount"`
	MessageCount  int64 `json:"messageCount"`
	CommentCount  int64 `json:"commentCount"`
	NewLoginCount int64 `json:"newLoginCount"`
	NewSignCount  int64 `json:"newSignCount"`
}

func (DataApi) SumView(c *gin.Context) {
	var data SumResponse
	data.FlowCount = redis_site.GetFlow()
	global.DB.Model(models.UserModel{}).Count(&data.UserCount)
	global.DB.Model(models.ArticleModel{}).Where("status = ?", enum.ArticleStatusPublished).Count(&data.ArticleCount)
	global.DB.Model(models.CommentModel{}).Count(&data.CommentCount)
	global.DB.Model(models.UserLoginModel{}).Where("date(created_at) = date(now())").Count(&data.NewLoginCount)
	global.DB.Model(models.UserModel{}).Where("date(created_at) = date(now())").Count(&data.NewSignCount)
	res.SuccessWithData(c, data)
}

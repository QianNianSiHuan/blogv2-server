package ai_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/service/ai_service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleAnalysisRequest struct {
	Content string `json:"content" binding:"required"`
}

type ArticleAnalysisResponse struct {
	Title    string   `json:"title"`
	Abstract string   `json:"abstract"`
	Category string   `json:"category"`
	Tag      []string `json:"tag"`
}

func (AiApi) AiArticleAnalysis(c *gin.Context) {
	var cr ArticleAnalysisRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if !global.Config.Ai.Enable {
		res.FailWithMsg(c, "站点未启用ai功能")
		return
	}
	fmt.Println(cr.Content)
	msg, err := ai_service.AiChat(cr.Content)
	if err != nil {
		logrus.Errorf("ai分析失败 %s %s", err, cr.Content)
		res.FailWithError(c, err)
		return
	}
	var data ArticleAnalysisResponse
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		logrus.Errorf("ai分析失败 %s %s", err, msg)
		res.FailWithMsg(c, "ai分析失败")
		return
	}
	res.SuccessWithData(c, data)
}

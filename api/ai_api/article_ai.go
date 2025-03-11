package ai_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/ai_service"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
)

type ArticleAiRequest struct {
	Content   string `form:"content" binding:"required"`
	ArticleID uint   `form:"articleID"`
	Type      int    `form:"type" binding:"required,oneof=1 2"` //1.文章推荐2.文章分析
}

func (a AiApi) ArticleAiView(c *gin.Context) {
	var cr ArticleAiRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.SSEFail(c, "query参数绑定失败")
		return
	}
	var writelist []llms.MessageContent
	if cr.Type == 2 {
		var article models.ArticleModel
		global.DB.Find(&article, "id = ?", cr.ArticleID)
		writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, "你是一个博客网站的文章概要总结助手,会文章内容进行文章总结,一会儿我会给你一篇文章,请你对其进行内容总结,字数不多于500字"))
		writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, article.Content))
		err = ai_service.ChatStream(writelist, c)
		if err != nil {
			res.SSEFail(c, err.Error())
		}
		return
	}
	ai_service.AiSearch(c, cr.Content)
}

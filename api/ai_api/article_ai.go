package ai_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/ai_service"
	"blogv2/service/redis_service/redis_article"
	"encoding/json"
	"fmt"
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
	}
	fmt.Println("cr:", cr)
	var writelist []llms.MessageContent
	idList, _ := redis_article.GetArticleSearchIndex(cr.Content)
	if len(idList) >= 10 {
		idList = idList[0:10]
	}
	if cr.Type == 2 {
		var article models.ArticleModel
		global.DB.Find(&article, "id = ?", cr.ArticleID)
		writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, "你是一个博客网站的文章概要总结助手,会文章内容进行文章总结"))
		writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, "请对下面文章尽行总结:", article.Content))
		err = ai_service.ChatStream(writelist, c)
		if err != nil {
			res.SSEFail(c, err.Error())
		}
		return
	}
	var articleList []models.ArticleModel
	global.DB.Select("id", "title", "abstract", "tag_list").Where("id in ?", idList).Find(&articleList)
	fmt.Println(articleList)
	_articleList, _ := json.Marshal(articleList)
	fmt.Println(_articleList)
	writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, "你是一个博客网站的文章检索推荐助手,会根据文章列表和用户的输入进行文章推荐"))
	writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem, "这是一个json格式的文章列表:", string(_articleList)))
	writelist = append(writelist, llms.TextParts(llms.ChatMessageTypeSystem,
		"请根据:\"", cr.Content, "\" 推荐文章列表中的一篇或多篇文章 按照固定格式 文章题目:<a href=/article/id>title</a>\n 推荐理由:\n 进行返回"))
	err = ai_service.ChatStream(writelist, c)
	if err != nil {
		res.SSEFail(c, err.Error())
	}
	//idList, words := redis_article.GetArticleSearchIndex()
}

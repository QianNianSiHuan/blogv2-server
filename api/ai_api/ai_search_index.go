package ai_api

import (
	"blogv2/common/res"
	"blogv2/service/ai_service"
	"github.com/gin-gonic/gin"
)

func (AiApi) AiSearchIndexView(c *gin.Context) {
	ai_service.AiSearchIndex()
	res.SuccessWithMsg(c, "文章Ai索引重建中...")
}

package ai_service

import (
	"blogv2/common/res"
	"blogv2/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func AiSearch(c *gin.Context, key string) {
	llm, err := openai.New(
		openai.WithModel("qwen-plus"),
		openai.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
		openai.WithToken(global.Config.Ai.SecretKey),
		openai.WithEmbeddingModel("text-embedding-v2"),
	)
	redisUrl := "redis://:" + global.Config.Redis.Password + "@" + global.Config.Redis.Addr + "/0"
	// 创建embedder
	openAiEmbedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		logrus.Error(err)
	}
	// 基于redis存储向量
	redisStore, err := redisvector.New(context.Background(),
		redisvector.WithConnectionURL(redisUrl),
		redisvector.WithIndexName("vector_idx", true),
		redisvector.WithEmbedder(openAiEmbedder),
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你是一个博客网站的文章推荐助手,给你一组json列表数据,你需要根据文章数据给用户推荐文章,推荐的文章需要理由,推荐文章格式需要有标题 如果有符合要求的文章返回，则按照固定格式 <a href=\"/article/id\">title</a> 进行文章推荐"),
	}
	docs, err := redisStore.SimilaritySearch(context.Background(), key, 10, vectorstores.WithScoreThreshold(0.5))
	fmt.Println(len(docs))
	for _, s := range docs {
		_s, err := json.Marshal(s)
		if err != nil {
			logrus.Error(err)
			return
		}
		content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, string(_s)))
	}
	content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, key))
	//将vector检索接入chains中
	_, err = llm.GenerateContent(context.Background(), content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		res.SSESuccess(c, string(chunk))
		return nil
	}))
	if err != nil {
		logrus.Error(err)
		return
	}
}

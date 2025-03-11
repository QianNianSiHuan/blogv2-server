package ai_service

import (
	"blogv2/global"
	"blogv2/models"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func AiSearchIndex() {
	llm, err := openai.New(
		openai.WithModel("qwen-omni-turbo"),
		openai.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
		openai.WithToken(global.Config.Ai.SecretKey),
		openai.WithEmbeddingModel("text-embedding-v2"),
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	redisUrl := "redis://:" + global.Config.Redis.Password + "@" + global.Config.Redis.Addr + "/0"
	logrus.Info(redisUrl)
	// 创建embedder
	openAiEmbedder, err := embeddings.NewEmbedder(llm, embeddings.WithBatchSize(25))
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
	}
	err = redisStore.DropIndex(context.Background(), "vector_idx", true)
	if err != nil {
		// 处理错误
		logrus.Error(err)
	}
	var articleList []models.ArticleModel
	global.DB.Preload("UserModel").Find(&articleList)
	mark := textsplitter.NewMarkdownTextSplitter(textsplitter.WithCodeBlocks(true), textsplitter.WithJoinTableRows(true))

	var count int
	for _, article := range articleList {
		var indexCount int
		textList, err := mark.SplitText(article.Content)
		if err != nil {
			logrus.Error(err)
			return
		}
		mp := map[string]any{
			"auth":      article.UserID,
			"articleID": article.ID,
			"title":     article.Title,
			"abstract":  article.Abstract,
		}
		for k, val := range textList {
			indexCount++
			var data []schema.Document
			data = append(data, schema.Document{
				PageContent: val,
				Metadata:    mp,
			})
			_, err = redisStore.AddDocuments(context.Background(), data, vectorstores.WithDeduplicater(func(ctx context.Context, doc schema.Document) bool {
				if k == 0 {
					return false
				}
				return textList[k-1] == doc.PageContent
			}))
			if err != nil {
				logrus.Error(err)
				return
			}
		}
		count++
		logrus.Infof("共 %d 篇文章,当前文章为第 %d 篇,当前文章共分 %d 段, 当前完成第 %d 段索引创建", len(articleList), count, len(textList), indexCount)
	}
}

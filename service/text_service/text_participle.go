package text_service

import (
	"blogv2/global/global_gse"
	"blogv2/models/ctype"
	"blogv2/service/redis_service/redis_article"
	"github.com/sirupsen/logrus"
	"sync"
)

type ParticipleTextModel struct {
	ID uint `json:"id"`
	TextModel
}
type ParticipleArticleModel struct {
	ID           uint       `json:"id"`
	Title        string     `json:"title"`
	Abstract     string     `json:"abstract"`
	LookCount    int        `json:"lookCount"`
	DiggCount    int        `json:"diggCount"`
	CommentCount int        `json:"commentCount"`
	CollectCount int        `json:"collectCount"`
	TagList      ctype.List `json:"tagList" `
}

func TextSearchParticiple(textList ...ParticipleTextModel) {
	allCount := len(textList)
	count := 0
	var wg sync.WaitGroup
	for _, text := range textList {
		go func() {
			wg.Add(1)
			count++
			logrus.Infof("开始创建分词... (文章ID:%d 段落总数:%d 当前分词段落:%d)", text.ArticleID, allCount, count)
			words := TextParticiple(text.Head, text.Body)
			redis_article.SetTextSearchIndex(text.ID, words)
			redis_article.SetTextSearchWords(text.ArticleID, text.ID, words)
			logrus.Infof("分词创建成功 (文章ID:%d 段落总数:%d 当前分词段落:%d)", text.ArticleID, allCount, count)
			wg.Done()
		}()
	}
	wg.Wait()
}

func ArticleSearchParticiple(textList ...ParticipleArticleModel) {
	for _, text := range textList {
		words := TextParticiple(text.Title, text.Abstract)
		redis_article.SetArticleSearchIndex(text.ID, words)
		redis_article.SetArticleSearchWords(text.ID, words)
	}
}

func DeleteArticleParticiple(articleID uint) {
	redis_article.DeleteArticleSearchIndexWords(articleID)
}
func DeleteTextParticiple(articleID uint) {
	redis_article.DeleteTextSearchIndexWords(articleID)
}

func TextParticiple(textList ...string) (words []string) {
	for _, text := range textList {
		word := global_gse.Gse.CutSearch(text, true)
		words = append(words, word...)
	}
	return
}

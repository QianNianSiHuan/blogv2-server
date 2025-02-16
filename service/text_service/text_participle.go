package text_service

import (
	"blogv2/global/global_gse"
	"blogv2/models/ctype"
	"blogv2/service/redis_service/redis_article"
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

func TextParticiple(textList []ParticipleTextModel) {
	for _, text := range textList {
		words := textParticiple(text.Head, text.Body)
		redis_article.SetTextSearchIndex(text.ID, words)
		redis_article.SetTextSearchWords(text.ArticleID, text.ID, words)
	}
}

func ArticleParticiple(textList []ParticipleArticleModel) {
	for _, text := range textList {
		words := textParticiple(text.Title, text.Abstract)
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

func textParticiple(textList ...string) (words []string) {
	for _, text := range textList {
		word := global_gse.Gse.CutSearch(text, true)
		words = append(words, word...)
	}
	return
}

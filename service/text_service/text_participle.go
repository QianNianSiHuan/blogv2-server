package text_service

import (
	"blogv2/global/global_gse"
	"blogv2/service/redis_service/redis_article"
)

type ParticipleTextModel struct {
	ID uint `json:"id"`
	TextModel
}

func Participle(textList []ParticipleTextModel) {
	for _, text := range textList {
		words := textParticiple(text.Head, text.Body)
		redis_article.SetTextSearchIndex(text.ID, words)
		redis_article.SetTextSearchWords(text.ArticleID, text.ID, words)
	}
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

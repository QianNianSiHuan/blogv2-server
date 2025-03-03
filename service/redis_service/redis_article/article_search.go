package redis_article

import (
	"blogv2/global"
	"blogv2/global/global_gse"
	"encoding/json"
	"fmt"
	"strconv"
)

func SetArticleSearchIndex(textID uint, words []string) {
	for _, word := range words {
		global.Redis.SAdd(fmt.Sprintf("article_%s", word), textID)
	}
}
func GetArticleSearchIndex(text string) (idList []string, words []string) {
	words = global_gse.Gse.CutSearch(text, true)
	var _words []string
	for _, word := range words {
		_words = append(_words, fmt.Sprintf("article_%s", word))
	}
	idList, _ = global.Redis.SUnion(_words...).Result()
	return
}

func DeleteArticleSearchIndex(words []string, articleID uint) {
	for _, word := range words {
		global.Redis.SRem(fmt.Sprintf("article_%s", word), articleID)
	}
}

type articleSearchType string

const (
	ArticleSearchWords articleSearchType = "article_search_words"
)

func SetArticleSearchWords(articleID uint, words []string) {
	_words, _ := json.Marshal(words)
	global.Redis.HSet(string(ArticleSearchWords), strconv.Itoa(int(articleID)), _words)
}

func GetArticleSearchWords(articleID uint) string {
	result, _ := global.Redis.HGet(string(ArticleSearchWords), strconv.Itoa(int(articleID))).Result()
	return result
}

func DeleteArticleSearchIndexWords(articleID uint) {
	result := GetArticleSearchWords(articleID)
	var words []string
	_ = json.Unmarshal([]byte(result), &words)
	DeleteArticleSearchIndex(words, articleID)
	deleteArticleSearchWordsIndex(articleID)
}

func deleteArticleSearchWordsIndex(articleID uint) {
	global.Redis.HDel(string(ArticleSearchWords), strconv.Itoa(int(articleID)))
}

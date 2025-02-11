package redis_article

import (
	"blogv2/global"
	"encoding/json"
	"fmt"
	"strconv"
)

func SetTextSearchIndex(textID uint, words []string) {
	for _, word := range words {
		global.Redis.SAdd(word, textID)
	}
}
func GetTextSearchIndex(word string) []string {
	vals, _ := global.Redis.SMembers(word).Result()
	return vals
}

func DeleteTextSearchIndex(words []string, textID uint) {
	for _, word := range words {
		global.Redis.SRem(word, textID)
	}
}

type articleSearchType string

const (
	ArticleSearchWordsType articleSearchType = "article_search_words"
)

func SetTextSearchWords(articleID uint, textID uint, words []string) {
	_words, _ := json.Marshal(words)
	global.Redis.HSet(string(ArticleSearchWordsType)+fmt.Sprintf("_%d", articleID), strconv.Itoa(int(textID)), _words)
}

func GetTextSearchWords(articleID uint) map[string]string {
	result, _ := global.Redis.HGetAll(string(ArticleSearchWordsType) + fmt.Sprintf("_%d", articleID)).Result()
	return result
}

func DeleteTextSearchIndexWords(articleID uint) {
	result := GetTextSearchWords(articleID)
	for key, val := range result {
		var words []string
		var textID int
		_ = json.Unmarshal([]byte(val), &words)
		textID, _ = strconv.Atoi(key)
		DeleteTextSearchIndex(words, uint(textID))
	}
	deleteTextSearchWordsIndex(articleID)
}

func deleteTextSearchWordsIndex(articleID uint) {
	global.Redis.HDel(string(ArticleSearchWordsType) + fmt.Sprintf("_%d", articleID))
}

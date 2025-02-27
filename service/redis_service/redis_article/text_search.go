package redis_article

import (
	"blogv2/global"
	"blogv2/global/global_gse"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

func SetTextSearchIndex(textID uint, words []string) {
	for _, word := range words {
		if word == "" {
			continue
		}
		global.Redis.SAdd(fmt.Sprintf("text_%s", word), textID)
	}
}
func GetTextSearchIndex(text string) []string {
	words := global_gse.Gse.CutSearch(text, true)
	var _words []string
	for _, word := range words {
		_words = append(_words, fmt.Sprintf("text_%s", word))
	}
	vals, _ := global.Redis.SUnion(_words...).Result()
	return vals
}

func DeleteTextSearchIndex(words []string, textID uint) {
	fmt.Println(len(words))
	for _, word := range words {
		if word == "" {
			continue
		}
		global.Redis.SRem(fmt.Sprintf("text_%s", word), textID)
	}
}

type textSearchType string

const (
	TextSearchWords textSearchType = "text_search_words"
)

func SetTextSearchWords(articleID uint, textID uint, words []string) {
	_words, _ := json.Marshal(words)
	global.Redis.HSet(string(TextSearchWords)+fmt.Sprintf("_%d", articleID), strconv.Itoa(int(textID)), _words)
}

func GetTextSearchWords(articleID uint) map[string]string {
	result, err := global.Redis.HGetAll(string(TextSearchWords) + fmt.Sprintf("_%d", articleID)).Result()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return result
}

func DeleteTextSearchIndexWords(articleID uint) {
	result := GetTextSearchWords(articleID)
	for key, val := range result {
		var words []string
		var textID int
		_ = json.Unmarshal([]byte(val), &words)
		fmt.Println(words)
		textID, _ = strconv.Atoi(key)
		DeleteTextSearchIndex(words, uint(textID))
	}
	deleteTextSearchWordsIndex(articleID)
}

func deleteTextSearchWordsIndex(articleID uint) {
	global.Redis.HDel(string(TextSearchWords) + fmt.Sprintf("_%d", articleID))
}

package text_service

import (
	"blogv2/global"
	"github.com/siddontang/go/log"
	"strings"
)

type TextModel struct {
	ArticleID uint   `json:"article_id"`
	Head      string `json:"head"`
	Body      string `json:"body"`
}

func MdContentTransformation(id uint, title string, content string) (list []TextModel) {
	lines := strings.Split(content, "\n")
	var headList []string
	var bodyList []string
	var body string
	headList = append(headList, title)
	var flag bool
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			flag = !flag
		}
		if !flag && strings.HasPrefix(line, "#") {
			// 标题行
			headList = append(headList, getHead(line))
			if strings.TrimSpace(body) != "" {
				bodyList = append(bodyList, getBody(body))
			}
			body = ""
			continue
		}
		body += line
	}
	if body != "" {
		bodyList = append(bodyList, getBody(body))
	}
	if len(headList) > len(bodyList) {
		bodyList = append(bodyList, "")
	}

	if len(headList) != len(bodyList) {
		log.Errorf("headList与bodyList 不一致 \n headList:%q  %d\\\n bodyList: %q  %d\n", headList, len(headList), bodyList, len(bodyList))
		return
	}

	for i := 0; i < len(headList); i++ {
		list = append(list, TextModel{
			ArticleID: id,
			Head:      headList[i],
			Body:      bodyList[i],
		})
	}
	return
}

func getHead(head string) string {
	s := strings.TrimSpace(strings.Join(strings.Split(head, " ")[1:], " "))
	return s
}

func getBody(body string) string {
	body = strings.TrimSpace(body)
	return body
}

// ReplaceSensitiveWords 替换敏感词
func ReplaceSensitiveWords(text string, replaceWord string) string {
	// 将匹配位置转换为区间
	//hits := ahocorasick.NewStringMatcher(global_gse.SensitiveWords).Match([]byte(text))
	hits := global.AhoCorasick.Match([]byte(text))
	for _, val := range hits {
		oldReplaceWord := global.SensitiveWords[val]
		text = strings.Replace(text, oldReplaceWord, strings.Repeat(replaceWord, len([]rune(oldReplaceWord))), -1)
	}
	return text
}

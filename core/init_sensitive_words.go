package core

import (
	"blogv2/global"
	"bufio"
	_ "embed"
	"github.com/cloudflare/ahocorasick"
	"github.com/sirupsen/logrus"
	"strings"
)

//go:embed files/sensitive_words_lines.txt
var sensitiveWordsFile string

func InitSensitiveWords() (sensitiveWords []string) {
	scanner := bufio.NewScanner(strings.NewReader(sensitiveWordsFile))
	for scanner.Scan() {
		line := scanner.Text()
		sensitiveWords = append(sensitiveWords, line)
	}
	logrus.Info("SensitiveWords加载成功")
	return
}

func InitAhoCorasick() (ahoCorasick *ahocorasick.Matcher) {
	logrus.Info("开始加载,敏感词匹配")
	ahoCorasick = ahocorasick.NewStringMatcher(global.SensitiveWords)
	logrus.Info("敏感词匹配AhoCorasick加载成功")
	return
}

package xss

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

func XSSFilter(content string) (newContent string) {
	// 文章正文防xss注入
	contentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
	if err != nil {
		return
	}
	contentDoc.Find("script").Remove()
	contentDoc.Find("img").Remove()
	contentDoc.Find("iframe").Remove()

	newContent = contentDoc.Text()
	return
}

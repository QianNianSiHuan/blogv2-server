package markdown

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

//提取字符

func ExtractContent(content string, length int) (abs string, err error) {
	// 把markdown转成html，再取文本
	html := MdToHTML(content)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		return
	}
	htmlText := doc.Text()
	abs = htmlText
	if len(htmlText) > length {
		// 如果大于200，就取前200
		abs = string([]rune(htmlText)[:length])
	}
	return
}

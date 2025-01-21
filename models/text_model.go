package models

import _ "embed"

type TextModel struct {
	Model
	ArticleID uint   `json:"articleID"`
	Head      string `json:"head"`
	Body      string `json:"body"`
}

//go:embed mappings/text_mapping.json
var textMapping string

func (TextModel) Mapping() string {
	return textMapping
}
func (TextModel) Index() string {
	return "text_index"
}

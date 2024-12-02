package models

import "fmt"

type ImageModel struct {
	Model
	Filename string `gorm:"64" json:"filename"`
	Path     string `gorm:"256" json:"path"`
	Size     int    `json:"size"`
	Hash     string `gorm:"32" json:"hash"`
}

func (i ImageModel) WebPath() string {
	return fmt.Sprintf("/")
}

package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type ImageModel struct {
	Model
	Filename string `gorm:"64" json:"filename"`
	Path     string `gorm:"256" json:"path"`
	Size     int64  `json:"size"`
	Hash     string `gorm:"32" json:"hashTool"`
}

func (i ImageModel) WebPath() string {
	return fmt.Sprintf("http://localhost:8080" + "/" + i.Path)
}
func (i ImageModel) BeforeDelete(tx *gorm.DB) error {
	err := os.Remove(i.Path)
	if err != nil {
		logrus.Warnf("删除文件失败 %s", err)
	}
	return nil
}

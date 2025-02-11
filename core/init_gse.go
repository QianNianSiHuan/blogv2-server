package core

import (
	"blogv2/global/global_gse"
	"github.com/go-ego/gse"
	"github.com/sirupsen/logrus"
)

func InitGse() {
	newGse, _ := gse.New()
	global_gse.Gse = newGse
	logrus.Info("gse构建完成")
}

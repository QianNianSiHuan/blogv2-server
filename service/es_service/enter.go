package es_service

import (
	"blogv2/artFontFiles"
	"blogv2/global"
	"context"
	"github.com/sirupsen/logrus"
)

func CreatIndexV2(index, mapping string) {
	if ExistsIndex(index) {
		DeleteIndex(index)
	}
	CreatIndex(index, mapping)
}
func CreatIndex(index, mapping string) {
	_, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).Do(context.Background())
	if err != nil {
		artFontFiles.OutPutArtisticFont(artFontFiles.FAIL)
		logrus.Errorf("%s 索引创建失败 %s", index, err)
		return
	}
	logrus.Infof("%s 索引穿件成功", index)
}
func ExistsIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}
func DeleteIndex(index string) {
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		artFontFiles.OutPutArtisticFont(artFontFiles.FAIL)
		logrus.Errorf("%s 索引删除失败 %s", index, err)
		return
	}
	logrus.Infof("%s 索引删除成功", index)
}

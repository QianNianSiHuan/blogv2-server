package es_service

import (
	"blogv2/global"
	"context"
	"github.com/sirupsen/logrus"
)

// 先删除存在的，在创建索引
func CreatIndexV2(index, mapping string) error {
	if ExistsIndex(index) {
		err := DeleteIndex(index)
		if err != nil {
			return err
		}
	}
	err := CreatIndex(index, mapping)
	if err != nil {
		return err
	}
	return nil
}

// 创建索引
func CreatIndex(index, mapping string) error {
	_, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s 索引创建失败 %s", index, err)
		return err
	}
	logrus.Infof("%s 索引创建成功", index)
	return nil
}

// 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}

// 删除索引
func DeleteIndex(index string) error {
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s 索引删除失败 %s", index, err)
		return err
	}
	logrus.Infof("%s 索引删除成功", index)
	return nil
}

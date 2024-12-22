package core

import (
	"blogv2/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	es := global.Config.ES
	if es.Url == "" {
		return nil
	}
	client, err := elastic.NewClient(
		elastic.SetURL("http://"+es.Url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.Username, es.Password),
	)
	if err != nil {
		logrus.Panic(err.Error())
		return nil
	}
	logrus.Info("es连接成功")
	return client
}

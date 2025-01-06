package cron_service

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

func Cron() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	//每天两点同步redis数据
	crontab.Start()
	EntryID, err := crontab.AddFunc("0 0  2 * * *", SyncArticle)
	if err != nil {
		logrus.Errorf("EntryID: %d , err: %s", EntryID, err)
	}
	EntryID, err = crontab.AddFunc("0 0  3 * * *", SyncArticle)
	if err != nil {
		logrus.Errorf("EntryID: %d , err: %s", EntryID, err)
	}
}

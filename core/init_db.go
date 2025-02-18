package core

import (
	"blogv2/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func InitDB() *gorm.DB {
	//ProgressbarMsg <- "数据库初始化..."
	if len(global.Config.DB) == 0 {
		logrus.Fatalf("数据库未配置")
	}
	dc := global.Config.DB[0]
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成外建约束
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败: %s", err)
	}
	sqlDB, err := db.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	logrus.Infof("数据库连接成功")

	if len(global.Config.DB) > 1 {
		var readList []gorm.Dialector
		for _, d := range global.Config.DB[1:] {
			readList = append(readList, mysql.Open(d.DSN()))
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc.DSN())}, //读写
			Replicas: readList,                               //读
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			logrus.Fatalf("读写配置错误 %s", err)
		}
		logrus.Info("数据库读写配置成功")
		return db
	}
	if global.Config.DB[0].Debug {
		return db.Debug()
	}
	return db
}

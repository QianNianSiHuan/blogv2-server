package log_service

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"time"
)

type RuntimeLog struct {
	level           enum.LogLevelType
	title           string
	itemList        []string
	serviceName     string
	runtimeDateType RuntimeDateType
}

func (ac *RuntimeLog) Save() {
	ac.SetNowTime()
	var log models.LogModel
	//判断创建还是更新
	global.DB.Find(&log, fmt.Sprintf("service_name = ? and log_type = ? and created_at >= date_sub(now(),%s)",
		ac.runtimeDateType.GetSqlTime()), ac.serviceName, enum.RuntimeLogType)
	content := strings.Join(ac.itemList, "\n")
	if log.ID != 0 {
		//更新
		c := strings.Join(ac.itemList, "\n")
		newContent := log.Content + "\n" + c
		global.DB.Model(&log).Updates(map[string]any{
			"content": newContent,
		})
		ac.itemList = []string{}
		return
	}
	err := global.DB.Create(&models.LogModel{
		LogType:     enum.RuntimeLogType,
		Title:       ac.title,
		Content:     content,
		Level:       ac.level,
		ServiceName: ac.serviceName,
	}).Error
	if err != nil {
		logrus.Errorf("创建运行日志错误 %s", err)
		return
	}
	ac.itemList = []string{}
}

func (ac *RuntimeLog) SetTitle(title string) {
	ac.title = title
}
func (ac *RuntimeLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}
func (ac *RuntimeLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("连接: %s%s"), label, href)
}
func (ac *RuntimeLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("照片:%s", src))
}
func (ac *RuntimeLog) setItem(label string, value any, logLevelType enum.LogLevelType) {
	var v string
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%s", value)
	}
	ac.itemList = append(ac.itemList, fmt.Sprintf("%s%s%s",
		logLevelType,
		label,
		v,
	))
}
func (ac *RuntimeLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}
func (ac *RuntimeLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}
func (ac *RuntimeLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LofWarnLevel)
}
func (ac *RuntimeLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LofErrLevel)
}
func (ac *RuntimeLog) SetNowTime() {
	ac.itemList = append(ac.itemList, fmt.Sprintf("更新时间:%s", time.Now().Format("2006-01-02 15:04:05")))
}
func (ac *RuntimeLog) SetError(label string, err error) {
	msg := errors.WithStack(err)
	logrus.Errorf(err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("错误: %s%s%s%s", label, err, err, msg))
	fmt.Println(label, err)
}

type RuntimeDateType int8

const (
	RuntimeDateHour  = 1
	RuntimeDateDay   = 2
	RuntimeDateWeek  = 3
	RuntimeDateMonth = 4
)

func (ac RuntimeDateType) GetSqlTime() string {
	switch ac {
	case RuntimeDateHour:
		return "interval 1 hour"
	case RuntimeDateDay:
		return "interval 1 day"
	case RuntimeDateWeek:
		return "interval 1 week"
	case RuntimeDateMonth:
		return "interval 1 month"
	default:
		return "interval 1 day"
	}
}
func NewRuntimeLog(serviceName string, dateType RuntimeDateType) *RuntimeLog {
	return &RuntimeLog{
		serviceName:     serviceName,
		runtimeDateType: dateType,
	}
}

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

// RuntimeLog 结构体用于记录运行时日志。
// 它包含了所有关于运行时日志的信息，如日志级别、标题、日志项列表、服务名称和服务时间类型。
type RuntimeLog struct {
	level           enum.LogLevelType // 日志级别
	title           string            // 日志标题
	itemList        []string          // 日志项列表
	serviceName     string            // 服务名称
	runtimeDateType RuntimeDateType   // 运行日期类型，定义了日志的时间范围
}

// Save 将运行时日志保存到数据库中。
// 如果在指定时间内存在相同的服务名和类型的日志，则更新现有日志；否则创建新的日志条目。
func (ac *RuntimeLog) Save() {
	ac.SetNowTime()
	var log models.LogModel

	// 查找是否存在符合条件的日志（同一服务名、类型，并且在指定时间范围内）
	global.DB.Find(&log, fmt.Sprintf("service_name = ? and log_type = ? and created_at >= date_sub(now(),%s)",
		ac.runtimeDateType.GetSqlTime()), ac.serviceName, enum.RuntimeLogType)
	content := strings.Join(ac.itemList, "\n")

	if log.ID != 0 { // 存在匹配的日志，进行更新操作
		newContent := log.Content + "\n" + content
		global.DB.Model(&log).Updates(map[string]any{
			"content": newContent,
		})
	} else { // 不存在匹配的日志，创建新的日志条目
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
	}
	ac.itemList = []string{} // 清空日志项列表
}

// SetTitle 设置日志标题
func (ac *RuntimeLog) SetTitle(title string) {
	ac.title = title
}

// SetLevel 设置日志级别
func (ac *RuntimeLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

// SetLink 添加一个链接到日志项列表
func (ac *RuntimeLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\">%s</a></div></div>\n",
		label, href, href))
}

// SetImage 添加图片信息到日志项列表
func (ac *RuntimeLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>", src))
}

// setItem 是一个辅助函数，用于设置日志项，并根据日志级别格式化输出。
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
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>",
		logLevelType, label, v))
}

// SetItem 设置信息级别的日志项
func (ac *RuntimeLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}

// SetItemInfo 设置信息级别的日志项
func (ac *RuntimeLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}

// SetItemWarn 设置警告级别的日志项
func (ac *RuntimeLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LofWarnLevel)
}

// SetItemError 设置错误级别的日志项
func (ac *RuntimeLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LofErrLevel)
}

// SetNowTime 向日志项列表添加当前时间戳
func (ac *RuntimeLog) SetNowTime() {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_time\">%s</div>", time.Now().Format("2006-01-02 15:04:05")))
}

// SetError 设置错误日志项，并记录堆栈跟踪
func (ac *RuntimeLog) SetError(label string, err error) {
	msg := errors.WithStack(err)
	logrus.Errorf(err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("\n<div class=\"log_error\"><div class=\"line\"><div class=\"label\">%s</div><div class=\"value\">%s</div><div class=\"type\">%T</div></div><div class=\"stack\">%+v</div></div>",
		label, err, err, msg))
}

// RuntimeDateType 定义了日志的时间范围类型
type RuntimeDateType int8

const (
	RuntimeDateHour  = 1 // 按小时记录日志
	RuntimeDateDay   = 2 // 按天记录日志
	RuntimeDateWeek  = 3 // 按周记录日志
	RuntimeDateMonth = 4 // 按月记录日志
)

// GetSqlTime 返回对应的时间间隔SQL字符串
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

// NewRuntimeLog 创建一个新的RuntimeLog实例
func NewRuntimeLog(serviceName string, dateType RuntimeDateType) *RuntimeLog {
	return &RuntimeLog{
		serviceName:     serviceName,
		runtimeDateType: dateType,
	}
}

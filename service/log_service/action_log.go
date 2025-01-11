package log_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"reflect"
	"strings"

	"blogv2/core"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
)

// ActionLog 结构体用于记录操作日志。
// 它包含了所有关于日志的信息，如上下文、日志级别、标题、请求和响应体等。
type ActionLog struct {
	c                  *gin.Context      // Gin框架的上下文
	level              enum.LogLevelType // 日志级别
	title              string            // 日志标题
	requestBody        []byte            // 请求体内容
	responseBody       []byte            // 响应体内容
	log                *models.LogModel  // 持久化到数据库的日志模型
	showRequest        bool              // 是否显示请求体
	showResponse       bool              // 是否显示响应体
	itemList           []string          // 日志项列表
	showRequestHeader  bool              // 是否显示请求头
	showResponseHeader bool              // 是否显示响应头
	responseHeader     http.Header       // 响应头信息
	isMiddleware       bool              // 标记是否作为中间件调用
}

// ShowRequest 设置是否显示请求体
func (ac *ActionLog) ShowRequest() {
	ac.showRequest = true
}

// ShowResponse 设置是否显示响应体
func (ac *ActionLog) ShowResponse() {
	ac.showResponse = true
}

// SetTitle 设置日志标题
func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

// setItem 是一个辅助函数，用于设置日志项，并根据日志级别格式化输出。
func (ac *ActionLog) setItem(label string, value any, logLevelType enum.LogLevelType) {
	var v string
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%s", value)
	}
	// 添加日志项到列表中
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>",
		logLevelType, label, v))
}

// SetItem 设置日志项，默认为信息级别
func (ac *ActionLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}

// SetItemInfo 设置信息级别的日志项
func (ac *ActionLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}

// SetItemWarn 设置警告级别的日志项
func (ac *ActionLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LofWarnLevel)
}

// SetItemError 设置错误级别的日志项
func (ac *ActionLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LofErrLevel)
}

// SetLevel 设置日志级别
func (ac *ActionLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

// SetLink 设置链接日志项
func (ac *ActionLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\">%s</a></div></div>\n",
		label, href, href))
}

// SetImage 设置图片日志项
func (ac *ActionLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>", src))
}

// ShowResquestHeader 设置是否显示请求头
func (ac *ActionLog) ShowResquestHeader() {
	ac.showRequestHeader = true
}

// ShowResponseHeader 设置是否显示响应头
func (ac *ActionLog) ShowResponseHeader() {
	ac.showResponseHeader = true
}

// SetRequest 设置请求体，同时处理HTTP请求体只能读取一次的问题
func (ac *ActionLog) SetRequest(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	// 将原始请求体替换为一个新的Reader，以便可以多次读取
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	ac.requestBody = byteData
}

// SetResponse 设置响应体
func (ac *ActionLog) SetResponse(data []byte) {
	ac.responseBody = data
}

// SetResponseHeader 设置响应头
func (ac *ActionLog) SetResponseHeader(header http.Header) {
	ac.responseHeader = header
}

// SetError 设置错误日志项，并记录堆栈跟踪
func (ac *ActionLog) SetError(label string, err error) {
	msg := errors.WithStack(err)
	logrus.Errorf("%s : %s", label, err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("\n<div class=\"log_error\"><div class=\"line\"><div class=\"label\">%s</div><div class=\"value\">%s</div><div class=\"type\">%T</div></div><div class=\"stack\">%+v</div></div>",
		label, err, err, msg))
	fmt.Println(label, err)
}

// MiddlewareSave 在中间件中保存日志
func (ac *ActionLog) MiddlewareSave() {
	_saveLog, _ := ac.c.Get("saveLog")
	saveLog, _ := _saveLog.(bool)
	if !saveLog {
		return
	}
	if ac.log == nil {
		ac.isMiddleware = true
		ac.Save()
		return
	}
	// 添加响应头信息
	if ac.showResponseHeader {
		byteDate, _ := json.Marshal(ac.responseHeader)
		ac.itemList = append(ac.itemList, fmt.Sprintf("响应头: %s", byteDate))
	}
	// 添加响应体信息
	if ac.showResponse {
		ac.itemList = append(ac.itemList, fmt.Sprintf("响应体: %s", ac.responseBody))
	}
	ac.Save()
}

// Save 保存日志到数据库
func (ac *ActionLog) Save() (id uint) {
	if ac.log != nil {
		newContent := strings.Join(ac.itemList, "\n")
		content := ac.log.Content + "\n" + newContent
		global.DB.Model(ac.log).Updates(map[string]any{
			"content": content,
		})
		ac.itemList = []string{}
		return ac.log.ID
	}
	var newItemList []string
	// 添加请求头信息
	if ac.showRequestHeader {
		byteDate, _ := json.Marshal(ac.c.Request.Header)
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_request_header\"><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>", byteDate))
	}
	// 添加请求体信息
	if ac.showRequest {
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_request\"><div class=\"log_request_head\"><span class=\"log_request_method %s\">%s</span><span class=\"log_request_path\">%s</span></div><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>",
			strings.ToLower(ac.c.Request.Method),
			ac.c.Request.Method,
			ac.c.Request.URL,
			string(ac.requestBody),
		))
	}

	newItemList = append(newItemList, ac.itemList...)
	// 如果是中间件调用，则添加响应头和响应体信息
	if ac.isMiddleware {
		if ac.showResponseHeader {
			byteDate, _ := json.Marshal(ac.responseHeader)
			newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_response_header\"><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>", byteDate))
		}
		if ac.showResponse {
			newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_response\"><pre class=\"log_json_body\">%s</pre></div>",
				string(ac.responseBody)))
		}
	}
	ip := ac.c.ClientIP()
	addr := core.GetIpAddr(ip)
	userID := uint(0)
	claims, err := jwts.ParseTokenByGin(ac.c)
	if err == nil && claims != nil {
		userID = claims.UserID
	}

	log := models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.title,
		Content: strings.Join(newItemList, "\n"),
		Level:   ac.level,
		UserID:  userID,
		IP:      ip,
		Addr:    addr,
	}
	err = global.DB.Create(&log).Error
	if err != nil {
		logrus.Errorf("日志创建失败 %s", err)
		return
	}
	ac.log = &log
	ac.itemList = []string{}
	return log.ID
}

// NewActionLogByGin 创建一个新的ActionLog实例
func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}

// GetLog 从Gin上下文中获取ActionLog实例，如果不存在则创建新的
func GetLog(c *gin.Context) *ActionLog {
	_log, ok := c.Get("log")
	if !ok {
		return NewActionLogByGin(c)
	}
	log, ok := _log.(*ActionLog)
	if !ok {
		return NewActionLogByGin(c)
	}
	c.Set("saveLog", true)
	return log
}

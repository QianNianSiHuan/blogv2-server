package log_service

import (
	"blogv2/core"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/unitls/jwt"
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
)

type ActionLog struct {
	c                  *gin.Context
	level              enum.LogLevelType
	title              string
	requestBody        []byte
	responseBody       []byte
	log                *models.LogModel
	showRequest        bool
	showResponse       bool
	itemList           []string
	showRequestHeader  bool
	showResponseHeader bool
	responseHeader     http.Header
	isMiddleware       bool
}

func (ac *ActionLog) ShowRequest() {
	ac.showRequest = true
}
func (ac *ActionLog) ShowResponse() {
	ac.showResponse = true
}
func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

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
	ac.itemList = append(ac.itemList, fmt.Sprintf("%s%s%s",
		logLevelType,
		label,
		v,
	))
}
func (ac *ActionLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}
func (ac *ActionLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LofInfoLevel)
}
func (ac *ActionLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LofWarnLevel)
}
func (ac *ActionLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LofErrLevel)
}
func (ac *ActionLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}
func (ac *ActionLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("连接: %s%s"), label, href)
}
func (ac *ActionLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("照片:%s", src))
}

func (ac *ActionLog) ShowResquestHeader() {
	ac.showRequestHeader = true
}
func (ac *ActionLog) ShowResponseHeader() {
	ac.showResponseHeader = true
}
func (ac *ActionLog) SetRequest(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	//logrus.Info("body: ", string(byteData))
	ac.requestBody = byteData
}
func (ac *ActionLog) SetResponse(data []byte) {
	ac.responseBody = data
}
func (ac *ActionLog) SetResponseHeader(header http.Header) {
	ac.responseHeader = header
}

func (ac *ActionLog) SetError(label string, err error) {
	msg := errors.WithStack(err)
	logrus.Errorf(err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("错误: %s%s%s%s", label, err, err, msg))
	fmt.Println(label, err)
}
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
	//响应头
	if ac.showResponseHeader {
		byteDate, _ := json.Marshal(ac.responseHeader)
		ac.itemList = append(ac.itemList, fmt.Sprintf("响应头: %s", byteDate))
	}
	//设置响应
	if ac.showResponse {
		ac.itemList = append(ac.itemList, fmt.Sprintf("响应体: %s",
			ac.responseBody,
		))
	}
	ac.Save()
}
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
	//设置请求
	if ac.showRequestHeader {
		byteDate, _ := json.Marshal(ac.c.Request.Header)
		newItemList = append(newItemList, fmt.Sprintf("请求头: %s", byteDate))
	}
	if ac.showRequest {
		newItemList = append(newItemList, fmt.Sprintf("请求体: %s%s%s%s",
			strings.ToLower(ac.c.Request.Method),
			ac.c.Request.Method,
			ac.c.Request.URL,
			string(ac.requestBody),
		))
	}

	newItemList = append(newItemList, ac.itemList...)
	//请求头

	if ac.isMiddleware {
		//响应头
		if ac.showResponseHeader {
			byteDate, _ := json.Marshal(ac.responseHeader)
			newItemList = append(newItemList, fmt.Sprintf("响应头 %s", byteDate))
		}
		//设置响应
		if ac.showResponse {
			newItemList = append(newItemList, fmt.Sprintf("响应体 %s",
				ac.responseBody,
			))
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
func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}
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

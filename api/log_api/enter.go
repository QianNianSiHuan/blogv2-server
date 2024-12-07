package log_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}
type LogListRequest struct {
	common.PageInfo
	LogType     enum.LogType      `form:"logType" `
	Level       enum.LogLevelType `form:"level"`
	UserID      uint              `form:"userID"`
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"`
	ServiceName string            `form:"serviceName"`
}
type LogListResponse struct {
	models.LogModel
	UserNickName string `json:"userNickName"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	//精确查询,模糊匹配
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	list, count, err := common.ListQuery(models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, common.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickName: logModel.UserModel.Nickname,
			UserAvatar:   logModel.UserModel.Avatar,
		})
	}
	res.SuccessWithList(c, _list, count)
	return
}

// 改变日志状态为已读
func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var log models.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg(c, "不存在的日志")
		return
	}
	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)
	}
	res.FailWithMsg(c, "日志读取成功")
}

// 日志删除
func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var logList []models.LogModel
	global.DB.Find(&logList, "id in ?", cr.IDList)
	if len(logList) > 0 {
		global.DB.Delete(&logList)
	}
	msg := fmt.Sprintf("日志删除成功，共删除 %d 条", len(logList))
	res.SuccessWithMsg(c, msg)
}

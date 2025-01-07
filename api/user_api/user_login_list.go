package user_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserLoginListRequest struct {
	common.PageInfo
	UserID   uint   `form:"userID"`
	IP       string `form:"ip"`
	Addr     string `form:"addr"`
	StarTime string `form:"startTime"` //年月日时分秒
	EndTime  string `form:"endTime"`
	Type     int8   `form:"type" binding:"required,oneof=1 2"`
}
type UserLoginListResponse struct {
	models.UserLoginModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (UserApi) UserLoginListView(c *gin.Context) {
	var cr UserLoginListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	if cr.Type == 1 {
		cr.UserID = claims.UserID
	}
	var query = global.DB.Where("")
	if cr.StarTime != "" {
		_, err = time.Parse("2006-01-02 15:04:05", cr.StarTime)
		if err != nil {
			res.FailWithMsg(c, "开始时间解析出错")
			return
		}
		query.Where("created_at > ?", cr.StarTime)
	}
	if cr.EndTime != "" {
		_, err = time.Parse("2006-01-02 15:04:05", cr.EndTime)
		if err != nil {
			res.FailWithMsg(c, "结束时间解析出错")
			return
		}
		query.Where("created_at < ?", cr.EndTime)
	}
	var preloads []string
	if cr.Type == 2 {
		preloads = []string{"UserModel"}
	}
	_list, count, _ := common.ListQuery(models.UserLoginModel{
		UserID: cr.UserID,
		IP:     cr.IP,
		Addr:   cr.Addr,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: preloads,
	})
	var list = make([]UserLoginListResponse, 0)
	for _, model := range _list {
		list = append(list, UserLoginListResponse{
			UserLoginModel: model,
			UserNickname:   model.UserModel.Nickname,
			UserAvatar:     model.UserModel.Avatar,
		})
	}
	res.SuccessWithList(c, list, count)
}

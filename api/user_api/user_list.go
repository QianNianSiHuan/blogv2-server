package user_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserListRequest struct {
	common.PageInfo
}

func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	_list, count, _ := common.ListQuery(models.UserModel{}, common.Options{
		PageInfo:     cr.PageInfo,
		DefaultOrder: "created_at desc",
	})

	list := make([]models.UserModel, 0)

	for _, v := range _list {
		list = append(list, v)
	}
	logrus.Info(_list)
	res.SuccessWithList(c, list, count)

}

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	claims := jwts.GetClaims(c)
	if claims.Role != enum.AdminRole {
		res.FailWithMsg(c, "权限不足")
		return
	}
	var userList []models.UserModel
	global.DB.Find(&userList, "id in ?", cr.IDList)
	if len(userList) > 0 {
		global.DB.Delete(&userList)
	}
	msg := fmt.Sprintf("用户删除成功，共删除 %d 个用户", len(userList))
	res.SuccessWithMsg(c, msg)
}

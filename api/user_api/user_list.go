package user_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/models"
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

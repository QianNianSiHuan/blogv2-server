package image_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImageApi struct {
}
type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"web_path"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	var cr common.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	_list, count, err := common.ListQuery(models.ImageModel{}, common.Options{
		PageInfo: cr,
		Likes:    []string{"fileTool"},
		Preloads: nil,
	})
	var list = make([]ImageListResponse, 0)
	for _, model := range _list {
		list = append(list, ImageListResponse{
			ImageModel: model,
			WebPath:    model.WebPath(),
		})
	}
	res.SuccessWithList(c, list, count)
}
func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var list []models.ImageModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	//删除对应文件
	var successCount, errCount int64
	if len(list) > 0 {
		successCount = global.DB.Delete(&list).RowsAffected
	}

	errCount = int64(len(list)) - successCount

	msg := fmt.Sprintf("操作成功，成功%d 失败%d", successCount, errCount)

	res.SuccessWithMsg(c, msg)
}

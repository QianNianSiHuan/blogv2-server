package banner_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BannerApi struct {
}
type BannerCreatRequest struct {
	Cover string `json:"cover" binding:"required"`
	Href  string `json:"href"`
	Show  bool   `json:"show"`
	Type  int8   `json:"type" binding:"required,oneof=1 2"`
}

func (BannerApi) BannerCreatView(c *gin.Context) {
	var cr BannerCreatRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if cr.Type == 0 {
		cr.Type = 1
	}
	err = global.DB.Create(&models.BannerModel{
		Show:  cr.Show,
		Cover: cr.Cover,
		Href:  cr.Href,
		Type:  cr.Type,
	}).Error
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "添加banner成功")
}

type BannerListRequest struct {
	common.PageInfo
	Show bool `from:"show"`
	Type int8 `from:"type"`
}

func (BannerApi) BannerListView(c *gin.Context) {
	var cr BannerListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	list, count, _ := common.ListQuery(models.BannerModel{
		Show: cr.Show,
		Type: cr.Type,
	}, common.Options{
		PageInfo: cr.PageInfo,
	})
	res.SuccessWithList(c, list, count)
}
func (BannerApi) BannerRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	var successCount, errCount int64
	if len(list) > 0 {
		successCount = global.DB.Delete(&list).RowsAffected
	}

	errCount = int64(len(list)) - successCount

	msg := fmt.Sprintf("操作成功，成功%d 失败%d", successCount, errCount)

	res.SuccessWithMsg(c, msg)
}

func (BannerApi) BannerUpdateView(c *gin.Context) {
	var id models.IDRequest
	err := c.ShouldBindUri(&id)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var cr BannerCreatRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var model models.BannerModel
	err = global.DB.Take(&model, id.ID).Error
	if err != nil {
		res.FailWithMsg(c, "不存在的banner")
		return
	}
	err = global.DB.Model(&model).Updates(map[string]any{
		"cover": cr.Cover,
		"href":  cr.Href,
		"show":  cr.Show,
		"type":  cr.Type,
	}).Error
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithMsg(c, "更新成功")
}

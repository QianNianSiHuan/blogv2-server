package data_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type GrowthDataRequest struct {
	Type int8 `form:"type" binding:"required,oneof=1 2 3"`
}

type GrowthDataResponse struct {
	GrowthRate int      `json:"growthRate"` // 增长率
	GrowthNum  int      `json:"growthNum"`  // 增长数
	CountList  []int    `json:"countList"`
	DateList   []string `json:"dateList"`
}
type Table struct {
	Date  string `gorm:"column:date"`
	Count int    `gorm:"column:count"`
}

func (DataApi) GrowthDataView(c *gin.Context) {
	var cr GrowthDataRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	now := time.Now()
	before7 := now.AddDate(0, 0, -7)
	var dataList []Table

	switch cr.Type {
	case 1:
		global.DB.Model(models.SiteFlowModel{}).Where("created_at >= ? and created_at <= ?",
			before7.Format("2006-01-02")+" 00:00:00",
			now.Format("2006-01-02 15:04:05"),
		).
			Select("date(created_at) as date", "sum(count) as count").
			Group("date").Scan(&dataList)
	case 2:
		global.DB.Model(models.ArticleModel{}).Where("created_at >= ? and created_at <= ? and status = ?",
			before7.Format("2006-01-02")+" 00:00:00",
			now.Format("2006-01-02 15:04:05"),
			enum.ArticleStatusPublished).
			Select("date(created_at) as date", "count(id) as count").
			Group("date").Scan(&dataList)
	case 3:
		global.DB.Model(models.UserModel{}).Where("created_at >= ? and created_at <= ?",
			before7.Format("2006-01-02")+" 00:00:00",
			now.Format("2006-01-02 15:04:05"),
		).
			Select("date(created_at) as date", "count(id) as count").
			Group("date").Scan(&dataList)
	}
	var dateMap = map[string]int{}
	for _, model := range dataList {
		date := strings.Split(model.Date, "T")[0]
		dateMap[date] = model.Count
	}

	response := GrowthDataResponse{}
	for i := 0; i < 7; i++ {
		date := before7.AddDate(0, 0, i+1)
		dateS := date.Format("2006-01-02")
		count, _ := dateMap[dateS]
		response.CountList = append(response.CountList, count)
		response.DateList = append(response.DateList, dateS)
	}
	// 算增长，找最后一个和最后一个的前一个
	response.GrowthNum = response.CountList[6] - response.CountList[5]
	if response.CountList[5] == 0 {
		response.GrowthRate = 100
	} else {
		response.GrowthRate = int(float64(response.GrowthNum) / float64(response.CountList[5]) * 100)
	}
	res.SuccessWithData(c, response)
}

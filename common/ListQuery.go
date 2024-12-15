package common

import (
	"blogv2/global"
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Order string `form:"order"`
}

func (p PageInfo) GetPage() int {
	if p.Page > 20 || p.Page <= 0 {
		return 1
	}
	return p.Page
}
func (p PageInfo) GetLimit() int {
	if p.Limit <= 0 || p.Limit > 50 {
		return 10
	}
	return p.Limit
}
func (p PageInfo) GetOffset() int {
	offset := (p.GetPage() - 1) * p.GetLimit()
	return offset
}

type Options struct {
	PageInfo     PageInfo
	Likes        []string
	Preloads     []string
	Where        *gorm.DB
	DefaultOrder string
}

func ListQuery[T any](model T, option Options) (list []T, count int64, err error) {
	//基础查询
	query := global.DB.Model(model).Where(model)
	//模糊匹配
	if len(option.Likes) > 0 && option.PageInfo.Key != "" {
		likes := global.DB.Where("")
		for _, column := range option.Likes {
			likes.Or(
				fmt.Sprintf("%s like ?", column),
				fmt.Sprintf("%%%s%%", option.PageInfo.Key))
		}
		query = query.Where(likes)
	}
	//定制化高级查询
	if option.Where != nil {
		query.Where(option.Where)
	}
	//预加载
	for _, preload := range option.Preloads {
		query = query.Preload(preload)
	}
	//总数
	query.Count(&count)

	//分页
	limit := option.PageInfo.GetLimit()
	Offset := option.PageInfo.GetOffset()
	//排序
	if option.PageInfo.Order != "" {
		query.Order(option.PageInfo.Order)
	} else {
		if option.DefaultOrder != "" {
			query = query.Order(option.DefaultOrder)
		}
	}
	err = query.Offset(Offset).Limit(limit).Find(&list).Error
	return
}

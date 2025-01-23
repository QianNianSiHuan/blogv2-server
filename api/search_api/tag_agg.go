package search_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"sort"
)

type TagAggResponse struct {
	Tag          string `json:"tag"`
	ArticleCount int    `json:"articleCount"`
}

func (SearchApi) TagAggView(c *gin.Context) {
	var cr common.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var list = make([]TagAggResponse, 0)
	if global.ESClient == nil {
		var articleList []models.ArticleModel
		global.DB.Find(&articleList, "tag_list <> ''")
		var tagMap = map[string]int{}
		for _, model := range articleList {
			for _, tag := range model.TagList {
				count, ok := tagMap[tag]
				if !ok {
					tagMap[tag] = 1
					continue
				}
				tagMap[tag] = count + 1
			}
		}
		for tag, count := range tagMap {
			list = append(list, TagAggResponse{
				Tag:          tag,
				ArticleCount: count,
			})
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].ArticleCount > list[j].ArticleCount
		})
		res.SuccessWithList(c, list, int64(len(list)))
		return
	}

	agg := elastic.NewTermsAggregation().Field("tag_list")
	agg.SubAggregation("page",
		elastic.NewBucketSortAggregation().
			From(cr.GetOffset()).
			Size(cr.Limit))
	query := elastic.NewBoolQuery()
	query.MustNot(elastic.NewTermQuery("tag_list", ""))
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Aggregation("tags1", elastic.NewCardinalityAggregation().Field("tag_list")).
		Size(0).Do(context.Background())
	if err != nil {
		logrus.Errorf("查询失败 %s", err)
		res.FailWithMsg(c, "查询失败")
		return
	}
	var t AggType
	var val = result.Aggregations["tags"]
	err = json.Unmarshal(val, &t)
	if err != nil {
		logrus.Errorf("解析json失败 %s %s", err, string(val))
		res.FailWithMsg(c, "查询失败")
		return
	}
	var co Agg1Type
	err = json.Unmarshal(result.Aggregations["tags1"], &co)
	if err != nil {
		logrus.Errorf("解析json失败 %s %s", err, string(val))
		res.FailWithMsg(c, "查询失败")
		return
	}
	for _, bucket := range t.Buckets {
		list = append(list, TagAggResponse{
			Tag:          bucket.Key,
			ArticleCount: bucket.DocCount,
		})
	}
	res.SuccessWithList(c, list, int64(co.Value))
	return
}

type AggType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
	} `json:"buckets"`
}
type Agg1Type struct {
	Value int `json:"value"`
}

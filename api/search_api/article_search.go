package search_api

import (
	"blogv2/common"
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	jwts "blogv2/utils/jwt"
	"blogv2/utils/sql"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"slices"
	"strconv"
)

type ArticleSearchRequest struct {
	common.PageInfo
	Tag  string `form:"tag"`
	Type int8   `form:"type"` // 0 猜你喜欢  1 最新发布  2最多回复 3最多点赞 4最多收藏
}

type ArticleBaseInfo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}

type ArticleListResponse struct {
	models.ArticleModel
	AdminTop      bool    `json:"adminTop"` // 是否是管理员置顶
	CategoryTitle *string `json:"categoryTitle"`
	UserNickname  string  `json:"userNickname"`
	UserAvatar    string  `json:"userAvatar"`
}

func (SearchApi) ArticleSearchView(c *gin.Context) {
	var cr ArticleSearchRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var sortMap = map[int8]string{
		0: "_score",
		1: "created_at",
		2: "comment_count",
		3: "digg_count",
		4: "collect_count",
	}
	sortKey := sortMap[cr.Type]
	if sortKey == "" {
		res.FailWithMsg(c, "搜索类型错误")
		return
	}
	topArticleIDList := getAdminTopArticleIDList()
	collectMap := redis_article.GetAllCacheCollect(0)
	diggMap := redis_article.GetAllCacheDigg(0)
	lookMap := redis_article.GetAllCacheLook(0)
	commentMap := redis_article.GetAllCacheComment(0)
	if global.Redis != nil {
		var tagsMap map[string][]string
		var tagArticleList []string
		var sortArticleList []string
		var articleList []string
		where := global.DB.Where("")
		where.Where("status = ?", enum.ArticleStatusPublished)
		if cr.Tag != "" {
			tagsMap = redis_article.GetTagAggAll()
			tagArticleList = tagsMap[cr.Tag]
		}
		var articleTopMap = map[uint]bool{}
		for _, u := range topArticleIDList {
			articleTopMap[u] = true
		}
		switch cr.Type {
		case 0:
			sortArticleList = redis_article.GetAllCacheAllSort()
		case 1:
			where.Order("created_at desc")
		case 2:
			sortArticleList = redis_article.GetAllCacheLookSort()
		case 3:
			sortArticleList = redis_article.GetAllCacheDiggSort()
		case 4:
			sortArticleList = redis_article.GetAllCacheCollectSort()
		}
		//tag加指定排序
		if len(tagArticleList) > 0 && len(sortArticleList) > 0 {
			slices.Reverse(sortArticleList)
			for _, sortArticle := range sortArticleList {
				for _, tagArticle := range tagArticleList {
					if sortArticle == tagArticle {
						articleList = append(articleList, tagArticle)
					}
				}
			}
			where.Where("id in ?", articleList)
		}
		//有tag,没有排序
		if len(sortArticleList) == 0 && len(tagArticleList) > 0 {
			slices.Reverse(tagArticleList)
			articleList = tagArticleList
			where.Where("id in ?", articleList)
		}
		//没有tag,有排序
		if len(tagArticleList) == 0 && len(sortArticleList) > 0 {
			articleList = sortArticleList
			where.Where("id in ?", articleList)
		}
		if cr.Type != 2 {
			var _article []uint
			for _, article := range articleList {
				id, _ := strconv.Atoi(article)
				_article = append(_article, uint(id))
			}
			topArticleIDList = append(topArticleIDList, _article...)
		}

		_list, count, _ := common.ListQuery(models.ArticleModel{}, common.Options{
			Preloads:     []string{"CategoryModel", "UserModel"},
			PageInfo:     cr.PageInfo,
			Likes:        []string{"title", "abstract"},
			DefaultOrder: sql.ConvertSliceOrderSql(topArticleIDList),
			Where:        where,
		})
		var list = make([]ArticleListResponse, 0)
		for _, model := range _list {
			model.Content = ""
			model.DiggCount = model.DiggCount + diggMap[model.ID]
			model.CollectCount = model.CollectCount + collectMap[model.ID]
			model.LookCount = model.LookCount + lookMap[model.ID]
			model.CommentCount = model.CommentCount + commentMap[model.ID]
			item := ArticleListResponse{
				ArticleModel: model,
				AdminTop:     articleTopMap[model.ID],
				UserNickname: model.UserModel.Nickname,
				UserAvatar:   model.UserModel.Avatar,
			}
			if model.CategoryModel != nil {
				item.CategoryTitle = &model.CategoryModel.Title
			}
			list = append(list, item)
		}
		res.SuccessWithList(c, list, count)
		return
	}
	if global.ESClient == nil {
		// 服务降级，用户可能没有配置es
		where := global.DB.Where("")
		if cr.Tag != "" {
			where.Where("tag_list like ?", fmt.Sprintf("%%%s%%", cr.Tag))
		}
		var articleTopMap = map[uint]bool{}
		for _, u := range topArticleIDList {
			articleTopMap[u] = true
		}
		sortMap = map[int8]string{
			0: "",
			1: "created_at desc",
			2: "comment_count desc",
			3: "digg_count desc",
			4: "collect_count desc",
		}
		sort, _ := sortMap[cr.Type]
		cr.Order = sort
		_list, count, _ := common.ListQuery(models.ArticleModel{}, common.Options{
			Preloads:     []string{"CategoryModel", "UserModel"},
			PageInfo:     cr.PageInfo,
			Likes:        []string{"title", "abstract"},
			DefaultOrder: sql.ConvertSliceOrderSql(topArticleIDList),
			Where:        where,
		})
		var list = make([]ArticleListResponse, 0)
		for _, model := range _list {
			model.Content = ""
			model.DiggCount = model.DiggCount + diggMap[model.ID]
			model.CollectCount = model.CollectCount + collectMap[model.ID]
			model.LookCount = model.LookCount + lookMap[model.ID]
			model.CommentCount = model.CommentCount + commentMap[model.ID]
			item := ArticleListResponse{
				ArticleModel: model,
				AdminTop:     articleTopMap[model.ID],
				UserNickname: model.UserModel.Nickname,
				UserAvatar:   model.UserModel.Avatar,
			}
			if model.CategoryModel != nil {
				item.CategoryTitle = &model.CategoryModel.Title
			}
			list = append(list, item)
		}

		res.SuccessWithList(c, list, count)
		return
	}
	query := elastic.NewBoolQuery()
	if cr.Key != "" {
		query.Should(
			elastic.NewMatchQuery("title", cr.Key),
			elastic.NewMatchQuery("abstract", cr.Key),
			elastic.NewMatchQuery("content", cr.Key),
		)
	}
	if cr.Tag != "" {
		query.Must(
			elastic.NewTermQuery("tag_list", cr.Tag),
		)
	}

	// 只能查发布的文章
	query.Must(elastic.NewTermQuery("status", 3))

	var articleIDList []uint

	// 把管理员置顶的文章查出来

	var articleTopMap = map[uint]bool{}
	if len(topArticleIDList) > 0 {
		var topArticleIDListAny []interface{}
		for _, u := range topArticleIDList {
			topArticleIDListAny = append(topArticleIDListAny, u)
			articleTopMap[u] = true
			articleIDList = append(articleIDList, u)
		}
		query.Should(elastic.NewTermsQuery("id", topArticleIDListAny...).Boost(10))
	}

	if cr.Type == 0 {
		// 只有猜你喜欢，才会把用户喜欢的标签带入查询
		claims, err := jwts.ParseTokenByGin(c)
		if err == nil && claims != nil {
			// 用户登录了
			// 查用户感兴趣的分类
			var userConf models.UserConfModel
			err = global.DB.Take(&userConf, "user_id = ?", claims.UserID).Error
			if err != nil {
				res.FailWithMsg(c, "用户配置不存在")
				return
			}
			if len(userConf.LikeTags) > 0 {
				tagQuery := elastic.NewBoolQuery()
				var tagAnyList []interface{}
				for _, tag := range userConf.LikeTags {
					tagAnyList = append(tagAnyList, tag)
				}
				tagQuery.Should(elastic.NewTermsQuery("tag_list", tagAnyList...))
				query.Must(tagQuery)
			}
		}
	}

	highlight := elastic.NewHighlight()
	highlight.Field("title")
	highlight.Field("abstract")

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Highlight(highlight).
		From(cr.GetOffset()).
		Size(cr.GetLimit()).
		Sort(sortKey, false).
		Do(context.Background())
	if err != nil {
		source, _ := query.Source()
		byteData, _ := json.Marshal(source)
		logrus.Errorf("查询失败 %s \n %s", err, string(byteData))
		res.FailWithMsg(c, "查询失败")
		return
	}

	count := result.Hits.TotalHits.Value
	var searchArticleMap = map[uint]ArticleBaseInfo{}

	for _, hit := range result.Hits.Hits {

		var art ArticleBaseInfo
		err = json.Unmarshal(hit.Source, &art)
		if err != nil {
			logrus.Warnf("解析失败 %s  %s", err, string(hit.Source))
			continue
		}
		if hit.Score != nil {
			fmt.Println(*hit.Score, art.Title, art.ID)
		}
		if len(hit.Highlight["title"]) > 0 {
			art.Title = hit.Highlight["title"][0]
		}
		if len(hit.Highlight["abstract"]) > 0 {
			art.Abstract = hit.Highlight["abstract"][0]
		}

		searchArticleMap[art.ID] = art
		articleIDList = append(articleIDList, art.ID)
	}

	where := global.DB.Where("id in ?", articleIDList)

	_list, _, _ := common.ListQuery(models.ArticleModel{}, common.Options{
		Where:        where,
		Preloads:     []string{"CategoryModel", "UserModel"},
		DefaultOrder: sql.ConvertSliceOrderSql(articleIDList),
	})

	var list = make([]ArticleListResponse, 0)
	for _, model := range _list {
		model.Content = ""
		model.DiggCount = model.DiggCount + diggMap[model.ID]
		model.CollectCount = model.CollectCount + collectMap[model.ID]
		model.LookCount = model.LookCount + lookMap[model.ID]
		model.CommentCount = model.CommentCount + commentMap[model.ID]
		item := ArticleListResponse{
			ArticleModel: model,
			AdminTop:     articleTopMap[model.ID],
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
		}
		if model.CategoryModel != nil {
			item.CategoryTitle = &model.CategoryModel.Title
		}
		item.Title = searchArticleMap[model.ID].Title
		item.Abstract = searchArticleMap[model.ID].Abstract
		list = append(list, item)
	}
	res.SuccessWithList(c, list, count)
}

func getAdminTopArticleIDList() (topArticleIDList []uint) {
	var userIDList []uint
	global.DB.Model(models.UserModel{}).Where("role = ?", enum.AdminRole).Select("id").Scan(&userIDList)
	global.DB.Model(models.UserTopArticleModel{}).Where("user_id in ?", userIDList).Select("article_id").Scan(&topArticleIDList)
	return
}

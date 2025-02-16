package redis_article

import (
	"blogv2/global"
	"github.com/go-redis/redis"
	"strconv"
)

type ArticleSortType string

const (
	articleCommentSort ArticleSortType = "article_comment_sort"
	articleDiggSort    ArticleSortType = "article_digg_sort"
	articleCollectSort ArticleSortType = "article_collect_sort"
	articleLookSort    ArticleSortType = "article_look_sort"
)

// set 修改指定类型的文章计数（增加或减少）。
func setArticleSort(t ArticleSortType, articleID uint, n int) {
	global.Redis.ZIncrBy(string(t), float64(n), strconv.Itoa(int(articleID)))
}

// SetCacheLookSort  设置文章的查看次数。
func SetCacheLookSort(articleID uint, n int) {
	setArticleSort(articleLookSort, articleID, n)
}

// SetCacheDiggSort  设置文章的点赞数。
func SetCacheDiggSort(articleID uint, n int) {
	setArticleSort(articleDiggSort, articleID, n)
}

// SetCacheCollectSort  设置文章的收藏数。
func SetCacheCollectSort(articleID uint, n int) {
	setArticleSort(articleCollectSort, articleID, n)
}
func SetCacheCommentSort(articleID uint, n int) {
	setArticleSort(articleCommentSort, articleID, n)
}

func setArticleSortByCount(t ArticleSortType, articleID uint, count int) {
	global.Redis.ZAdd(string(t), redis.Z{
		Score:  float64(count),
		Member: articleID,
	})
}

// SetCacheLookSortByCount 设置文章的查看次数。
func SetCacheLookSortByCount(articleID uint, count int) {
	setArticleSortByCount(articleLookSort, articleID, count)
}

// SetCacheDiggSortByCount 设置文章的点赞数。
func SetCacheDiggSortByCount(articleID uint, count int) {
	setArticleSortByCount(articleDiggSort, articleID, count)
}

// SetCacheCollectSortByCount 设置文章的收藏数。
func SetCacheCollectSortByCount(articleID uint, count int) {
	setArticleSortByCount(articleCollectSort, articleID, count)
}
func SetCacheCommentSortByCount(articleID uint, count int) {
	setArticleSortByCount(articleCommentSort, articleID, count)
}

// get 获取指定类型的文章计数。
func getArticleSort(t ArticleSortType, articleID uint) int {
	num, _ := global.Redis.HGet(string(t), strconv.Itoa(int(articleID))).Int()
	return num
}

// GetCacheLookSort  获取文章的查看次数。
func GetCacheLookSort(articleID uint) int {
	return getArticleSort(articleLookSort, articleID)
}

// GetCacheDiggSort 获取文章的点赞数。
func GetCacheDiggSort(articleID uint) int {
	return getArticleSort(articleDiggSort, articleID)
}

// GetCacheCollectSort 获取文章的收藏数。
func GetCacheCollectSort(articleID uint) int {
	return getArticleSort(articleCollectSort, articleID)
}
func GetCacheCommentSort(articleID uint) int {
	return getArticleSort(articleCommentSort, articleID)
}

// GetAllSort  获取所有文章的指定类型计数。
func GetAllSort(t ArticleSortType) (IDSortList []redis.Z) {
	IDList, _ := global.Redis.ZRangeByScoreWithScores(string(t), redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  0,
	}).Result()
	return IDList
}

// GetAllCacheLookSort 获取所有文章的查看次数。
func GetAllCacheLookSort() []redis.Z {
	return GetAllSort(articleLookSort)
}

// GetAllCacheDiggSort  获取所有文章的点赞数。
func GetAllCacheDiggSort() []redis.Z {
	return GetAllSort(articleDiggSort)
}

// GetAllCacheCollectSort  获取所有文章的收藏数。
func GetAllCacheCollectSort() []redis.Z {
	return GetAllSort(articleCollectSort)
}
func GetAllCacheCommentSort() []redis.Z {
	return GetAllSort(articleCommentSort)
}

func ClearArticleSortByID(articleID uint) {
	global.Redis.HDel(string(articleCommentSort), strconv.Itoa(int(articleID)))
	global.Redis.HDel(string(articleDiggSort), strconv.Itoa(int(articleID)))
	global.Redis.HDel(string(articleCollectSort), strconv.Itoa(int(articleID)))
	global.Redis.HDel(string(articleLookSort), strconv.Itoa(int(articleID)))
}

func ArticleSortClear() {
	global.Redis.Del(string(articleCommentSort), string(articleDiggSort), string(articleCollectSort), string(articleLookSort))
}

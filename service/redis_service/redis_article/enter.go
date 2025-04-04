package redis_article

import (
	"blogv2/global"
	"blogv2/utils/date"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type articleCacheType string

const (
	articleCacheLook    articleCacheType = "article_look_key"
	articleCacheDigg    articleCacheType = "article_digg_key"
	articleCacheCollect articleCacheType = "article_collect_key"
	articleCacheComment articleCacheType = "article_comment_key"
)

// set 修改指定类型的文章计数（增加或减少）。
func set(t articleCacheType, articleID uint, n int) {
	nowDate := time.Now().Format("20060102")
	global.Redis.HIncrBy(string(t)+fmt.Sprintf("_%s", nowDate), strconv.Itoa(int(articleID)), int64(n))
}

// SetCacheLook 设置文章的查看次数。
func SetCacheLook(articleID uint, n int) {
	set(articleCacheLook, articleID, n)
}

// SetCacheDigg 设置文章的点赞数。
func SetCacheDigg(articleID uint, n int) {
	set(articleCacheDigg, articleID, n)
}

// SetCacheCollect 设置文章的收藏数。
func SetCacheCollect(articleID uint, n int) {
	set(articleCacheCollect, articleID, n)
}
func SetCacheComment(articleID uint, n int) {
	set(articleCacheComment, articleID, n)
}

// get 获取指定类型的文章计数。
func get(t articleCacheType, articleID uint, day int) int {
	nowDate := time.Now().Add(time.Duration(-day*24) * time.Hour).Format("20060102")
	num, _ := global.Redis.HGet(string(t)+fmt.Sprintf("_%s", nowDate), strconv.Itoa(int(articleID))).Int()
	return num
}

// GetCacheLook 获取文章的查看次数。
func GetCacheLook(articleID uint, day int) int {
	return get(articleCacheLook, articleID, day)
}

// GetCacheDigg 获取文章的点赞数。
func GetCacheDigg(articleID uint, day int) int {
	return get(articleCacheDigg, articleID, day)
}

// GetCacheCollect 获取文章的收藏数。
func GetCacheCollect(articleID uint, day int) int {
	return get(articleCacheCollect, articleID, day)
}
func GetCacheComment(articleID uint, day int) int {
	return get(articleCacheComment, articleID, day)
}

// getAll 获取所有文章的指定类型计数。
func getAll(t articleCacheType, day int) (mps map[uint]int) {
	nowDate := time.Now().Add(time.Duration(-day*24) * time.Hour).Format("20060102")
	res, err := global.Redis.HGetAll(string(t) + fmt.Sprintf("_%s", nowDate)).Result()
	if err != nil {
		return
	}
	mps = make(map[uint]int)
	for key, numS := range res {
		iK, _ := strconv.Atoi(key)
		iN, _ := strconv.Atoi(numS)
		mps[uint(iK)] = iN
	}
	return mps
}

// GetAllCacheLook 获取所有文章的查看次数。
func GetAllCacheLook(day int) (mps map[uint]int) {
	return getAll(articleCacheLook, day)
}

// GetAllCacheDigg 获取所有文章的点赞数。
func GetAllCacheDigg(day int) (mps map[uint]int) {
	return getAll(articleCacheDigg, day)
}

// GetAllCacheCollect 获取所有文章的收藏数。
func GetAllCacheCollect(day int) (mps map[uint]int) {
	return getAll(articleCacheCollect, day)
}
func GetAllCacheComment(day int) (mps map[uint]int) {
	return getAll(articleCacheComment, day)
}

func SetUserArticleHistoryCache(articleID, userID uint) {
	// 创建一个基于用户ID的Redis键名，用来存储用户的浏览历史
	key := fmt.Sprintf("history_%d", userID)

	// 创建一个字段名，表示特定的文章ID
	field := fmt.Sprintf("%d", articleID)

	// 获取当前时间
	now := time.Now()

	// 通过 date.GetNowAfter() 函数获取一个将来的时间点（可能是文章过期时间）
	endTime := date.GetNowAfter()

	// 计算当前时间和结束时间之间的差值
	subTime := endTime.Sub(now)

	// 将字段设置到Redis中，并设置其存活时间为subTime，如果出错则记录错误日志并返回
	err := global.Redis.Set(key, field, subTime).Err()
	if err != nil {
		logrus.Error(err)
		return
	}

	// 设置Redis键的过期时间为endTime指定的时间点，如果出错则记录错误日志并返回
	err = global.Redis.ExpireAt(key, endTime).Err()
	if err != nil {
		logrus.Error(err)
		return
	}
}

func GetUserArticleHistoryCache(articleID, userID uint) bool {
	// 创建一个基于用户ID的Redis键名，用来存储用户的浏览历史
	key := fmt.Sprintf("history_%d", userID)

	// 创建一个字段名，表示特定的文章ID
	field := fmt.Sprintf("%d", articleID)

	// 尝试从Redis的哈希表中获取指定键和字段的值
	err := global.Redis.HGet(key, field).Err()

	// 如果获取过程中发生错误（比如键或字段不存在），则返回false
	if err != nil {
		return false
	}

	// 如果没有发生错误，说明找到了对应的浏览历史记录，返回true
	return true
}

func Clear() {
	err := global.Redis.Del("article_look_key", "article_digg_key", "article_collect_key").Err()
	if err != nil {
		logrus.Error(err)
	}
}

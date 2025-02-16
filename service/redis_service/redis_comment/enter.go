package redis_comment

import (
	"blogv2/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type commentCacheType string

const (
	commentCacheApply commentCacheType = "comment_apply_key"
	commentCacheDigg  commentCacheType = "comment_digg_key"
)

func set(t commentCacheType, commentID uint, n int) {
	nowTime := time.Now().Format("20060102")
	global.Redis.HIncrBy(string(t)+fmt.Sprintf("_%s", nowTime), strconv.Itoa(int(commentID)), int64(n))
}
func SetCacheApply(commentID uint, n int) {
	set(commentCacheApply, commentID, n)
}
func SetCacheDigg(commentID uint, n int) {
	set(commentCacheDigg, commentID, n)
}

func get(t commentCacheType, commentID uint) int {
	nowTime := time.Now().Format("20060102")
	num, _ := global.Redis.HGet(string(t)+fmt.Sprintf("_%s", nowTime), strconv.Itoa(int(commentID))).Int()
	return num
}
func GetCacheApply(commentID uint) int {
	return get(commentCacheApply, commentID)
}
func GetCacheDigg(commentID uint) int {
	return get(commentCacheDigg, commentID)
}

func GetAll(t commentCacheType) (mps map[uint]int) {
	YesterdayTime := time.Now().Add(-24 * time.Hour).Format("20060102")
	res, err := global.Redis.HGetAll(string(t) + fmt.Sprintf("_%s", YesterdayTime)).Result()
	if err != nil {
		return
	}
	mps = make(map[uint]int)
	for key, numS := range res {
		iK, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		iN, err := strconv.Atoi(numS)
		if err != nil {
			continue
		}
		mps[uint(iK)] = iN
	}
	return mps
}

func GetAllCacheApply() (mps map[uint]int) {
	return GetAll(commentCacheApply)
}
func GetAllCacheDigg() (mps map[uint]int) {
	return GetAll(commentCacheDigg)
}

func Clear() {
	YesterdayTime := time.Now().Add(-24 * time.Hour).Format("20060102")
	err := global.Redis.Del("comment_apply_key"+fmt.Sprintf("_%s", YesterdayTime),
		"comment_digg_key"+fmt.Sprintf("_%s", YesterdayTime)).Err()
	if err != nil {
		logrus.Error(err)
	}
}

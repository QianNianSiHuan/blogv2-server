package redis_article

import (
	"blogv2/global"
	"fmt"
	"github.com/siddontang/go/log"
)

func SetTagAgg(tag string, articleID uint) {
	global.Redis.SAdd(fmt.Sprintf("tag_%s", tag), articleID)
}

func SetTagAggAdd(tag string) {
	global.Redis.SAdd("tagadd", tag)
}
func getTagAggAdd() []string {
	tags, _ := global.Redis.SMembers("tagadd").Result()
	return tags
}

func GetTagAggAllCount() (mp map[string]int) {
	tags := getTagAggAdd()
	mp = make(map[string]int)
	for _, tag := range tags {
		_count, _ := global.Redis.SCard(fmt.Sprintf("tag_%s", tag)).Result()
		if _count == 0 {
			RemoveTagAggAdd(tag)
		}
		mp[tag] = int(_count)
	}
	return
}

func GetTagAggAll() (mp map[string][]string) {
	tags := getTagAggAdd()
	mp = make(map[string][]string)
	for _, tag := range tags {
		articleList, err := global.Redis.SMembers(fmt.Sprintf("tag_%s", tag)).Result()
		if err != nil {
			log.Error(err)
			continue
		}
		mp[tag] = articleList
	}
	return
}
func RemoveTagAgg(articleID uint, tags ...string) {
	for _, tag := range tags {
		global.Redis.SRem(fmt.Sprintf("tag_%s", tag), articleID)
	}
}

func RemoveTagAggAdd(tags ...string) {
	global.Redis.SRem("tagadd", tags)
}

func ClearAllTagAgg() {
	tags := getTagAggAdd()
	for _, tag := range tags {
		global.Redis.Del(fmt.Sprintf("tag_%s", tag))
	}
}

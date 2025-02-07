package global

import (
	"blogv2/conf"
	"github.com/cloudflare/ahocorasick"
	"github.com/go-redis/redis"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/mojocn/base64Captcha"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"sync"
)

const Version = "10.0.1"

var (
	Config           *conf.Config
	DB               *gorm.DB
	Redis            *redis.Client
	CaptchaStore     = base64Captcha.DefaultMemStore
	EmailVerifyStore = sync.Map{}
	ESClient         *elastic.Client
	IP               *xdb.Searcher
	SensitiveWords   []string //敏感词切片
	AhoCorasick      *ahocorasick.Matcher
)

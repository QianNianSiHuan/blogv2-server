package global

import (
	"blogv2/conf"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
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
)

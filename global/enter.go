package global

import (
	"blogv2/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const Version = "10.0.1"

var (
	Config *conf.Config
	DB     *gorm.DB
	Redis  *redis.Client
)

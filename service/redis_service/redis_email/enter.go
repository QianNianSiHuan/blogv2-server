package redis_email

import (
	"blogv2/global"
	"fmt"
	"time"
)

type EmailSendType string

const (
	EmailFeedBackIDCache EmailSendType = "email_feedBack_id"
)

func SetEmailFeedBackID(id string) {
	global.Redis.Set(string(EmailFeedBackIDCache)+fmt.Sprintf("_%s", id), true, 60*time.Second)
}
func GetEmailFeedBackID(id string) (val int) {
	val, _ = global.Redis.Get(string(EmailFeedBackIDCache) + fmt.Sprintf("_%s", id)).Int()
	return
}

func IsExistExEmailFeedBackID(id string) bool {
	val := GetEmailFeedBackID(id)
	if val != 1 {
		return false
	}
	return true
}

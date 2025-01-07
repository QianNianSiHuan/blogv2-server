package email_store

import (
	"blogv2/global"
)

type EmailStoreInfo struct {
	Email string
	Code  string
}

func Verify(emailID, emailCode string) (info EmailStoreInfo, ok bool) {
	value, ok := global.EmailVerifyStore.Load(emailID)
	if !ok {
		return
	}
	info, ok = value.(EmailStoreInfo)
	if !ok {
		return
	}
	if info.Code != emailCode {
		global.EmailVerifyStore.Delete(emailID)
		return
	}
	global.EmailVerifyStore.Delete(emailID)
	return info, true
}

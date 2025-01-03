package date

import "time"

// 获取当天结束时间
func GetNowAfter() time.Time {
	now := time.Now()
	location := time.Local
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)
	return endTime
}

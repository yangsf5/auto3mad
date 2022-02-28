package daily

import (
	"time"
)

type retDay struct {
	PassedDay int `json:"passed_day"`
	Weekend   int `json:"weekend"`
	Holiday   int `json:"holiday"`
	Adapter   int `json:"adapter"`
	Ret       int `json:"ret"`
}

func GetDDUCount(firstDate string) int {
	diffDay := getDayPassed(firstDate)

	// 计算周末天数
	weekends := diffDay / 7 * 2

	// 国家节假日，不算周末，则每年 11 天假
	// 即每 365 天，有 11 天假
	holidays := diffDay * 11 / 365

	// 每年给自己 10 天适配时间
	adapter := diffDay * 10 / 365

	d := diffDay - weekends - holidays - adapter

	return int(d)
}

func GetWeekPassed(firstDate string) int {
	return int(getDayPassed(firstDate) / 7)
}

func getDayPassed(firstDate string) float64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	start, _ := time.ParseInLocation("2006-01-02", firstDate, loc)
	now := time.Now().In(loc) // 这里本地时间不太有必要，但为了防止放到服务器时，时区不一致问题

	diff := now.Sub(start)
	return diff.Hours() / 24
}

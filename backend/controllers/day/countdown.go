package day

import (
	"time"

	"backend/controllers/base"
)

type CountdownController struct {
	base.BaseController
}

type retDay struct {
	PassedDay int `json:"passed_day"`
	Weekend   int `json:"weekend"`
	Holiday   int `json:"holiday"`
	Adapter   int `json:"adapter"`
	Ret       int `json:"ret"`
}

func (c *CountdownController) Get() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	start, _ := time.ParseInLocation("2006-01-02", "2021-02-25", loc)
	now := time.Now().In(loc) // 这里本地时间不太有必要，但为了防止放到服务器时，时区不一致问题

	diff := now.Sub(start)
	diffDay := diff.Hours() / 24

	// 计算周末天数
	// 2021-02-25 是周四
	// 即第一周有四天，有两个工作日，这个单独算，否则后续偏移不准确
	// 所以先从 diffDay 里减掉第一周四天，最后再加第一周的周末两天
	weekends := (diffDay-4)/7*2 + 2

	// 国家节假日，不算周末，则每年 11 天假
	// 即每 365 天，有 11 天假
	holidays := diffDay * 11 / 365

	// 每年给自己 10 天适配时间
	adapter := diffDay * 10 / 365

	d := diffDay - weekends - holidays - adapter

	c.JSONOK(map[string]interface{}{
		"data": retDay{
			int(diffDay),
			int(weekends),
			int(holidays),
			int(adapter),
			int(d),
		},
	})
}

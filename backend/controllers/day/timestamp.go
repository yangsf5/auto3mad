package day

import (
	"time"

	"backend/controllers/base"
)

type TimestampController struct {
	base.BaseController
}

type retTime struct {
	Area     string `json:"area"`
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
	offset   int
}

func (c *TimestampController) Get() {
	ts, err := c.GetInt64("timestamp")
	c.JSONErrorAbort(err)

	rets := []retTime{
		{"China", "UTC+8", "", 8 * 60 * 60},
		{"UTC-0", "UTC-0", "", 0 * 60 * 60},
	}

	for i, ret := range rets {
		loc := time.FixedZone(ret.Timezone, ret.offset)
		t := time.Unix(ts, 0).In(loc)
		ret.Time = t.Format("2006-01-02 15:04:05")
		rets[i] = ret
	}

	c.JSONOK(rets)
}

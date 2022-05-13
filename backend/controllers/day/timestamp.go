package day

import (
	"time"

	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/util"
)

type TimestampController struct {
	base.Controller
}

type retD2T struct {
	FirstSecond int64 `json:"first_second"`
	LastSecond  int64 `json:"last_second"`
}

type retT2D struct {
	Area     string `json:"area"`
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
	offset   int
}

func (c *TimestampController) Get() {
	t := c.GetString("type")
	if t == "d2t" {
		date := c.GetString("date")
		f, l := util.GetDateTimestamp(date)
		c.JSONOK(retD2T{f, l})
	} else if t == "t2d" {
		c.getT2D()
	}
	c.JSONErrorAbort("type must be d2t or t2d.")
}

func (c *TimestampController) getT2D() {
	ts, err := c.GetInt64("timestamp")
	c.JSONErrorAbort(err)

	rets := []retT2D{
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

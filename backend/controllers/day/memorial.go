package day

import (
	"backend/controllers/base"
	"backend/models/db"
	"time"
)

type MemorialController struct {
	base.BaseController
}

type retMemorial struct {
	Desc       string `json:"desc"`
	Date       string `json:"date"`
	Passed     int    `json:"passed"`      // 过去了多少天
	NextLeft   int    `json:"next_left"`   // 离下次还剩多少天
	NextDate   string `json:"next_date"`   // 下次日期
	CycleCount int    `json:"cycle_count"` // 周期数
}

type retEdit struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
	Date string `json:"date"`
}

func (c *MemorialController) Get() {
	kind := c.GetString("kind")

	days, err := (&db.ModelDayMemorial{}).GetAllDays()
	c.JSONErrorAbort(err, err != nil)

	if kind == "full" {
		rets := make([]retMemorial, 0)
		for _, day := range days {
			passedDayCount, cycleCount, nextDate, nextLeft := c.calcDate(day.Date, day.RemindType)
			item := retMemorial{}
			item.Desc = day.Desc
			item.Date = day.Date
			item.Passed = passedDayCount
			item.NextLeft = nextLeft
			item.NextDate = nextDate
			item.CycleCount = cycleCount
			rets = append(rets, item)
		}
		c.JSONOK(rets)
	} else if kind == "edit" {
		rets := make([]retEdit, 0)
		for _, day := range days {
			item := retEdit{}
			item.ID = day.ID
			item.Desc = day.Desc
			item.Date = day.Date
			rets = append(rets, item)
		}
		c.JSONOK(rets)
	}
}

func (c *MemorialController) calcDate(memorialDate string, remindType int) (passedDayCount int, cycleCount int, nextDate string, nextLeft int) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", memorialDate, loc)
	now := time.Now()
	passedDayCount = diffDay(now, t)

	memorialYear := t.Year()
	thisYear := now.Year()

	if remindType == 1 {
		tNext := t.AddDate(thisYear-memorialYear, 0, 0)

		cycleCount = thisYear - memorialYear
		if tNext.Unix() > now.Unix() {
			cycleCount -= 1
		}

		if tNext.Unix() < now.Unix() {
			tNext = tNext.AddDate(1, 0, 0)
		}
		nextDate = tNext.Format("2006-01-02")
		nextLeft = diffDay(tNext, now)
	}

	return
}

func diffDay(bigT, litteT time.Time) int {
	diff := (bigT.Unix() - litteT.Unix()) / (24 * 60 * 60)
	return int(diff)
}

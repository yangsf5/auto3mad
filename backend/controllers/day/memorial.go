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
	Desc           string
	Date           string
	Passed         int    // 过去了多少天
	NextLeft       int    // 离下次还剩多少天
	NextDate       string // 下次日期
	RemindTypeDesc string // 纪念类型
	CycleCount     int    // 周期数
}

func (c *MemorialController) Get() {
	remindTypeDescs := map[int]string{
		1: "每年",
		2: "每季",
		3: "每月",
	}

	rets := make([]retMemorial, 0)
	days, _ := (&db.ModelDayMemorial{}).GetAllDays()
	for _, day := range days {
		passedDayCount, cycleCount, nextDate, nextLeft := c.calcDate(day.Date, day.RemindType)
		item := retMemorial{}
		item.Desc = day.Desc
		item.Date = day.Date
		item.Passed = passedDayCount
		item.NextLeft = nextLeft
		item.NextDate = nextDate
		item.RemindTypeDesc = remindTypeDescs[day.RemindType]
		item.CycleCount = cycleCount
		rets = append(rets, item)
	}

	c.JSONOK(rets)
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

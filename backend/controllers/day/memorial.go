package day

import (
	"encoding/json"
	"time"

	"backend/controllers/base"
	"backend/models/db/day"
)

type MemorialController struct {
	base.BaseController
	modelMemo day.MemorialModel
}

func (c *MemorialController) Prepare() {
	c.modelMemo = *day.NewMemorialModel()
	c.BaseController.Prepare()
}

type retMemorial struct {
	Desc       string `json:"desc"`
	Date       string `json:"date"`
	Passed     int    `json:"passed"`      // 过去了多少天
	NextLeft   int    `json:"next_left"`   // 离下次还剩多少天
	NextDate   string `json:"next_date"`   // 下次日期
	CycleCount int    `json:"cycle_count"` // 周期数
}

type EditInfo struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
	Date string `json:"date"`
}

func (c *MemorialController) Get() {
	kind := c.GetString("kind")

	var days []day.Memorial
	err := c.modelMemo.GetAllOrderBy(&days, "date")
	c.JSONErrorAbort(err)

	if kind == "full" {
		rets := make([]retMemorial, 0)
		for _, day := range days {
			passedDayCount, cycleCount, nextDate, nextLeft := c.calcDate(day.Date)
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
		rets := make([]EditInfo, 0)
		for _, day := range days {
			item := EditInfo{}
			item.ID = day.ID
			item.Desc = day.Desc
			item.Date = day.Date
			rets = append(rets, item)
		}
		c.JSONOK(rets)
	} else {
		c.JSONErrorAbort("kind must be 'full' or 'edit'")
	}
}

func (c *MemorialController) calcDate(memorialDate string) (passedDayCount int, cycleCount int, nextDate string, nextLeft int) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", memorialDate, loc)
	now := time.Now()
	passedDayCount = diffDay(now, t)

	memorialYear := t.Year()
	thisYear := now.Year()

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

	return
}

func diffDay(bigT, litteT time.Time) int {
	diff := (bigT.Unix() - litteT.Unix()) / (24 * 60 * 60)
	return int(diff)
}

func (c *MemorialController) Post() {
	info := EditInfo{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	c.JSONErrorAbort(err)

	rmd := &day.Memorial{
		ID:   info.ID,
		Desc: info.Desc,
		Date: info.Date,
	}
	err = c.modelMemo.Upsert(rmd)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func (c *MemorialController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.modelMemo.DeleteByID(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

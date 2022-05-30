package daily

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/daily"
)

type RoutineController struct {
	base.Controller
	mr daily.RoutineModel
	me daily.EventModel
}

func (c *RoutineController) Prepare() {
	c.mr = *daily.NewRoutineModel()
	c.me = *daily.NewEventModel()
	c.Controller.Prepare()
}

type routineInfo struct {
	daily.Routine
	TodaySpend     int     `json:"today_spend"`
	WeekWillSpend  int     `json:"week_will_spend"`
	WeekSpend      int     `json:"week_spend"`
	MonthWillSpend int     `json:"month_will_spend"`
	MonthSpend     int     `json:"month_spend"`
	TotalWillSpend float64 `json:"total_will_spend"`
	TotalSpend     float64 `json:"total_spend"`
	WeekPassed     int     `json:"week_passed"`
}

func (c *RoutineController) Get() {
	date := c.GetString("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var rrs []daily.Routine
	err := c.mr.GetAllOrderBy(&rrs, "sort")
	c.JSONErrorAbort(err)

	todaySpends, err := c.me.GetTodaySpendGroupByRoutine(date)
	c.JSONErrorAbort(err)

	weekSpends, err := c.me.GetWeekSpendGroupByRoutine(date)
	c.JSONErrorAbort(err)

	monthSpends, err := c.me.GetMonthSpendGroupByRoutine(date)
	c.JSONErrorAbort(err)

	totalSpends, err := c.me.GetTotalSpendGroupByRoutine()
	c.JSONErrorAbort(err)

	rets := []routineInfo{}

	for _, rr := range rrs {
		ret := routineInfo{Routine: rr}

		strID := strconv.Itoa(rr.ID)
		if v, ok := todaySpends[strID]; ok {
			ret.TodaySpend, err = strconv.Atoi(v.(string))
			c.JSONErrorAbort(err)
		}

		ret.WeekWillSpend = 5 * rr.WillSpend // nolint

		if v, ok := weekSpends[strID]; ok {
			ret.WeekSpend, err = strconv.Atoi(v.(string))
			c.JSONErrorAbort(err)
		}

		ret.MonthWillSpend = 22 * rr.WillSpend // nolint

		if v, ok := monthSpends[strID]; ok {
			ret.MonthSpend, err = strconv.Atoi(v.(string))
			c.JSONErrorAbort(err)
		}

		tws := float64(GetDDUCount(rr.StartDate)*rr.WillSpend) / 60 // nolint
		stws := fmt.Sprintf("%0.1f", tws)
		ret.TotalWillSpend, err = strconv.ParseFloat(stws, 64) // nolint
		c.JSONErrorAbort(err)

		if v, ok := totalSpends[strID]; ok {
			ret.TotalSpend, err = strconv.ParseFloat(v.(string), 64) // nolint
			c.JSONErrorAbort(err)
		}

		ret.WeekPassed = GetWeekPassed(rr.StartDate)

		rets = append(rets, ret)
	}

	c.JSONOK(rets)
}

func (c *RoutineController) Post() {
	info := daily.Routine{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)	
	c.JSONErrorAbort(err)
    if info.ID == -1 {
		info.ID = 0
	}

	err = c.mr.Upsert(&info)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func (c *RoutineController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.mr.DeleteByID(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

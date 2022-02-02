package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
	"strconv"
	"time"
)

type RoutineController struct {
	base.BaseController
	mr daily.RoutineModel
	me daily.EventModel
}

func (c *RoutineController) Prepare() {
	c.mr = *daily.NewRoutineModel()
	c.me = *daily.NewEventModel()
	c.BaseController.Prepare()
}

type routineInfo struct {
	daily.Routine
	TodaySpend int `json:"today_spend"`
	TotalSpend int `json:"total_spend"`
}

func (c *RoutineController) Get() {
	date := c.GetString("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var rrs []daily.Routine
	err := c.mr.GetAll(&rrs)
	c.JSONErrorAbort(err)

	todaySpends, err := c.me.GetTodaySpendGroupByRoutine(date)
	c.JSONErrorAbort(err)

	totalSpends, err := c.me.GetTotalSpendGroupByRoutine()
	c.JSONErrorAbort(err)

	rets := []routineInfo{}
	for _, rr := range rrs {
		ret := routineInfo{Routine: rr}

		strID := strconv.Itoa(rr.ID)
		if todaySpends[strID] != nil {
			ret.TodaySpend, err = strconv.Atoi(todaySpends[strID].(string))
			c.JSONErrorAbort(err)
		}

		if totalSpends[strID] != nil {
			ret.TotalSpend, err = strconv.Atoi(totalSpends[strID].(string))
			c.JSONErrorAbort(err)
		}

		rets = append(rets, ret)
	}

	c.JSONOK(rets)
}

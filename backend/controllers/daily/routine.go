package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
	"encoding/json"
	"fmt"
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
	TodaySpend     int     `json:"today_spend"`
	TotalWillSpend float64 `json:"total_will_spend"`
	TotalSpend     float64 `json:"total_spend"`
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

		if v, ok := totalSpends[strID]; ok {
			ret.TotalSpend, err = strconv.ParseFloat(v.(string), 64)
			c.JSONErrorAbort(err)
		}

		tws := float64(GetDDUCount(rr.StartDate)*rr.WillSpend) / 60
		stws := fmt.Sprintf("%0.1f", tws)
		ret.TotalWillSpend, err = strconv.ParseFloat(stws, 64)
		c.JSONErrorAbort(err)

		rets = append(rets, ret)
	}

	c.JSONOK(rets)
}

func (c *RoutineController) Post() {
	info := daily.Routine{}
	println(string(c.Ctx.Input.RequestBody))
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	c.JSONErrorAbort(err)

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

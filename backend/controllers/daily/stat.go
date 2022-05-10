package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
	"strconv"
	"time"
)

type StatController struct {
	base.BaseController
	mr *daily.RoutineModel
	me *daily.EventModel
}

func (c *StatController) Prepare() {
	c.mr = daily.NewRoutineModel()
	c.me = daily.NewEventModel()
	c.BaseController.Prepare()
}

type statInfo struct {
	Routine string `json:"routine"`
	Month   string `json:"month"`
	Spend   int    `json:"spend"`
}

func (c *StatController) Get() {
	firstMonth := c.GetString("first_month")
	lastMonth := c.GetString("last_month")

	if firstMonth == "" || lastMonth == "" {
		firstMonth = time.Now().Format("2006-01")
		lastMonth = firstMonth
	}

	var rrs []daily.Routine
	err := c.mr.GetAllOrderBy(&rrs, "sort")
	c.JSONErrorAbort(err)

	spends, err := c.me.GetPeriodMonthsSpendGroupByRoutine(firstMonth, lastMonth)
	c.JSONErrorAbort(err)

	rets := []statInfo{}
	for _, spend := range spends {
		ret := statInfo{
			Routine: strconv.Itoa(spend.RoutineID),
			Month:   spend.Month,
			Spend:   spend.Spend,
		}

		rets = append(rets, ret)
	}

	c.JSONOK(rets)
}

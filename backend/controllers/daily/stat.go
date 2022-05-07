package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
	"time"
)

type StatController struct {
	base.BaseController
	me *daily.EventModel
}

func (c *StatController) Prepare() {
	c.me = daily.NewEventModel()
	c.BaseController.Prepare()
}

func (c *StatController) Get() {
	firstMonth := c.GetString("first_month")
	lastMonth := c.GetString("last_month")

	if firstMonth == "" || lastMonth == "" {
		firstMonth = time.Now().Format("2006-01")
		lastMonth = firstMonth
	}

	spends, err := c.me.GetPeriodMonthsSpendGroupByRoutine(firstMonth, lastMonth)
	c.JSONErrorAbort(err)

	c.JSONOK(spends)
}

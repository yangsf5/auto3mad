package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
	"time"
)

type EventController struct {
	base.BaseController
	mr daily.RoutineModel
	me daily.EventModel
}

func (c *EventController) Prepare() {
	c.mr = *daily.NewRoutineModel()
	c.me = *daily.NewEventModel()
	c.BaseController.Prepare()
}

func (c *EventController) Get() {
	date := c.GetString("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	var rs []daily.Event
	err := c.me.GetEventByDate(date, &rs)
	c.JSONErrorAbort(err)

	c.JSONOK(rs)
}

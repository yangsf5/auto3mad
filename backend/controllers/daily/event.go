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

type retEvent struct {
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	SpecificEvent string `json:"specific_event"`
	RoutineId     int    `json:"routine_id"`
	Date          string `json:"date"`
	Spend         int    `json:"spend"`
}

func (c *EventController) Get() {
	date := c.GetString("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	var es []daily.Event
	err := c.me.GetEventByDate(date, &es)
	c.JSONErrorAbort(err)

	rets := []retEvent{}
	for _, e := range es {
		re := retEvent{
			StartTime:     time.Unix(e.StartTime, 0).Format("15:04"),
			EndTime:       time.Unix(e.EndTime, 0).Format("15:04"),
			SpecificEvent: e.SpecificEvent,
			RoutineId:     e.RoutineId,
			Date:          date,
			Spend:         int(e.EndTime-e.StartTime) / 60,
		}
		rets = append(rets, re)
	}

	c.JSONOK(rets)
}

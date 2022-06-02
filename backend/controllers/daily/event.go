package daily

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/daily"
)

type EventController struct {
	base.Controller
	mr daily.RoutineModel
	me daily.EventModel
}

func (c *EventController) Prepare() {
	c.Controller.Prepare()
	c.mr = *daily.NewRoutineModel()
	c.me = *daily.NewEventModel(c.GetMyUserID())
}

type retEvent struct {
	editEventInfo
	Spend int `json:"spend"`
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
		edit := editEventInfo{
			ID:        e.ID,
			StartTime: time.Unix(e.StartTime, 0).Format("15:04"),
			EndTime:   time.Unix(e.EndTime, 0).Format("15:04"),
			RoutineID: e.RoutineID,
			Date:      e.Date,
		}
		re := retEvent{
			editEventInfo: edit,
			Spend:         int(e.EndTime-e.StartTime) / 60, // nolint
		}
		rets = append(rets, re)
	}

	c.JSONOK(map[string]interface{}{
		"events": rets,
	})
}

type editEventInfo struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	RoutineID int    `json:"routine_id"`
}

func (c *EventController) Post() {
	info := editEventInfo{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	c.JSONErrorAbort(err)

	re := &daily.Event{
		ID:        info.ID,
		UserID:    c.GetMyUserID(),
		StartTime: parseTime(info.Date, info.StartTime),
		EndTime:   parseTime(info.Date, info.EndTime),
		RoutineID: info.RoutineID,
		Date:      info.Date,
		Month:     info.Date[0:7],
	}
	err = c.me.Upsert(re)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func parseTime(date, hm string) int64 {
	ft := fmt.Sprintf("%s %s:00", date, hm)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", ft, loc)

	return t.Unix()
}

func (c *EventController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.me.Delete(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

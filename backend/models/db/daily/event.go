package daily

import (
	"backend/models/db/base"
	"backend/models/util"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type Event struct {
	StartTime     int64  `orm:"pk" json:"start_time"`
	EndTime       int64  `json:"end_time"`
	SpecificEvent string `json:"specific_event"`
	RoutineId     int    `json:"routine_id"`
}

func (o *Event) TableName() string {
	return "daily_time_use"
}

func (o *Event) GetID() int {
	return int(o.StartTime)
}

func (o *Event) NewObjectOnlyID(id int) interface{} {
	ooid := new(Event)
	ooid.StartTime = int64(id)
	return ooid
}

type EventModel struct {
	base.BaseModel
}

func NewEventModel() *EventModel {
	m := new(EventModel)
	m.BaseModel = *base.NewBaseModel(&Event{})
	return m
}

func (m *EventModel) GetEventByDate(date string, events *[]Event) error {
	firstSecond, lastSecond := util.GetDateTimestamp(date)
	sql := fmt.Sprintf("SELECT start_time, end_time, specific_event, routine_id FROM %s WHERE start_time BETWEEN %d AND %d ORDER BY start_time DESC", m.TableName, firstSecond, lastSecond)
	_, err := m.ORM.Raw(sql).QueryRows(events)
	return err
}

func (m *EventModel) GetTodaySpendGroupByRoutine(date string) (spends orm.Params, err error) {
	firstSecond, lastSecond := util.GetDateTimestamp(date)
	return m.getSpendGroupByRoutine(firstSecond, lastSecond)
}

func (m *EventModel) GetWeekSpendGroupByRoutine(date string) (spends orm.Params, err error) {
	firstSecond, lastSecond := util.GetWeekTimestamp(date)
	return m.getSpendGroupByRoutine(firstSecond, lastSecond)
}

func (m *EventModel) getSpendGroupByRoutine(firstSecond, lastSecond int64) (spends orm.Params, err error) {
	sql := fmt.Sprintf("SELECT routine_id, ROUND(SUM(end_time-start_time)/60) as `spend` FROM %s WHERE start_time BETWEEN %d AND %d GROUP BY routine_id", m.TableName, firstSecond, lastSecond)
	_, err = m.ORM.Raw(sql).RowsToMap(&spends, "routine_id", "spend")
	return
}

func (m *EventModel) GetTotalSpendGroupByRoutine() (spends orm.Params, err error) {
	sql := fmt.Sprintf("SELECT routine_id, ROUND(SUM(end_time-start_time)/3600, 1) as `spend` FROM %s GROUP BY routine_id", m.TableName)
	_, err = m.ORM.Raw(sql).RowsToMap(&spends, "routine_id", "spend")
	return
}

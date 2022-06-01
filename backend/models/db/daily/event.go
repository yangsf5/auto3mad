package daily

import (
	"fmt"

	"github.com/yangsf5/auto3mad/backend/models/db/base"
	"github.com/yangsf5/auto3mad/backend/models/util"

	"github.com/beego/beego/v2/client/orm"
)

type Event struct {
	StartTime int64 `orm:"pk"`
	EndTime   int64
	RoutineID int `orm:"column(routine_id)"`
	Date      string
	Month     string
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
	base.Model
}

func NewEventModel() *EventModel {
	m := new(EventModel)
	m.Model = *base.NewModel(&Event{})

	return m
}

func (m *EventModel) GetEventByDate(date string, events *[]Event) error {
	sql := fmt.Sprintf("SELECT start_time, end_time, routine_id, date FROM %s WHERE date='%s' ORDER BY start_time DESC",
		m.TableName, date)
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

func (m *EventModel) GetMonthSpendGroupByRoutine(date string) (spends orm.Params, err error) {
	firstSecond, lastSecond := util.GetMonthTimestamp(date)
	return m.getSpendGroupByRoutine(firstSecond, lastSecond)
}

func (m *EventModel) getSpendGroupByRoutine(firstSecond, lastSecond int64) (spends orm.Params, err error) {
	sql := fmt.Sprintf(
		`SELECT routine_id, ROUND(SUM(end_time-start_time)/60) as spend 
			FROM %s 
			WHERE start_time BETWEEN %d AND %d GROUP BY routine_id`,
		m.TableName, firstSecond, lastSecond)
	_, err = m.ORM.Raw(sql).RowsToMap(&spends, "routine_id", "spend")

	return
}

func (m *EventModel) GetTotalSpendGroupByRoutine() (spends orm.Params, err error) {
	sql := fmt.Sprintf(
		"SELECT routine_id, ROUND(SUM(end_time-start_time)/3600, 1) as `spend` FROM %s GROUP BY routine_id",
		m.TableName)
	_, err = m.ORM.Raw(sql).RowsToMap(&spends, "routine_id", "spend")

	return
}

func (m *EventModel) GetRoutineSpendGroupByMonth(
	routineID int, firstMonth, lastMonth string,
) (spends orm.Params, err error) {
	sql := fmt.Sprintf(
		`SELECT month, ROUND(SUM(end_time-start_time)/60) as spend 
			FROM %s 
			WHERE routine_id='%d' AND month BETWEEN '%s' AND '%s' 
			GROUP BY month
			ORDER BY month`,
		m.TableName, routineID, firstMonth, lastMonth)
	_, err = m.ORM.Raw(sql).RowsToMap(&spends, "month", "spend")

	return
}

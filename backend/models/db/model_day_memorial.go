package db

import (
	"fmt"
)

const TABLE_NAME_DAY_MEMORIAL = "day_memorial"

type ModelDayMemorial struct {
}

type MemorailDay struct {
	ID         int `orm:"column(id)"`
	Date       string
	Desc       string
	RemindType int
}

func (m *ModelDayMemorial) GetAllDays() (days []MemorailDay, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY date", TABLE_NAME_DAY_MEMORIAL)
	_, err = getOrm().Raw(sql).QueryRows(&days)
	return
}

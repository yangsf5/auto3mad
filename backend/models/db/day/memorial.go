package day

import (
	"fmt"

	"backend/models/db/base"

	"github.com/beego/beego/v2/client/orm"
)

const DB_TABLE_DAY_MEMORIAL = "day_memorial"

func init() {
	orm.RegisterModel(new(MemorailDay))
}

type ModelDayMemorial struct {
}

type MemorailDay struct {
	ID   int `orm:"column(id)"`
	Date string
	Desc string
}

func (i *MemorailDay) TableName() string {
	return DB_TABLE_DAY_MEMORIAL
}

func (m *ModelDayMemorial) GetAllDays() (days []MemorailDay, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY date", DB_TABLE_DAY_MEMORIAL)
	_, err = base.GetOrm().Raw(sql).QueryRows(&days)
	return
}

func (m *ModelDayMemorial) Upsert(item MemorailDay) (err error) {
	key := MemorailDay{
		ID: item.ID,
	}

	o := base.GetOrm()
	err = o.Read(&key)
	if err == orm.ErrNoRows {
		_, err = o.Insert(&item)
	} else if err != nil {
		return
	} else {
		_, err = o.Update(&item)
	}
	return
}

func (m *ModelDayMemorial) Delete(id int) (err error) {
	_, err = base.GetOrm().Delete(&MemorailDay{ID: id})
	return
}

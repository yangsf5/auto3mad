package day

import (
	"fmt"

	"backend/models/db/base"

	"github.com/beego/beego/v2/client/orm"
)

const DB_TABLE_DAY_MEMORIAL = "day_memorial"

func init() {
	orm.RegisterModel(new(Memorial))
}

type MemorialModel struct {
	base.BaseModel
}

func NewMemorialModel() *MemorialModel {
	m := new(MemorialModel)
	m.BaseModel = *base.NewBaseModel(DB_TABLE_DAY_MEMORIAL, &Memorial{})
	return m
}

type Memorial struct {
	ID   int `orm:"column(id)"`
	Date string
	Desc string
}

func (o *Memorial) TableName() string {
	return DB_TABLE_DAY_MEMORIAL
}

func (o *Memorial) GetID() int {
	return o.ID
}

func (o *Memorial) NewObjectOnlyID(id int) interface{} {
	ooid := new(Memorial)
	ooid.ID = id
	return ooid
}

func (m *MemorialModel) GetAllDays() (days []Memorial, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY date", DB_TABLE_DAY_MEMORIAL)
	_, err = m.ORM.Raw(sql).QueryRows(&days)
	return
}

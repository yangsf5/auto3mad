package day

import (
	"fmt"

	"backend/models/db/base"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(Memorial))
}

type Memorial struct {
	ID   int `orm:"column(id)"`
	Date string
	Desc string
}

func (o *Memorial) TableName() string {
	return "day_memorial"
}

func (o *Memorial) GetID() int {
	return o.ID
}

func (o *Memorial) NewObjectOnlyID(id int) interface{} {
	ooid := new(Memorial)
	ooid.ID = id
	return ooid
}

type MemorialModel struct {
	base.BaseModel
}

func NewMemorialModel() *MemorialModel {
	m := new(MemorialModel)
	o := &Memorial{}
	m.BaseModel = *base.NewBaseModel(o.TableName(), o)
	return m
}

func (m *MemorialModel) GetAllDays() (days []Memorial, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY date", m.TableName)
	_, err = m.ORM.Raw(sql).QueryRows(&days)
	return
}

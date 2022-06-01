package day

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

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
	base.Model
}

func NewMemorialModel() *MemorialModel {
	m := new(MemorialModel)
	m.Model = *base.NewModel(&Memorial{})

	return m
}

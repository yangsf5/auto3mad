package daily

import (
	"backend/models/db/base"
)

type Routine struct {
	ID         int    `orm:"column(id)" json:"id"`
	Icon       string `json:"icon"`
	ShortName  string `json:"short_name"`
	EventScope string `json:"event_scope"`
	WillSpend  int    `json:"will_spend"`
	Sort       int
}

func (o *Routine) TableName() string {
	return "daily_time_routine"
}

func (o *Routine) GetID() int {
	return o.ID
}

func (o *Routine) NewObjectOnlyID(id int) interface{} {
	ooid := new(Routine)
	ooid.ID = id
	return ooid
}

type RoutineModel struct {
	base.BaseModel
}

func NewRoutineModel() *RoutineModel {
	m := new(RoutineModel)
	m.BaseModel = *base.NewBaseModel(&Routine{})
	return m
}

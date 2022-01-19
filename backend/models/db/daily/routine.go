package daily

import (
	"backend/models/db/base"
)

type Routine struct {
	ID        int    `orm:"column(id)" json:"id"`
	ShortName string `json:"short_name"`
	Event     string `json:"event"`
	WillSpend int    `json:"will_spend"`
	Icon      string `json:"icon"`
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

package daily

import (
	"backend/models/db/base"
)

type Routine struct {
	Id           int    `form:"id"`
	ShortName    string `form:"short_name"`
	Event        string `form:"event"`
	CurrentFocus string `form:"current_focus"`
	WillSpend    int    `form:"will_spend"`
	Icon         string `form:"icon"`
}

func (o *Routine) TableName() string {
	return "daily_time_routine"
}

func (o *Routine) GetID() int {
	return o.Id
}

func (o *Routine) NewObjectOnlyID(id int) interface{} {
	ooid := new(Routine)
	ooid.Id = id
	return ooid
}

type RoutineModel struct {
	base.BaseModel
}

func NewGroupModel() *RoutineModel {
	m := new(RoutineModel)
	m.BaseModel = *base.NewBaseModel(&Routine{})
	return m
}

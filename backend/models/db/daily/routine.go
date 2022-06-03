package daily

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Routine struct {
	ID           int     `orm:"pk;column(id)" json:"id"`
	UserID       int     `orm:"column(user_id)"`
	Icon         string  `json:"icon"`
	ShortName    string  `json:"short_name"`
	EventScope   string  `json:"event_scope"`
	WillSpend    int     `json:"will_spend"`
	HistorySpend float64 `json:"history_spend"`
	Object       int     `json:"object"`
	ObjectUnit   string  `json:"object_unit"`
	Progress     int     `json:"progress"`
	StartDate    string  `json:"start_date"`
	Sort         int
}

func (o *Routine) TableName() string {
	return "daily_time_routine"
}

func (o *Routine) GetID() int {
	return o.ID
}

func (o *Routine) NewObject() interface{} {
	return new(Routine)
}

type RoutineModel struct {
	base.UserBaseModel
}

func NewRoutineModel(userID int) *RoutineModel {
	m := new(RoutineModel)
	m.UserBaseModel = *base.NewUserBaseModelSTD(&Routine{}, userID)

	return m
}

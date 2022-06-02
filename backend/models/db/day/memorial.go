package day

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Memorial struct {
	ID     int `orm:"column(id)"`
	UserID int `orm:"column(user_id)"`
	Date   string
	Desc   string
}

func (o *Memorial) TableName() string {
	return "day_memorial"
}

func (o *Memorial) GetID() int {
	return o.ID
}

func (o *Memorial) NewObject() interface{} {
	return new(Memorial)
}

type MemorialModel struct {
	base.UserBaseModel
}

func NewMemorialModel(userID int) *MemorialModel {
	m := new(MemorialModel)
	config := base.UserBaseModelConfig{
		Object: &Memorial{},
		UserID: userID,
	}
	m.UserBaseModel = *base.NewUserBaseModel(config)

	return m
}

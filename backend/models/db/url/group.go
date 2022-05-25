package url

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Group struct {
	ID   int `orm:"column(id)"`
	Desc string
}

func (o *Group) TableName() string {
	return "url_group"
}

func (o *Group) GetID() int {
	return o.ID
}

func (o *Group) NewObjectOnlyID(id int) interface{} {
	ooid := new(Group)
	ooid.ID = id

	return ooid
}

type GroupModel struct {
	base.BaseModel
}

func NewGroupModel() *GroupModel {
	m := new(GroupModel)
	m.BaseModel = *base.NewBaseModel(&Group{})

	return m
}

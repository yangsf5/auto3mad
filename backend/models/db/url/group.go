package url

import (
	"backend/models/db/base"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(Group))
}

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

package url

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Group struct {
	ID     int `orm:"pk;column(id)"`
	UserID int `orm:"column(user_id)"`
	Desc   string
}

func (o *Group) TableName() string {
	return "url_group"
}

func (o *Group) GetID() int {
	return o.ID
}

func (o *Group) NewObject() interface{} {
	return new(Group)
}

type GroupModel struct {
	base.UserBaseModel
}

func NewGroupModel(userID int) *GroupModel {
	m := new(GroupModel)
	m.UserBaseModel = *base.NewUserBaseModelSTD(&Group{}, userID)

	return m
}

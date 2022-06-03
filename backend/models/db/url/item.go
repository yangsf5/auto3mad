package url

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Item struct {
	ID      int `orm:"column(id)"`
	UserID  int `orm:"column(user_id)"`
	Icon    string
	URL     string `orm:"column(url)"`
	Title   string
	GroupID int `orm:"column(group_id)"`
}

func (o *Item) TableName() string {
	return "url_item"
}

func (o *Item) GetID() int {
	return o.ID
}

func (o *Item) NewObject() interface{} {
	return new(Item)
}

type ItemModel struct {
	base.UserBaseModel
}

func NewItemModel(userID int) *ItemModel {
	m := new(ItemModel)
	m.UserBaseModel = *base.NewUserBaseModelSTD(&Item{}, userID)

	return m
}

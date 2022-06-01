package url

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type Item struct {
	ID      int `orm:"column(id)"`
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

func (o *Item) NewObjectOnlyID(id int) interface{} {
	ooid := new(Item)
	ooid.ID = id

	return ooid
}

type ItemModel struct {
	base.Model
}

func NewItemModel() *ItemModel {
	m := new(ItemModel)
	m.Model = *base.NewModel(&Item{})

	return m
}

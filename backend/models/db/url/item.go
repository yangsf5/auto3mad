package url

import (
	"backend/models/db/base"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(Item))
}

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
	base.BaseModel
}

func NewItemModel() *ItemModel {
	m := new(ItemModel)
	m.BaseModel = *base.NewBaseModel(&Item{})
	return m
}

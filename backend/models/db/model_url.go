package db

import (
	"fmt"
)

const DB_TABLE_URL_GROUP = "url_group"
const DB_TABLE_URL_ITEM = "url_item"

type ModelURL struct {
}

type URLGroup struct {
	ID   int `orm:"column(id)"`
	Desc string
}

func (i *URLGroup) TableName() string {
	return DB_TABLE_URL_GROUP
}

type URLItem struct {
	ID      int `orm:"column(id)"`
	Icon    string
	URL     string `orm:"column(url)"`
	Title   string
	GroupID int `orm:"column(group_id)"`
}

func (i *URLItem) TableName() string {
	return DB_TABLE_URL_ITEM
}

func (m *ModelURL) GetAllItems() (items []URLItem, err error) {
	sql := fmt.Sprintf("SELECT * from %s", DB_TABLE_URL_ITEM)
	_, err = getOrm().Raw(sql).QueryRows(&items)
	return
}

func (m *ModelURL) GetAllGroups() (groups []URLGroup, err error) {
	sql := fmt.Sprintf("SELECT * from %s", DB_TABLE_URL_GROUP)
	_, err = getOrm().Raw(sql).QueryRows(&groups)
	return
}

package db

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

const DB_TABLE_URL_GROUP = "url_group"
const DB_TABLE_URL_ITEM = "url_item"

func init() {
	orm.RegisterModel(new(URLGroup))
}

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
	sql := fmt.Sprintf("SELECT * FROM %s", DB_TABLE_URL_ITEM)
	_, err = getOrm().Raw(sql).QueryRows(&items)
	return
}

func (m *ModelURL) GetAllGroups() (groups []URLGroup, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s", DB_TABLE_URL_GROUP)
	_, err = getOrm().Raw(sql).QueryRows(&groups)
	return
}

func (m *ModelURL) UpsertGroup(group URLGroup) (err error) {
	key := URLGroup{
		ID: group.ID,
	}

	o := getOrm()
	err = o.Read(&key)
	if err == orm.ErrNoRows {
		_, err = o.Insert(&group)
	} else if err != nil {
		return
	} else {
		_, err = o.Update(&group)
	}
	return
}

func (m *ModelURL) DeleteGroup(id int) (err error) {
	_, err = getOrm().Delete(&URLGroup{ID: id})
	return
}

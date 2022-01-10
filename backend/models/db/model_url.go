package db

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

const DB_TABLE_URL_GROUP = "url_group"
const DB_TABLE_URL_ITEM = "url_item"

func init() {
	orm.RegisterModel(new(URLGroup), new(URLItem))
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
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY group_id", DB_TABLE_URL_ITEM)
	_, err = getOrm().Raw(sql).QueryRows(&items)
	return
}

func (m *ModelURL) UpsertItem(item URLItem) (err error) {
	key := URLItem{
		ID: item.ID,
	}

	o := getOrm()
	err = o.Read(&key)
	if err == orm.ErrNoRows {
		_, err = o.Insert(&item)
	} else if err != nil {
		return
	} else {
		_, err = o.Update(&item)
	}
	return
}

func (m *ModelURL) DeleteItem(id int) (err error) {
	_, err = getOrm().Delete(&URLItem{ID: id})
	return
}

func (m *ModelURL) GetMaxID(kind string) (max int, err error) {
	type Ret struct {
		Max int
	}

	dbTable := ""
	if kind == "group" {
		dbTable = DB_TABLE_URL_GROUP
	} else if kind == "item" {
		dbTable = DB_TABLE_URL_ITEM
	} else {
		return -1, errors.New("Kind must be group or item.")
	}

	ret := Ret{}
	sql := fmt.Sprintf("SELECT MAX(id) AS max FROM %s", dbTable)
	err = getOrm().Raw(sql).QueryRow(&ret)
	return ret.Max, err
}

package db

import (
	"backend/models/db/base"
	"backend/models/db/url"
	"errors"
	"fmt"
)

func GetMaxID(kind string) (max int, err error) {
	type Ret struct {
		Max int
	}

	dbTable := ""
	if kind == "group" {
		dbTable = url.DB_TABLE_URL_GROUP
	} else if kind == "item" {
		dbTable = url.DB_TABLE_URL_ITEM
	} else if kind == "memorial" {
		dbTable = "day_memorial"
	} else {
		return -1, errors.New("Kind must be one of group/item/memorial.")
	}

	ret := Ret{}
	sql := fmt.Sprintf("SELECT MAX(id) AS max FROM %s", dbTable)
	err = base.GetOrm().Raw(sql).QueryRow(&ret)
	return ret.Max, err
}

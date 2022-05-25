package db

import (
	"errors"
	"fmt"

	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

func GetMaxID(kind string) (max int, err error) {
	type Ret struct {
		Max int
	}

	dbTable := ""

	if kind == "group" {
		dbTable = "url_group"
	} else if kind == "item" {
		dbTable = "url_item"
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

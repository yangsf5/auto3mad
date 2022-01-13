package db

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conn)
}

func getOrm() orm.Ormer {
	o := orm.NewOrmUsingDB("default")
	return o
}

func GetMaxID(kind string) (max int, err error) {
	type Ret struct {
		Max int
	}

	dbTable := ""
	if kind == "group" {
		dbTable = DB_TABLE_URL_GROUP
	} else if kind == "item" {
		dbTable = DB_TABLE_URL_ITEM
	} else if kind == "memorial" {
		dbTable = DB_TABLE_DAY_MEMORIAL
	} else {
		return -1, errors.New("Kind must be one of group/item/memorial.")
	}

	ret := Ret{}
	sql := fmt.Sprintf("SELECT MAX(id) AS max FROM %s", dbTable)
	err = getOrm().Raw(sql).QueryRow(&ret)
	return ret.Max, err
}

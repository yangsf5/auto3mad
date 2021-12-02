package db

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "user:pwd@tcp(127.0.0.1:3333)/auto3mad?charset=utf8")
}

func getOrm() orm.Ormer {
	o := orm.NewOrmUsingDB("default")
	return o
}

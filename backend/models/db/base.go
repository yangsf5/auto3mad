package db

import (
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

package base

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn, err := web.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conn)
}

func GetOrm() orm.Ormer {
	o := orm.NewOrmUsingDB("default")
	return o
}

type BaseModel struct {
	Orm orm.Ormer
}

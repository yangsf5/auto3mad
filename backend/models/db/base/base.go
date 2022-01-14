package base

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn, err := web.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conn)

	o = orm.NewOrmUsingDB("default")
}

func GetOrm() orm.Ormer {
	return o
}

type BaseModel struct {
}

func (m *BaseModel) GetORM() orm.Ormer {
	return o
}

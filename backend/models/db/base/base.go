package base

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql" // for MySQL Driver
)

var (
	defaultORM     orm.Ormer
	registerModels map[string]bool
)

func Init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	conn, err := web.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}

	_ = orm.RegisterDataBase("default", "mysql", conn)

	orm.Debug = web.AppConfig.DefaultBool("OrmDebug", false)

	defaultORM = orm.NewOrmUsingDB("default")
	registerModels = make(map[string]bool)
}

func GetOrm() orm.Ormer {
	return defaultORM
}

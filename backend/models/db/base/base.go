package base

import (
	"sync"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql" // for MySQL Driver
)

var (
	defaultORM     orm.Ormer
	registerModels sync.Map
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
}

func GetOrm() orm.Ormer {
	return defaultORM
}

type RegisterModelObject interface {
	TableName() string
}

func RegisterModel(obj RegisterModelObject) {
	if _, ok := registerModels.LoadOrStore(obj.TableName(), true); !ok {
		orm.RegisterModel(obj)
	}
}

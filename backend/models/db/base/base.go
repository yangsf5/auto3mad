package base

import (
	"fmt"

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

type BaseModelObject interface {
	TableName() string
	GetID() int
	NewObjectOnlyID(id int) interface{}
}

type BaseModel struct {
	ORM       orm.Ormer
	TableName string
	BMO       BaseModelObject
}

func NewBaseModel(tableName string, bmo BaseModelObject) *BaseModel {
	bm := new(BaseModel)
	bm.ORM = o
	bm.TableName = tableName
	bm.BMO = bmo

	return bm
}

func (m *BaseModel) GetMaxID() (int, error) {
	type Ret struct {
		Max int
	}

	ret := Ret{}
	sql := fmt.Sprintf("SELECT MAX(id) AS max FROM %s", m.TableName)
	err := m.ORM.Raw(sql).QueryRow(&ret)
	return ret.Max, err
}

func (m *BaseModel) GetAll(objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).All(objects)
	return err
}

func (m *BaseModel) GetAllOrderBy(objects interface{}, order string) error {
	_, err := m.ORM.QueryTable(m.TableName).OrderBy(order).All(objects)
	return err
}

func (m *BaseModel) Upsert(item BaseModelObject) error {
	key := m.BMO.NewObjectOnlyID(item.GetID())

	o := m.ORM
	err := o.Read(key)
	if err != nil {
		if err == orm.ErrNoRows {
			_, err = o.Insert(item)
			return err
		}
		return err
	}

	_, err = o.Update(item)
	return err
}

func (m *BaseModel) DeleteByID(id int) error {
	_, err := m.ORM.Delete(m.BMO.NewObjectOnlyID(id))
	return err
}

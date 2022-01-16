package base

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var (
	defaultORM     orm.Ormer
	registerModels map[string]bool
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn, err := web.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conn)

	defaultORM = orm.NewOrmUsingDB("default")
	registerModels = make(map[string]bool)
}

func GetOrm() orm.Ormer {
	return defaultORM
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

func NewBaseModel(bmo BaseModelObject) *BaseModel {
	bm := new(BaseModel)
	bm.ORM = defaultORM
	bm.TableName = bmo.TableName()
	bm.BMO = bmo

	if _, ok := registerModels[bmo.TableName()]; !ok {
		orm.RegisterModel(bmo)
		registerModels[bmo.TableName()] = true
	}

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

func (m *BaseModel) Upsert(obj BaseModelObject) error {
	key := m.BMO.NewObjectOnlyID(obj.GetID())
	err := m.ORM.Read(key)
	if err != nil {
		if err == orm.ErrNoRows {
			_, err = m.ORM.Insert(obj)
			return err
		}
		return err
	}

	_, err = m.ORM.Update(obj)
	return err
}

func (m *BaseModel) DeleteByID(id int) error {
	_, err := m.ORM.Delete(m.BMO.NewObjectOnlyID(id))
	return err
}

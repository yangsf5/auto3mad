package base

import (
	"errors"

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

	defaultORM = orm.NewOrmUsingDB("default")
	registerModels = make(map[string]bool)
}

func GetOrm() orm.Ormer {
	return defaultORM
}

type ModelObject interface {
	TableName() string
	GetID() int
	NewObjectOnlyID(id int) interface{}
}

type Model struct {
	ORM       orm.Ormer
	TableName string
	Object    ModelObject
}

func NewModel(obj ModelObject) *Model {
	bm := new(Model)
	bm.ORM = defaultORM
	bm.TableName = obj.TableName()
	bm.Object = obj

	if _, ok := registerModels[obj.TableName()]; !ok {
		orm.RegisterModel(obj)

		registerModels[obj.TableName()] = true
	}

	return bm
}

func (m *Model) GetAll(objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).All(objects)
	return err
}

func (m *Model) GetAllOrderBy(objects interface{}, order string) error {
	_, err := m.ORM.QueryTable(m.TableName).OrderBy(order).All(objects)
	return err
}

func (m *Model) Upsert(obj ModelObject) error {
	key := m.Object.NewObjectOnlyID(obj.GetID())

	err := m.ORM.Read(key)
	if err != nil {
		if errors.Is(err, orm.ErrNoRows) {
			_, err = m.ORM.Insert(obj)
			return err
		}

		return err
	}

	_, err = m.ORM.Update(obj)

	return err
}

func (m *Model) DeleteByID(id int) error {
	_, err := m.ORM.Delete(m.Object.NewObjectOnlyID(id))
	return err
}

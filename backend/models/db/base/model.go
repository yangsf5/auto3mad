package base

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type ModelObject interface {
	TableName() string
	GetID() int
	NewObject() interface{}
}

type ModelConfig struct {
	Object  ModelObject
	IDField string
}

type Model struct {
	ORM       orm.Ormer
	Config    ModelConfig
	TableName string
}

// 默认的标准形式：IDField 为常规值 id
func NewModelSTD(obj ModelObject) *Model {
	cfg := ModelConfig{
		Object:  obj,
		IDField: "id",
	}

	return NewModel(cfg)
}

func NewModel(cfg ModelConfig) *Model {
	m := new(Model)
	m.ORM = defaultORM
	m.Config = cfg
	m.TableName = cfg.Object.TableName()

	RegisterModel(cfg.Object)

	return m
}

func (m *Model) GetAll(objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).All(objects)
	return err
}

func (m *Model) GetAllOrderBy(order string, objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).OrderBy(order).All(objects)
	return err
}

func (m *Model) Upsert(obj ModelObject) error {
	oldObj := m.Config.Object.NewObject()

	err := m.ORM.QueryTable(m.TableName).Filter(m.Config.IDField, obj.GetID()).One(oldObj)
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

func (m *Model) Delete(id int) error {
	_, err := m.ORM.QueryTable(m.TableName).Filter(m.Config.IDField, id).Delete()
	return err
}

func (m *Model) ReadOrCreate(obj ModelObject, col1 string, cols ...string) (bool, int64, error) {
	return m.ORM.ReadOrCreate(obj, col1, cols...)
}

package base

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

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
	m := new(Model)
	m.ORM = defaultORM
	m.TableName = obj.TableName()
	m.Object = obj

	RegisterModel(obj)

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

func (m *Model) Delete(id int) error {
	_, err := m.ORM.Delete(m.Object.NewObjectOnlyID(id))
	return err
}

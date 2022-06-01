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

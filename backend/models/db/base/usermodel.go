package base

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type UID struct {
	ID     int
	UserID int
}

type UserBaseModelObject interface {
	TableName() string
	GetID() int
	NewObject() interface{}
}

type UserBaseModel struct {
	ORM       orm.Ormer
	TableName string
	Object    UserBaseModelObject
	UserID    int
}

func NewUserBaseModel(userID int, obj UserBaseModelObject) *UserBaseModel {
	bm := new(UserBaseModel)
	bm.ORM = defaultORM
	bm.TableName = obj.TableName()
	bm.Object = obj
	bm.UserID = userID

	if _, ok := registerModels[obj.TableName()]; !ok {
		orm.RegisterModel(obj)

		registerModels[obj.TableName()] = true
	}

	return bm
}

func (m *UserBaseModel) GetAll(objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).Filter("user_id", m.UserID).All(objects)
	return err
}

func (m *UserBaseModel) GetAllOrderBy(order string, objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).Filter("user_id", m.UserID).OrderBy(order).All(objects)
	return err
}

func (m *UserBaseModel) Upsert(obj UserBaseModelObject) error {
	oldObj := m.Object.NewObject()

	// 检查 id 是否属于 user_id 用户
	err := m.ORM.QueryTable(m.TableName).Filter("id", obj.GetID()).Filter("user_id", m.UserID).One(oldObj)
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

func (m *UserBaseModel) Delete(id int) error {
	// 检查 id 是否属于 user_id 用户
	_, err := m.ORM.QueryTable(m.TableName).Filter("id", id).Filter("user_id", m.UserID).Delete()
	return err
}

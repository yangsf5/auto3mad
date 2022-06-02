package base

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type UserBaseModelObject interface {
	TableName() string
	GetID() int
	NewObject() interface{}
}

type UserBaseModelConfig struct {
	Object      UserBaseModelObject
	IDField     string // 不设置，则默认为 id
	UserIDField string // 不设置，则默认为 user_id
	UserID      int
}

type UserBaseModel struct {
	ORM       orm.Ormer
	Config    UserBaseModelConfig
	TableName string
}

func NewUserBaseModel(config UserBaseModelConfig) *UserBaseModel {
	bm := new(UserBaseModel)
	bm.ORM = defaultORM

	// 不设置，则默认为 id
	if config.IDField == "" {
		config.IDField = "id"
	}

	// 不设置，则默认为 user_id
	if config.UserIDField == "" {
		config.UserIDField = "user_id"
	}

	bm.Config = config

	obj := config.Object
	bm.TableName = obj.TableName()

	if _, ok := registerModels[obj.TableName()]; !ok {
		orm.RegisterModel(obj)

		registerModels[obj.TableName()] = true
	}

	return bm
}

func (m *UserBaseModel) GetAll(objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).Filter(m.Config.UserIDField, m.Config.UserID).All(objects)
	return err
}

func (m *UserBaseModel) GetAllOrderBy(order string, objects interface{}) error {
	_, err := m.ORM.QueryTable(m.TableName).Filter(m.Config.UserIDField, m.Config.UserID).OrderBy(order).All(objects)
	return err
}

func (m *UserBaseModel) Upsert(obj UserBaseModelObject) error {
	oldObj := m.Config.Object.NewObject()

	// 检查 id 是否属于 user_id 用户
	err := m.ORM.QueryTable(m.TableName).
		Filter(m.Config.IDField, obj.GetID()).
		Filter(m.Config.UserIDField, m.Config.UserID).
		One(oldObj)
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
	_, err := m.ORM.QueryTable(m.TableName).
		Filter(m.Config.IDField, id).Filter(m.Config.UserIDField, m.Config.UserID).Delete()

	return err
}

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
	IDField     string
	UserIDField string
	UserID      int
}

type UserBaseModel struct {
	ORM       orm.Ormer
	Config    UserBaseModelConfig
	TableName string
}

// 默认的标准形式：IDField/UserIDField 为常规默认值
func NewUserBaseModelSTD(obj UserBaseModelObject, userID int) *UserBaseModel {
	config := UserBaseModelConfig{
		Object:      obj,
		IDField:     "id",
		UserIDField: "user_id",
		UserID:      userID,
	}

	return NewUserBaseModel(config)
}

func NewUserBaseModel(config UserBaseModelConfig) *UserBaseModel {
	bm := new(UserBaseModel)
	bm.ORM = defaultORM
	bm.Config = config
	bm.TableName = config.Object.TableName()

	RegisterModel(config.Object)

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

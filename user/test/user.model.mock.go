package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
)

type UserModelMock struct {
	*core.Model[*user.User]
	*core.Mockable
}

func NewUserModelMock(module *user.UserModule) *UserModelMock {
	return &UserModelMock{
		Model:    core.NewModel[*user.User]("UserModel"),
		Mockable: core.NewMockable(),
	}
}

func (um *UserModelMock) CountAll() (int64, error) {
	um.MethodCalled("CountAll")
	return um.CallFunc("CountAll").(int64), nil
}

func (um *UserModelMock) QueryFindAll(q query.QueryFilter) ([]*user.User, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.CallFunc("QueryFindAll").([]*user.User), nil
}

func (um *UserModelMock) DeleteByID(id string) error {
	um.MethodCalled("DeleteByID", id)
	um.CallFunc("DeleteByID")
	return nil
}

func (um *UserModelMock) FindByID(id string) (*user.User, error) {
	um.MethodCalled("FindByID", id)
	return um.CallFunc("FindByID").(*user.User), nil
}

func (um *UserModelMock) FindOneBy(field string, value any) (*user.User, error) {
	um.MethodCalled("FindOneBy", field, value)
	return um.CallFunc("FindOneBy").(*user.User), nil
}

func (um *UserModelMock) Create(item *user.User) error {
	um.MethodCalled("Create", item)
	um.CallFunc("Create")
	return nil
}

func (um *UserModelMock) UpdateByID(id string, item *user.User) error {
	um.MethodCalled("UpdateByID", id, item)
	um.CallFunc("UpdateByID")
	return nil
}

func (um *UserModelMock) CountBy(field string, value any) (int64, error) {
	um.MethodCalled("CountBy", field, value)
	return um.CallFunc("CountBy").(int64), nil
}

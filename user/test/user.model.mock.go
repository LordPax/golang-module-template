package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
)

type UserModelMock struct {
	*core.Model[*user.User]
	*core.Mockable[*user.User]
}

func NewUserModelMock(module *user.UserModule) *UserModelMock {
	return &UserModelMock{
		Model:    core.NewModel[*user.User]("UserModel"),
		Mockable: core.NewMockable[*user.User](),
	}
}

func (um *UserModelMock) CountAll() (int64, error) {
	um.MethodCalled("CountAll")
	return int64(len(um.GetItems())), nil
}

func (um *UserModelMock) QueryFindAll(q query.QueryFilter) ([]*user.User, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.GetItems(), nil
}

func (um *UserModelMock) DeleteByID(id string) error {
	um.MethodCalled("DeleteByID", id)
	return nil
}

func (um *UserModelMock) FindByID(id string) (*user.User, error) {
	um.MethodCalled("FindByID", id)
	return um.GetItems()[0], nil
}

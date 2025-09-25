package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
)

type UserModelMock struct {
	*core.ModelMock[*user.User]
}

func NewUserModelMock(module *user.UserModule) *UserModelMock {
	return &UserModelMock{
		ModelMock: core.NewModelMock[*user.User]("UserModel"),
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

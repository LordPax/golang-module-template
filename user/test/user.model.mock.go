package user_test

import (
	"golang-api/query"
	"golang-api/user"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type UserModelMock struct {
	*common.ModelMock[*user.User]
}

func NewUserModelMock(module core.IModule) *UserModelMock {
	return &UserModelMock{
		ModelMock: common.NewModelMock[*user.User]("UserModel"),
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

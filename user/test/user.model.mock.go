package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
)

type UserModelMock struct {
	*core.Model[*user.User]
	users []*user.User
}

func NewUserModelMock(module *user.UserModule) *UserModelMock {
	return &UserModelMock{
		Model: core.NewModel[*user.User]("UserModel"),
	}
}

func (um *UserModelMock) SetStubUsers(users []*user.User) {
	um.users = users
}

func (um *UserModelMock) SetStubUser(users *user.User) {
	um.users = []*user.User{users}
}

func (um *UserModelMock) CountAll() (int64, error) {
	return int64(len(um.users)), nil
}

func (um *UserModelMock) QueryFindAll(q query.QueryFilter) ([]*user.User, error) {
	return um.users, nil
}

func (um *UserModelMock) DeleteByID(id string) error {
	return nil
}

func (um *UserModelMock) FindByID(id string) (*user.User, error) {
	return um.users[0], nil
}

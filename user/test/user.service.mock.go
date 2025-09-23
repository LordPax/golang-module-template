package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"

	"github.com/LordPax/sockevent"
)

type UserServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewUserService(module *user.UserModule) *UserServiceMock {
	return &UserServiceMock{
		Provider: core.NewProvider("UserServiceMock"),
	}
}

func (us *UserServiceMock) FindAll(q query.QueryFilter) ([]*user.User, error) {
	us.MethodCalled("FindAll", q)
	return us.CallFunc("FindAll").([]*user.User), nil
}

func (us *UserServiceMock) FindByID(id string) (*user.User, error) {
	us.MethodCalled("FindByID", id)
	return us.CallFunc("FindByID").(*user.User), nil
}

func (us *UserServiceMock) FindOneBy(field string, value any) (*user.User, error) {
	us.MethodCalled("FindOneBy", field, value)
	return us.CallFunc("FindOneBy").(*user.User), nil
}

func (us *UserServiceMock) Create(user *user.User) error {
	us.MethodCalled("Create", user)
	us.CallFunc("Create")
	return nil
}

func (us *UserServiceMock) Update(user *user.User) error {
	us.MethodCalled("Update", user)
	us.CallFunc("Update")
	return nil
}

func (us *UserServiceMock) Delete(id string) error {
	us.MethodCalled("Delete", id)
	us.CallFunc("Delete")
	return nil
}

func (us *UserServiceMock) IsUserExists(email, username string) bool {
	us.MethodCalled("IsUserExists", email, username)
	return us.CallFunc("IsUserExists").(bool)
}

func (us *UserServiceMock) CountStats(ws *sockevent.Websocket) map[string]int {
	us.MethodCalled("CountStats", ws)
	return us.CallFunc("CountStats").(map[string]int)
}

package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewUserModuleTest() *user.UserModule {
	module := &user.UserModule{
		Module: core.NewModule("UserTestModule"),
	}

	module.AddProvider(NewUserModelMock(module))
	module.AddProvider(user.NewUserService(module))

	return module
}

func TestSetupUserModule(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	assert.NotNil(t, module)
	assert.NotNil(t, userService)
	assert.NotNil(t, userModel)
}

func TestUserService_FindAll(t *testing.T) {
	q := query.QueryFilter{}
	nbUsers := 3
	expectedUsers := CreateManyUsers(nbUsers)

	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)
	userModel.SetStubUsers(expectedUsers)

	newUsers, _ := userService.FindAll(q)

	assert.Len(t, newUsers, nbUsers)
	for i, user := range newUsers {
		assert.Equal(t, expectedUsers[i].ID, user.ID)
	}

}

func TestUserService_FindByID(t *testing.T) {
	expectedUser := CreateUser()

	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)
	userModel.SetStubUser(expectedUser)

	newUser, _ := userService.FindByID(expectedUser.ID)

	assert.Equal(t, expectedUser.ID, newUser.ID)
}

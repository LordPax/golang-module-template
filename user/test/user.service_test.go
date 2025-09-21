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

	assert.NotNil(t, module, "User module should be created")
	assert.NotNil(t, userService, "UserService should be created")
	assert.NotNil(t, userModel, "UserModel should be created")
}

func TestUserService_FindAll(t *testing.T) {
	q := query.QueryFilter{}
	nbUsers := 3
	expectedUsers := CreateManyUsers(nbUsers)

	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)
	userModel.SetItems(expectedUsers)

	newUsers, _ := userService.FindAll(q)

	called := userModel.IsMethodCalled("QueryFindAll")
	if !assert.Equal(t, called, true, "QueryFindAll method should be called") {
		return
	}

	if !assert.Len(t, newUsers, nbUsers, "Number of users should be equal to expected") {
		return
	}
	for i := 0; i < nbUsers; i++ {
		assert.Equal(t, expectedUsers[i].ID, newUsers[i].ID)
	}

}

func TestUserService_FindByID(t *testing.T) {
	expectedUser := CreateUser()

	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)
	userModel.AddItem(expectedUser)

	newUser, _ := userService.FindByID(expectedUser.ID)

	called := userModel.IsMethodCalled("FindByID")
	if !assert.Equal(t, called, true, "FindByID method should be called") {
		return
	}
	params := userModel.IsParamsEqual("FindByID", expectedUser.ID)
	if !assert.Equal(t, params, true, "FindByID parameter should be the user ID") {
		return
	}

	assert.Equal(t, expectedUser.ID, newUser.ID)
	assert.Equal(t, expectedUser.Email, newUser.Email)
	assert.Equal(t, expectedUser.Firstname, newUser.Firstname)
	assert.Equal(t, expectedUser.Lastname, newUser.Lastname)
	assert.Equal(t, expectedUser.Username, newUser.Username)
	assert.Equal(t, expectedUser.Password, newUser.Password)
}

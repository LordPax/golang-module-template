package user_test

import (
	"golang-api/core"
	"golang-api/query"
	"golang-api/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testUser(t *testing.T, expected *user.User, actual *user.User) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Firstname, actual.Firstname)
	assert.Equal(t, expected.Lastname, actual.Lastname)
	assert.Equal(t, expected.Username, actual.Username)
	assert.Equal(t, expected.Password, actual.Password)
}

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
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	q := query.QueryFilter{}
	nbUsers := 3
	expectedUsers := CreateManyUsers(nbUsers)

	userModel.MockMethod("QueryFindAll", func(params ...any) any { return expectedUsers })

	newUsers, _ := userService.FindAll(q)

	called := userModel.IsMethodCalled("QueryFindAll")
	if !assert.Equal(t, true, called, "QueryFindAll method should be called") {
		return
	}

	if !assert.Len(t, newUsers, nbUsers, "Number of users should be equal to expected") {
		return
	}
	for i := 0; i < nbUsers; i++ {
		testUser(t, expectedUsers[i], newUsers[i])
	}
}

func TestUserService_FindByID(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	expectedUser := CreateUser()

	userModel.MockMethod("FindByID", func(params ...any) any {
		return expectedUser
	})

	newUser, _ := userService.FindByID(expectedUser.ID)

	called := userModel.IsMethodCalled("FindByID")
	if !assert.Equal(t, true, called, "FindByID method should be called") {
		return
	}
	params := userModel.IsParamsEqual("FindByID", expectedUser.ID)
	if !assert.Equal(t, true, params, "FindByID parameter should be the user ID") {
		return
	}

	testUser(t, expectedUser, newUser)
}

func TestUserService_FindOneBy(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	expectedUser := CreateUser()

	userModel.MockMethod("FindOneBy", func(params ...any) any {
		return expectedUser
	})

	newUser, _ := userService.FindOneBy("email", expectedUser.Email)

	called := userModel.IsMethodCalled("FindOneBy")
	if !assert.Equal(t, true, called, "FindOneBy method should be called") {
		return
	}
	params := userModel.IsParamsEqual("FindOneBy", "email", expectedUser.Email)
	if !assert.Equal(t, true, params, "FindOneBy parameters should be the field and value") {
		return
	}

	testUser(t, expectedUser, newUser)
}

func TestUserService_Create(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	expectedUser := CreateUser()
	var createdUser *user.User

	userModel.MockMethod("Create", func(params ...any) any {
		createdUser = params[0].(*user.User)
		return nil
	})

	err := userService.Create(expectedUser)
	if !assert.Nil(t, err, "Create should not return an error") {
		return
	}

	called := userModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}
	params := userModel.GetMethodParams("Create")
	if !assert.Len(t, params, 1, "Create should have one parameter") {
		return
	}
	paramUser, ok := params[0].(*user.User)
	if !assert.Equal(t, true, ok, "Create parameter should be a user") {
		return
	}

	testUser(t, expectedUser, paramUser)
	testUser(t, expectedUser, createdUser)
}

func TestUserService_Update(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	expectedUser := CreateUser()
	var updatedUser *user.User

	userModel.MockMethod("UpdateByID", func(params ...any) any {
		updatedUser = params[1].(*user.User)
		return nil
	})

	err := userService.Update(expectedUser)
	if !assert.Nil(t, err, "Update should not return an error") {
		return
	}

	called := userModel.IsMethodCalled("UpdateByID")
	if !assert.Equal(t, true, called, "UpdateByID method should be called") {
		return
	}
	params := userModel.GetMethodParams("UpdateByID")
	if !assert.Len(t, params, 2, "UpdateByID should have two parameters") {
		return
	}
	assert.Equal(t, params[0], expectedUser.ID, "First parameter should be the user ID")
	paramUser, ok := params[1].(*user.User)
	if !assert.Equal(t, true, ok, "Second parameter should be a user") {
		return
	}

	testUser(t, expectedUser, paramUser)
	testUser(t, expectedUser, updatedUser)
}

func TestUserService_Delete(t *testing.T) {
	module := NewUserModuleTest()
	userService := module.Get("UserService").(*user.UserService)
	userModel := module.Get("UserModel").(*UserModelMock)

	expectedUser := CreateUser()

	userModel.MockMethod("DeleteByID", nil)

	err := userService.Delete(expectedUser.ID)
	if !assert.Nil(t, err, "Delete should not return an error") {
		return
	}

	called := userModel.IsMethodCalled("DeleteByID")
	if !assert.Equal(t, true, called, "DeleteByID method should be called") {
		return
	}
	params := userModel.IsParamsEqual("DeleteByID", expectedUser.ID)
	if !assert.Equal(t, true, params, "DeleteByID parameter should be the user ID") {
		return
	}
}

// func TestUserService_IsUserExists(t *testing.T) {
// 	expectedUser := CreateUser()

// 	module := NewUserModuleTest()
// 	userService := module.Get("UserService").(*user.UserService)
// 	userModel := module.Get("UserModel").(*UserModelMock)
// 	userModel.MockMethod("CountBy", func(params ...any) any { return int64(1) })
// 	userModel.AddItem(expectedUser)

// 	exists := userService.IsUserExists(expectedUser.Email, expectedUser.Username)
// 	if !assert.Equal(t, exists, true, "IsUserExists should return true") {
// 		return
// 	}

// 	called := userModel.IsMethodCalled("CountBy")
// 	if !assert.Equal(t, called, true, "CountBy method should be called") {
// 		return
// 	}
// 	paramsEmail := userModel.IsParamsEqual("CountBy", "email", expectedUser.Email)
// 	if !assert.Equal(t, paramsEmail, true, "CountBy parameters should be the email field and value") {
// 		return
// 	}
// }

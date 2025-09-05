package user

import (
	"golang-api/core"
	"golang-api/database"
)

var module *UserModule

type UserModule struct {
	*core.Module
}

func NewUserModule() *UserModule {
	module := &UserModule{
		Module: core.NewModule("UserModule"),
	}

	dbModule := database.Module()
	dbService := dbModule.GetProvider("DatabaseService").(*database.DatabaseService)
	userModel := NewUserModel(dbService)
	userService := NewUserService(userModel)
	userMiddleware := NewUserMiddleware(userService)

	module.AddModule(database.Module())
	module.AddProvider(userModel)
	module.AddProvider(userService)
	module.AddProvider(userMiddleware)
	module.AddProvider(NewUserController(userService, userMiddleware))

	return module
}

func Module() *UserModule {
	if module == nil {
		module = NewUserModule()
	}
	return module
}

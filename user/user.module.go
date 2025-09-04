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

	module.AddModule(database.Module())
	module.AddProvider(userModel)
	module.AddProvider(NewUserService(userModel))

	return module
}

func Module() *UserModule {
	if module == nil {
		module = NewUserModule()
	}
	return module
}

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

	module.AddModule(database.Module())
	module.AddProvider(NewUserModel(module))
	module.AddProvider(NewUserService(module))
	module.AddProvider(NewUserMiddleware(module))
	module.AddProvider(NewUserController(module))

	return module
}

func Module() *UserModule {
	if module == nil {
		module = NewUserModule()
	}
	return module
}

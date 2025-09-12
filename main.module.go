package main

import (
	"golang-api/auth"
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
	logUser "golang-api/log-user"
	"golang-api/user"
)

var module *MainModule

type MainModule struct {
	*core.Module
}

func NewMainModule() *MainModule {
	module := &MainModule{
		Module: core.NewModule("MainModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddModule(database.Module())
	module.AddModule(user.Module())
	module.AddModule(auth.Module())
	module.AddModule(logUser.Module())
	module.AddProvider(NewMainService(module))

	return module
}

func Module() *MainModule {
	if module == nil {
		module = NewMainModule()
	}
	return module
}

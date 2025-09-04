package main

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
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

	dotenvModule := dotenv.Module()
	dotenvService := dotenvModule.GetProvider("DotenvService").(*dotenv.DotenvService)

	dbModule := database.Module()
	dbService := dbModule.GetProvider("DatabaseService").(*database.DatabaseService)

	module.AddModule(dotenvModule)
	module.AddModule(dbModule)
	module.AddModule(user.Module())
	module.AddProvider(NewMainService(dbService, dotenvService))

	return module
}

func Module() *MainModule {
	if module == nil {
		module = NewMainModule()
	}
	return module
}

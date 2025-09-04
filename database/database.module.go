package database

import (
	"golang-api/core"
	"golang-api/dotenv"
)

var module *DatabaseModule

type DatabaseModule struct {
	*core.Module
}

func NewDatabaseModule() *DatabaseModule {
	module := &DatabaseModule{
		Module: core.NewModule("DatabaseModule"),
	}

	dotenvModule := dotenv.Module()
	dotenvService := dotenvModule.GetProvider("DotenvService").(*dotenv.DotenvService)

	module.AddModule(dotenvModule)
	module.AddProvider(NewDatabasePostgres(dotenvService))

	return module
}

func Module() *DatabaseModule {
	if module == nil {
		module = NewDatabaseModule()
	}
	return module
}

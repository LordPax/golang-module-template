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

	module.AddModule(dotenv.Module())
	module.AddProvider(NewDatabasePostgres(module))

	return module
}

func Module() *DatabaseModule {
	if module == nil {
		module = NewDatabaseModule()
	}
	return module
}

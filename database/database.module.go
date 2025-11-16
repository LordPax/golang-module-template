package database

import (
	"golang-api/dotenv"

	"github.com/LordPax/godular/core"
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
	module.AddProvider(NewDatabaseService(module))

	return module
}

func Module() *DatabaseModule {
	if module == nil {
		module = NewDatabaseModule()
	}
	return module
}

package log

import (
	"golang-api/database"

	"github.com/LordPax/godular/core"
)

var module *LogModule

type LogModule struct {
	*core.Module
}

func NewLogModule() *LogModule {
	module := &LogModule{
		Module: core.NewModule("LogModule"),
	}

	module.AddModule(database.Module())
	module.AddProvider(NewLogModel(module))
	module.AddProvider(NewLogService(module))
	module.AddProvider(NewLogMiddleware(module))

	return module
}

func Module() *LogModule {
	if module == nil {
		module = NewLogModule()
	}
	return module
}

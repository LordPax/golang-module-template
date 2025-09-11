package log

import (
	"golang-api/core"
	"golang-api/database"
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
	module.AddProvider(NewLogController(module))

	return module
}

func Module() *LogModule {
	if module == nil {
		module = NewLogModule()
	}
	return module
}

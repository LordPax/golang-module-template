package logUser

import (
	"golang-api/gin"
	"golang-api/log"
	"golang-api/query"
	"golang-api/user"

	"github.com/LordPax/godular/core"
)

var module *LogUserModule

type LogUserModule struct {
	*core.Module
}

func NewLogUserModule() *LogUserModule {
	module := &LogUserModule{
		Module: core.NewModule("LogUserModule"),
	}

	module.AddModule(gin.Module())
	module.AddModule(user.Module())
	module.AddModule(log.Module())
	module.AddModule(query.Module())
	module.AddProvider(NewLogUserController(module))

	return module
}

func Module() *LogUserModule {
	if module == nil {
		module = NewLogUserModule()
	}
	return module
}

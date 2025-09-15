package code

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/email"
	"golang-api/gin"
	"golang-api/log"
	"golang-api/query"
	"golang-api/user"
)

var module *CodeModule

type CodeModule struct {
	*core.Module
}

func NewCodeModule() *CodeModule {
	module := &CodeModule{
		Module: core.NewModule("CodeModule"),
	}

	module.AddModule(gin.Module())
	module.AddModule(database.Module())
	module.AddModule(log.Module())
	module.AddModule(query.Module())
	module.AddModule(user.Module())
	module.AddModule(email.Module())
	module.AddProvider(NewCodeModel(module))
	module.AddProvider(NewCodeService(module))
	module.AddProvider(NewCodeController(module))

	return module
}

func Module() *CodeModule {
	if module == nil {
		module = NewCodeModule()
	}
	return module
}

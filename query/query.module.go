package query

import (
	"golang-api/core"
)

var module *QueryModule

type QueryModule struct {
	*core.Module
}

func NewQueryModule() *QueryModule {
	module := &QueryModule{
		Module: core.NewModule("QueryModule"),
	}

	module.AddProvider(NewQueryService(module))

	return module
}

func Module() *QueryModule {
	if module == nil {
		module = NewQueryModule()
	}
	return module
}

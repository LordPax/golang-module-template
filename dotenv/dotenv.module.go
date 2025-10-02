package dotenv

import (
	"golang-api/core"
)

var module *DotenvModule

type DotenvModule struct {
	*core.Module
}

func NewDotenvModule() *DotenvModule {
	module := &DotenvModule{
		Module: core.NewModule("DotenvModule"),
	}

	module.AddProvider(NewDotenvService(module))

	return module
}

func Module() *DotenvModule {
	if module == nil {
		module = NewDotenvModule()
	}
	return module
}

package gin

import (
	"golang-api/core"
	"golang-api/dotenv"
)

var module *GinModule

type GinModule struct {
	*core.Module
}

func NewGinModule() *GinModule {
	module := &GinModule{
		Module: core.NewModule("GinModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddProvider(NewGinService(module))

	return module
}

func Module() *GinModule {
	if module == nil {
		module = NewGinModule()
	}
	return module
}

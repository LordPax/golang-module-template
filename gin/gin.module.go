package gin

import (
	"golang-api/dotenv"

	"github.com/LordPax/godular/core"
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

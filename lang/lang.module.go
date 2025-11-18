package lang

import (
	"github.com/LordPax/godular/core"
)

var module *LangModule

type LangModule struct {
	*core.Module
}

func NewLangModule() *LangModule {
	module := &LangModule{
		Module: core.NewModule("LangModule"),
	}

	module.AddProvider(NewLangService(module))

	return module
}

func Module() *LangModule {
	if module == nil {
		module = NewLangModule()
	}
	return module
}

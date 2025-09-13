package email

import (
	"golang-api/core"
	"golang-api/dotenv"
)

var module *EmailModule

type EmailModule struct {
	*core.Module
}

func NewEmailModule() *EmailModule {
	module := &EmailModule{
		Module: core.NewModule("EmailModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddProvider(NewEmailService(module))

	return module
}

func Module() *EmailModule {
	if module == nil {
		module = NewEmailModule()
	}
	return module
}

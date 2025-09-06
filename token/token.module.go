package token

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
)

var module *TokenModule

type TokenModule struct {
	*core.Module
}

func NewTokenModule() *TokenModule {
	module := &TokenModule{
		Module: core.NewModule("TokenModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddModule(database.Module())
	module.AddProvider(NewTokenModel(module))
	module.AddProvider(NewTokenService(module))

	return module
}

func Module() *TokenModule {
	if module == nil {
		module = NewTokenModule()
	}
	return module
}

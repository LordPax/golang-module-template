package auth

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
	"golang-api/user"
)

var module *AuthModule

type AuthModule struct {
	*core.Module
}

func NewAuthModule() *AuthModule {
	module := &AuthModule{
		Module: core.NewModule("AuthModule"),
	}

	module.AddModule(dotenv.Module())
	module.AddModule(database.Module())
	module.AddModule(user.Module())
	module.AddProvider(NewTokenModel(module))
	module.AddProvider(NewAuthService(module))
	module.AddProvider(NewAuthMiddleware(module))
	module.AddProvider(NewAuthController(module))

	return module
}

func Module() *AuthModule {
	if module == nil {
		module = NewAuthModule()
	}
	return module
}

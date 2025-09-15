package auth

import (
	"golang-api/code"
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
	"golang-api/email"
	"golang-api/gin"
	"golang-api/log"
	"golang-api/token"
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

	module.AddModule(gin.Module())
	module.AddModule(dotenv.Module())
	module.AddModule(email.Module())
	module.AddModule(database.Module())
	module.AddModule(log.Module())
	module.AddModule(token.Module())
	module.AddModule(user.Module())
	module.AddModule(code.Module())
	module.AddProvider(NewAuthService(module))
	module.AddProvider(NewAuthController(module))

	return module
}

func Module() *AuthModule {
	if module == nil {
		module = NewAuthModule()
	}
	return module
}

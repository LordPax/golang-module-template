package main

import (
	"golang-api/auth"
	"golang-api/code"
	"golang-api/core"
	"golang-api/database"
	"golang-api/gin"
	logUser "golang-api/log-user"
	"golang-api/user"
	userWebsocket "golang-api/user-websocket"
	"golang-api/websocket"
)

var module *MainModule

type MainModule struct {
	*core.Module
}

func NewMainModule() *MainModule {
	module := &MainModule{
		Module: core.NewModule("MainModule"),
	}

	module.AddModule(gin.Module())
	module.AddModule(database.Module())
	module.AddModule(user.Module())
	module.AddModule(auth.Module())
	module.AddModule(logUser.Module())
	module.AddModule(websocket.Module())
	module.AddModule(userWebsocket.Module())
	module.AddModule(code.Module())
	module.AddProvider(NewMainService(module))

	return module
}

func Module() *MainModule {
	if module == nil {
		module = NewMainModule()
	}
	return module
}

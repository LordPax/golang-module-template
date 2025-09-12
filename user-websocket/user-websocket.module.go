package userWebsocket

import (
	"golang-api/core"
	"golang-api/user"
	"golang-api/websocket"
)

var module *UserWebsocketModule

type UserWebsocketModule struct {
	*core.Module
}

func NewUserWebsocketModule() *UserWebsocketModule {
	module := &UserWebsocketModule{
		Module: core.NewModule("UserWebsocketModule"),
	}

	module.AddModule(user.Module())
	module.AddModule(websocket.Module())
	module.AddProvider(NewUserWebsocket(module))

	return module
}

func Module() *UserWebsocketModule {
	if module == nil {
		module = NewUserWebsocketModule()
	}
	return module
}

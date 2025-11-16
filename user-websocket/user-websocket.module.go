package userWebsocket

import (
	"golang-api/user"
	"golang-api/websocket"

	"github.com/LordPax/godular/core"
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

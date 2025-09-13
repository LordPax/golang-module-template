package websocket

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/user"
)

var module *WebsocketModule

type WebsocketModule struct {
	*core.Module
}

func NewWebsocketModule() *WebsocketModule {
	module := &WebsocketModule{
		Module: core.NewModule("WebsocketModule"),
	}

	module.AddModule(log.Module())
	module.AddModule(user.Module())
	module.AddProvider(NewWebsocketService(module))
	module.AddProvider(NewPingService(module))
	module.AddProvider(NewWebsocketController(module))

	return module
}

func Module() *WebsocketModule {
	if module == nil {
		module = NewWebsocketModule()
	}
	return module
}

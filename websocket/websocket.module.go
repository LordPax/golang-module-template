package websocket

import (
	"golang-api/gin"
	"golang-api/log"
	"golang-api/user"

	"github.com/LordPax/godular/core"
)

var module *WebsocketModule

type WebsocketModule struct {
	*core.Module
}

func NewWebsocketModule() *WebsocketModule {
	module := &WebsocketModule{
		Module: core.NewModule("WebsocketModule"),
	}

	module.AddModule(gin.Module())
	module.AddModule(log.Module())
	module.AddModule(user.Module())
	module.AddProvider(NewWebsocketService(module))
	module.AddProvider(NewPingWebsocket(module))
	module.AddProvider(NewWebsocketController(module))

	return module
}

func Module() *WebsocketModule {
	if module == nil {
		module = NewWebsocketModule()
	}
	return module
}

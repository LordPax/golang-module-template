package websocket_test

import (
	"golang-api/core"

	"github.com/LordPax/sockevent"
)

type PingWebsocketMock struct {
	*core.Provider
	*core.Mockable
}

func NewPingWebsocketMock(module core.IModule) *PingWebsocketMock {
	return &PingWebsocketMock{
		Provider: core.NewProvider("PingWebsocket"),
		Mockable: core.NewMockable(),
	}
}

func (ws *PingWebsocketMock) Ping(client *sockevent.Client, message any) error {
	ws.MethodCalled("Ping")
	ws.CallFunc("Ping")
	return nil
}

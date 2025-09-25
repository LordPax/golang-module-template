package userWebsocket_test

import (
	"golang-api/core"
	userWebsocket "golang-api/user-websocket"

	"github.com/LordPax/sockevent"
)

type UserWebsocketMock struct {
	*core.Provider
	*core.Mockable
}

func NewUserWebsocketMock(module *userWebsocket.UserWebsocketModule) *UserWebsocketMock {
	return &UserWebsocketMock{
		Provider: core.NewProvider("UserWebsocket"),
		Mockable: core.NewMockable(),
	}
}

func (uw *UserWebsocketMock) UserStats(client *sockevent.Client, message any) error {
	uw.MethodCalled("UserStats", client, message)
	uw.CallFunc("UserStats")
	return nil
}

package websocket_test

import (
	"golang-api/core"
	"net/http"

	"github.com/LordPax/sockevent"
)

type WebsocketServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewWebsocketServiceMock(module core.IModule) *WebsocketServiceMock {
	return &WebsocketServiceMock{
		Provider: core.NewProvider("WebsocketService"),
		Mockable: core.NewMockable(),
	}
}

func (ws *WebsocketServiceMock) GetWs() *sockevent.Websocket {
	return nil
}

func (ws *WebsocketServiceMock) SendNbUserToAdmin() error {
	ws.MethodCalled("SendNbUserToAdmin")
	ws.CallFunc("SendNbUserToAdmin")
	return nil
}

func (ws *WebsocketServiceMock) Connect(client *sockevent.Client, wr http.ResponseWriter, r *http.Request) error {
	ws.MethodCalled("Connect")
	ws.CallFunc("Connect")
	return nil
}

func (ws *WebsocketServiceMock) Disconnect(client *sockevent.Client) error {
	ws.MethodCalled("Disconnect")
	ws.CallFunc("Disconnect")
	return nil
}

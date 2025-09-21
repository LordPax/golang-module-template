package websocket

import (
	"fmt"
	"golang-api/core"
	"golang-api/user"

	"github.com/LordPax/sockevent"
)

type IPingWebsocket interface {
	core.IProvider
	Ping(client *sockevent.Client, message any) error
}

type PingWebsocket struct {
	*core.Provider
	websocketService IWebsocketService
}

func NewPingWebsocket(module *WebsocketModule) *PingWebsocket {
	return &PingWebsocket{
		Provider:         core.NewProvider("PingWebsocket"),
		websocketService: module.Get("WebsocketService").(IWebsocketService),
	}
}

func (ws *PingWebsocket) OnInit() error {
	fmt.Println("Registering Ping websocket")
	ws.websocketService.GetWs().On("ping", ws.Ping)
	return nil
}

func (ws *PingWebsocket) Ping(client *sockevent.Client, message any) error {
	logged := client.Get("logged").(bool)
	if !logged {
		fmt.Printf("Client %s sent message: %v\n", client.ID, message)
		return client.Emit("pong", "pong")
	}

	user := client.Get("user").(*user.User)
	fmt.Printf("Client %s (%s) sent message: %v\n", client.ID, user.Username, message)

	return client.Emit("pong", "pong")
}

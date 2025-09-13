package websocket

import (
	"fmt"
	"golang-api/core"
	"golang-api/user"

	"github.com/LordPax/sockevent"
)

type PingService struct {
	*core.Provider
	websocketService *WebsocketService
}

func NewPingService(module *WebsocketModule) *PingService {
	return &PingService{
		Provider:         core.NewProvider("PingService"),
		websocketService: module.Get("WebsocketService").(*WebsocketService),
	}
}

func (ws *PingService) OnInit() error {
	ws.websocketService.Ws.On("ping", ws.Ping)
	return nil
}

func (ws *PingService) Ping(client *sockevent.Client, message any) error {
	logged := client.Get("logged").(bool)
	if !logged {
		fmt.Printf("Client %s sent message: %v\n", client.ID, message)
		return client.Emit("pong", "pong")
	}

	user := client.Get("user").(*user.User)
	fmt.Printf("Client %s (%s) sent message: %v\n", client.ID, user.Username, message)

	return client.Emit("pong", "pong")
}

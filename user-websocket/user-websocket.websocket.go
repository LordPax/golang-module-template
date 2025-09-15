package userWebsocket

import (
	"fmt"
	"golang-api/core"
	"golang-api/user"
	"golang-api/websocket"

	"github.com/LordPax/sockevent"
)

type UserWebsocket struct {
	*core.Provider
	userService      *user.UserService
	websocketService *websocket.WebsocketService
}

func NewUserWebsocket(module *UserWebsocketModule) *UserWebsocket {
	return &UserWebsocket{
		Provider:         core.NewProvider("UserWebsocket"),
		userService:      module.Get("UserService").(*user.UserService),
		websocketService: module.Get("WebsocketService").(*websocket.WebsocketService),
	}
}

func (uw *UserWebsocket) OnInit() error {
	fmt.Println("Registering User websocket")
	uw.websocketService.Ws.On("user:stats", uw.UserStats)
	return nil
}

func (uw *UserWebsocket) UserStats(client *sockevent.Client, message any) error {
	logged := client.Get("logged").(bool)
	if !logged {
		return nil
	}

	cUser := client.Get("user").(*user.User)
	if !cUser.IsRole(user.ROLE_ADMIN) {
		return nil
	}

	wsData := uw.userService.CountStats(uw.websocketService.Ws)
	return client.Emit("user:connected", wsData)
}

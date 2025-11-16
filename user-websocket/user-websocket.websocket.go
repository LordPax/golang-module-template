package userWebsocket

import (
	"fmt"
	"golang-api/user"
	"golang-api/websocket"

	"github.com/LordPax/godular/core"
	"github.com/LordPax/sockevent"
)

type IUserWebsocket interface {
	core.IProvider
	OnInit() error
	UserStats(client *sockevent.Client, message any) error
}

type UserWebsocket struct {
	*core.Provider
	userService      user.IUserService
	websocketService websocket.IWebsocketService
}

func NewUserWebsocket(module core.IModule) *UserWebsocket {
	return &UserWebsocket{
		Provider:         core.NewProvider("UserWebsocket"),
		userService:      module.Get("UserService").(user.IUserService),
		websocketService: module.Get("WebsocketService").(websocket.IWebsocketService),
	}
}

func (uw *UserWebsocket) OnInit() error {
	fmt.Println("Registering User websocket")
	uw.websocketService.GetWs().On("user:stats", uw.UserStats)
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

	wsData := uw.userService.CountStats(uw.websocketService.GetWs())
	return client.Emit("user:connected", wsData)
}

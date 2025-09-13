package websocket

import (
	"fmt"
	"golang-api/core"
	"golang-api/user"
	"net/http"

	"github.com/LordPax/sockevent"
)

type WebsocketService struct {
	*core.Provider
	Ws          *sockevent.Websocket
	userService *user.UserService
}

func NewWebsocketService(module *WebsocketModule) *WebsocketService {
	return &WebsocketService{
		Provider:    core.NewProvider("WebsocketService"),
		Ws:          sockevent.GetWebsocket(),
		userService: module.Get("UserService").(*user.UserService),
	}
}

func (ws *WebsocketService) OnInit() error {
	ws.Ws.OnConnect(ws.Connect)
	ws.Ws.OnDisconnect(ws.Disconnect)
	return nil
}

func (ws *WebsocketService) SendNbUserToAdmin() error {
	wsData := ws.userService.CountStats(ws.Ws)
	return ws.Ws.Room("admin").Emit("user:connected", wsData)
}

func (ws *WebsocketService) Connect(client *sockevent.Client, wr http.ResponseWriter, r *http.Request) error {
	connectedUser := r.Context().Value("connectedUser")
	ok := connectedUser != nil && connectedUser.(*user.User).ID != ""
	client.Set("logged", ok)

	if err := ws.SendNbUserToAdmin(); err != nil {
		return err
	}

	if !ok {
		fmt.Printf("Client %s connected, len %d\n",
			client.ID,
			len(client.Ws.GetClients()),
		)
		return client.Emit("connected", nil)
	}

	cUser := connectedUser.(*user.User)
	client.Set("user", cUser)

	if cUser.IsRole(user.ROLE_ADMIN) {
		client.Ws.Room("admin").AddClient(client)
	}

	fmt.Printf("Client %s (%s) connected, len %d\n",
		client.ID,
		cUser.Username,
		len(client.Ws.GetClients()),
	)

	return client.Emit("connected", nil)
}

func (ws *WebsocketService) Disconnect(client *sockevent.Client) error {
	logged := client.Get("logged").(bool)

	if err := ws.SendNbUserToAdmin(); err != nil {
		return err
	}

	if !logged {
		fmt.Printf("Client %s disconnected, len %d\n",
			client.ID,
			len(client.Ws.GetClients()),
		)
		return nil
	}

	cUser := client.Get("user").(*user.User)

	fmt.Printf("Client %s (%s) disconnected, len %d\n",
		client.ID,
		cUser.Username,
		len(client.Ws.GetClients()),
	)

	return nil
}

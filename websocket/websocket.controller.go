package websocket

import (
	"context"
	"fmt"
	"golang-api/core"
	ginM "golang-api/gin"
	"golang-api/log"
	"golang-api/user"

	"github.com/LordPax/sockevent"
	"github.com/gin-gonic/gin"
)

type IWebsocketController interface {
	core.IProvider
	RegisterRoutes()
	WsHandler(ws *sockevent.Websocket) gin.HandlerFunc
}

type WebsocketController struct {
	*core.Provider
	websocketService IWebsocketService
	logService       log.ILogService
	userMiddleware   user.IUserMiddleware
	ginService       ginM.IGinService
}

func NewWebsocketController(module *WebsocketModule) *WebsocketController {
	return &WebsocketController{
		Provider:         core.NewProvider("WebsocketController"),
		websocketService: module.Get("WebsocketService").(IWebsocketService),
		logService:       module.Get("LogService").(log.ILogService),
		userMiddleware:   module.Get("UserMiddleware").(user.IUserMiddleware),
		ginService:       module.Get("GinService").(ginM.IGinService),
	}
}

func (wc *WebsocketController) OnInit() error {
	fmt.Println("Registering Websocket routes")
	wc.RegisterRoutes()
	return nil
}

func (wc *WebsocketController) RegisterRoutes() {
	wc.ginService.GetGroup().GET("/ws",
		wc.userMiddleware.IsLoggedIn(false),
		wc.WsHandler(wc.websocketService.GetWs()),
	)
}

func (wc *WebsocketController) WsHandler(ws *sockevent.Websocket) gin.HandlerFunc {
	return func(c *gin.Context) {
		connectedUser, ok := c.Get("connectedUser")
		if !ok {
			ws.WsHandler(c.Writer, c.Request)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "connectedUser", connectedUser)
		ws.WsHandler(c.Writer, c.Request.WithContext(ctx))
	}
}

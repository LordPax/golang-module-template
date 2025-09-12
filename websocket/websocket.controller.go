package websocket

import (
	"context"
	"golang-api/core"
	"golang-api/log"
	"golang-api/user"

	"github.com/LordPax/sockevent"
	"github.com/gin-gonic/gin"
)

type WebsocketController struct {
	*core.Provider
	websocketService *WebsocketService
	logService       *log.LogService
	userMiddleware   *user.UserMiddleware
}

func NewWebsocketController(module *WebsocketModule) *WebsocketController {
	return &WebsocketController{
		Provider:         core.NewProvider("WebsocketController"),
		websocketService: module.Get("WebsocketService").(*WebsocketService),
		logService:       module.Get("LogService").(*log.LogService),
		userMiddleware:   module.Get("UserMiddleware").(*user.UserMiddleware),
	}
}

func (wc *WebsocketController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/ws",
		wc.userMiddleware.IsLoggedIn(false),
		wc.WsHandler(wc.websocketService.Ws),
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

package sdkopen_websocket

import (
	"github.com/gorilla/websocket"
	webBaseHttp "github.com/sdkopen/sdkopen-go-webbase/http"
	"github.com/sdkopen/sdkopen-go-webbase/logging"
	"github.com/sdkopen/sdkopen-go-webbase/server"
	"net/http"
)

var Clients = make(map[*websocket.Conn]string)

var webSocketUpgGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketController struct {
}

type Response struct {
	Message string `json:"message"`
}

func NewWebSocketController() *WebSocketController {
	return &WebSocketController{}
}

func (cc *WebSocketController) Routes() []server.Route {
	return []server.Route{
		{
			URI:      "ws",
			Method:   webBaseHttp.MethodGet,
			Prefix:   server.Api,
			Function: cc.Connect,
		},
	}
}

func (cc *WebSocketController) Connect(ctx server.WebContext) {
	ws, err := webSocketUpgGrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		logging.Error("Error upgrading connection: %v", err)
	}
	reader(ws)
}

func reader(conn *websocket.Conn) {
	defer func() {
		delete(Clients, conn)
		conn.Close()
	}()

	var message EventMessage
	err := conn.ReadJSON(&message)
	if err != nil {
		logging.Error("Error reading message: %v", err)
		return
	}

	for _, webSocketEvent := range WebSocketEvents {
		fn := webSocketEvent.Consumer

		if message.Type == webSocketEvent.Type {
			fn(conn, message)
		}
	}
}

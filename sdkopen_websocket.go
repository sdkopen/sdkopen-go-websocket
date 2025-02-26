package sdkopen_websocket

import (
	"github.com/gorilla/websocket"
	"github.com/sdkopen/sdkopen-go-webbase/logging"
	"github.com/sdkopen/sdkopen-go-webbase/server"
)

var (
	WebSocketEvents []Event
	Clients         = make(map[*websocket.Conn]string)
)

func RegisterWebSocketEvent(event Event) {
	WebSocketEvents = append(WebSocketEvents, event)
}

func Broadcast(msg EventMessage, exclude *websocket.Conn) {
	for client := range Clients {
		if client != exclude {
			err := client.WriteJSON(msg)
			if err != nil {
				logging.Error("Error sending message: %v", err)
			}
		}
	}
}

func Initialize() {
	server.RegisterController(newWebSocketController())
}

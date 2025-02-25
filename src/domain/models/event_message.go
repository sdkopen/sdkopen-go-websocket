package sdkopen_websocket_model

import (
	"github.com/gorilla/websocket"
)

type EventMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Event struct {
	Type     string
	Consumer func(conn *websocket.Conn, message EventMessage)
}

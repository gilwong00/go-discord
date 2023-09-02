package websocket

import "encoding/json"

// Event is the Messages sent over the websocket
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event) error

const (
	EventSendMessage = "send_message"
	EventNewMessage  = "new_message"
	EventJoinChannel = "join_channel"
)

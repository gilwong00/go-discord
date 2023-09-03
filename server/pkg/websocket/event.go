package websocket

import (
	"encoding/json"
	"fmt"
	"time"
)

// Event is the Messages sent over the websocket
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type IncomingMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type ChangeChannelEvent struct {
	Name string `json:"name"`
}

// NewMessageEvent is returned when responding to send_message
type NewMessageEvent struct {
	IncomingMessageEvent
	Sent time.Time `json:"sent"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventBroadcastMessage = "broadcast_message"
	EventNewMessage       = "new_message"
	EventJoinChannel      = "join_channel"
)

func BroadcastMessageEventHandler(event Event, c *Client) error {
	var broadcastEvent IncomingMessageEvent
	if err := json.Unmarshal(event.Payload, &broadcastEvent); err != nil {
		return fmt.Errorf("unable to unmarshal payload in request: %v", err)
	}
	var broadMessage NewMessageEvent
	broadMessage.Sent = time.Now()
	broadMessage.Message = broadcastEvent.Message
	broadMessage.From = broadcastEvent.From
	payload, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}
	var outgoingEvent Event
	outgoingEvent.Payload = payload
	outgoingEvent.Type = EventNewMessage
	// save message to db

	// Broadcast to all other Clients
	for client := range c.manager.clients {
		// Only send to clients inside the same channel
		if client.channel == c.channel {
			client.egress <- outgoingEvent
		}
	}
	return nil
}

func JoinChannelHandler(event Event, c *Client) error {
	var changeChannelEvent ChangeChannelEvent
	if err := json.Unmarshal(event.Payload, &changeChannelEvent); err != nil {
		return fmt.Errorf("unable to unmarshal payload in request: %v", err)
	}
	c.channel = changeChannelEvent.Name
	return nil
}

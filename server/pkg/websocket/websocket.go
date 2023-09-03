package websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event // egress is used to avoid concurrent writes on the WebSocket
	channel    string
}

type ClientMap map[*Client]bool

var (
	// pongWait is how long we will await a pong response from client
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		// graceful close the Connection once this func is done
		c.manager.removeClient(c)
	}()
}

func (c *Client) writeMessages() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		// graceful close if ticker is closing
		c.manager.removeClient(c)
	}()
}

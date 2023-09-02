package websocket

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	/**
	upgrader is used to upgrade incomming HTTP requests into a websocket connection
	*/
	upgrader = websocket.Upgrader{
		// Apply the Origin Checker
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// checkOrigin will check origin and return true if its allowed
func checkOrigin(r *http.Request) bool {
	// Grab the request origin
	origin := r.Header.Get("Origin")
	switch origin {
	// Update this to HTTPS
	case "https://localhost:3000", "http://localhost:3000":
		return true
	default:
		return false
	}
}

// Manager is used to hold references to all registered clients
type Manager struct {
	handlers map[string]EventHandler
}

func NewManager(ctx context.Context) *Manager {
	return &Manager{
		handlers: make(map[string]EventHandler),
	}
}

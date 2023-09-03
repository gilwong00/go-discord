package websocket

import (
	"errors"
	"log"
	"net/http"
	"sync"

	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		switch origin {
		case "https://localhost:3000", "http://localhost:3000":
			return true
		default:
			return false
		}
	},
}

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// Manager is used to hold references to all registered clients
type Manager struct {
	sync.RWMutex
	clients  ClientMap
	handlers map[string]EventHandler
	store    db.Store
}

func NewManager(store db.Store) *Manager {
	return &Manager{
		store:    store,
		clients:  make(ClientMap),
		handlers: make(map[string]EventHandler),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// verify token
	log.Println("New connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, m, m.store)
	// Add the newly created client to the manager
	m.addClient(client)
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) SetupEventHandlers() {
	m.handlers[EventBroadcastMessage] = BroadcastMessageEventHandler
	m.handlers[EventJoinChannel] = JoinChannelHandler
}

func (m *Manager) addClient(client *Client) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

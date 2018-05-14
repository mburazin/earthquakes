package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// clientPusher handles the creation of websockets to clients
// and keeps track of connected clients.
// It is used also to push data to clients through created websockets.
type clientPusher struct {
	clients map[*websocket.Conn]bool // connected clients
}

func newClientPusher() *clientPusher {
	return &clientPusher{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (cp *clientPusher) registerRoutes() {
	http.HandleFunc("/ws", cp.handleConnections)
}

// handleConnections upgrades an incoming HTTP connection request to a websocket
// and keeps track of the client connected
func (cp *clientPusher) handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Cannot upgrade request to a websocket: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Register the new client
	log.Println("Registering new client")
	cp.clients[ws] = true
}

// pushDataToClients will take any arbitrary data structure, convert it to
// JSON using the JSON parser and send the JSON data to the client
func (cp *clientPusher) pushDataToClients(data interface{}) {
	for client := range cp.clients {
		log.Println("Pushing data to client...")
		err := client.WriteJSON(data)
		if err != nil {
			// assume client disconnected
			log.Printf("failed pushing data to client over websocket: %v", err)
			log.Println("removing subscribed client")
			client.Close()
			delete(cp.clients, client)
		}
	}
}

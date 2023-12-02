package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	clients = make(map[*websocket.Conn]bool)
)

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}

	clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, conn)
			return
		}

		// You can handle the received message here
		log.Printf("Received message: %s\n", msg)

		// Example: Echo the message back to the client
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func BroadcastMessages() {
	for {
		// Example: Periodically send a message to all connected clients
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte("Broadcast message"))
			if err != nil {
				log.Println("Error broadcasting message:", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

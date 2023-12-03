package websocket

import (
	// "log"
	// "net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	clients = make(map[*websocket.Conn]bool)
)

// type Room struct {
//     Clients map[*websocket.Conn]bool
// }

// // Map to store rooms
// var rooms map[string]*Room

// // Function to handle joining a room
// func JoinRoom(client *websocket.Conn, roomName string) {
//     room, exists := rooms[roomName]
//     if !exists {
//         rooms[roomName] = &Room{Clients: make(map[*websocket.Conn]bool)}
//         room = rooms[roomName]
//     }
//     room.Clients[client] = true
// }

// // Function to handle leaving a room
// func LeaveRoom(client *websocket.Conn, roomName string) {
//     if room, exists := rooms[roomName]; exists {
//         delete(room.Clients, client)
//     }
// }

// // Function to broadcast a message to a room
// func BroadcastToRoom(roomName string, message []byte) {
//     if room, exists := rooms[roomName]; exists {
//         for client := range room.Clients {
//             err := client.WriteMessage(websocket.TextMessage, message)
//             if err != nil {
//                 // Handle error if message cannot be sent
//             }
//         }
//     }
// }

// func ServeWs(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("Failed to set websocket upgrade:", err)
// 		return
// 	}

// 	clients[conn] = true

// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Error reading message:", err)
// 			delete(clients, conn)
// 			return
// 		}

// 		// You can handle the received message here
// 		log.Printf("Received message: %s\n", msg)

// 		// Example: Echo the message back to the client
// 		for client := range clients {
// 			err := client.WriteMessage(websocket.TextMessage, msg)
// 			if err != nil {
// 				log.Println("Error writing message:", err)
// 				delete(clients, client)
// 				client.Close()
// 			}
// 		}
// 	}
// }

// func BroadcastMessages(text string) {
// 	for {
// 		// Example: Periodically send a message to all connected clients
// 		for client := range clients {
// 			err := client.WriteMessage(websocket.TextMessage, []byte(text))
// 			if err != nil {
// 				log.Println("Error broadcasting message:", err)
// 				delete(clients, client)
// 				client.Close()
// 			}
// 		}
// 	}
// }

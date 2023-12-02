package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
} // use default options

func Ws(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
		//read data from ws
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		//write ws data
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}

		// ws.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		// time.Sleep(time.Second)
	}

	// i := 0
	// for {
	// 	i++
	// 	ws.WriteMessage(websocket.TextMessage, []byte("New message (#"+strconv.Itoa(i)+")"))
	// 	time.Sleep(time.Second)
	// }
}
